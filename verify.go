package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/exp/slices"
)

var daoContractAddress = common.HexToAddress("0x36F4D96Fe0D4Eb33cdC2dC6C0bCA15b9Cdd0d648")
var delegateCashAddress = common.HexToAddress("0x00000000000076a84fef008cdabe6409d2fe638b")
var zeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
var jsonRpcUrl = ""

type registration struct {
	CollectionID int
	Owner        common.Address
	TokenID      int
	IsOneOfOne   bool
	Timestamp    time.Time
}

type validatedRegistration struct {
	Owner   common.Address
	TokenID int
}

type ownerLoader interface {
	OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error)
}

type delegateLoader interface {
	GetDelegationsByDelegate(opts *bind.CallOpts, delegate common.Address) ([]IDelegationRegistryDelegationInfo, error)
}

func optInToRaffleList(inputFilePath string, outputFilePath string) error {
	// Read the CSV file
	f, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	// Connect to the Ethereum network
	client, err := ethclient.Dial(jsonRpcUrl)
	if err != nil {
		return err
	}

	// Bind the contract to an instance of the contract interface
	daoToken, err := NewErc721a(daoContractAddress, client)
	if err != nil {
		return err
	}

	delegateCash, err := NewDelegateCash(delegateCashAddress, client)
	if err != nil {
		return err
	}

	// Parse the records
	parsedRecords := []registration{}
	for _, record := range records {
		parsedRecords = append(parsedRecords, csvLineToRegistration(record))
	}
	debugPrintln("Parsed", len(parsedRecords), "records")

	// Validate the records
	validEntries, err := validateRegistrations(parsedRecords, delegateCash, daoToken)
	if err != nil {
		return err
	}

	debugPrintln("Found", len(validEntries), "valid entries")

	// If we have an output file, write to it, otherwise print.
	if outputFilePath != "" {
		err = writeCSV(outputFilePath, validEntries)
		if err != nil {
			return err
		}
		return nil
	}

	for _, entry := range validEntries {
		fmt.Printf("%s,%d\n", entry.Owner.Hex(), entry.TokenID)
	}

	return nil
}

func validateRegistrations(reg []registration, dc delegateLoader, gm ownerLoader) ([]validatedRegistration, error) {
	latestRecords := latestRegistrationsByToken(reg)

	debugPrintln("Found", len(latestRecords), "unique tokens")
	bar := progressbar.Default(int64(len(latestRecords)))

	validated := []validatedRegistration{}
	for _, entry := range latestRecords {
		bar.Add(1)

		valid, err := isValidEntry(dc, gm, entry)
		if err != nil {
			return nil, err
		}
		if valid {
			validated = append(validated, validatedRegistration{
				Owner:   entry.Owner,
				TokenID: entry.TokenID,
			})
		}
		if !valid {
			debugPrintln("Invalid entry:", entry.Owner, entry.TokenID)
		}
	}
	return validated, nil
}

// Check that the token is owned by the owner or one of the owner's vaults
func isValidEntry(dc delegateLoader, gm ownerLoader, entry registration) (bool, error) {
	vaults, err := getVaultWalletsForEntry(dc, entry)
	if err != nil {
		return false, err
	}

	owner, err := gm.OwnerOf(nil, big.NewInt(int64(entry.TokenID)))
	if err != nil {
		return false, err
	}

	// If any of the vaults contains the token, then the entry is valid
	return slices.Contains(vaults, owner), nil
}

func getVaultWalletsForEntry(dc delegateLoader, entry registration) ([]common.Address, error) {
	delegations, err := dc.GetDelegationsByDelegate(nil, entry.Owner)
	if err != nil {
		return nil, err
	}

	// If there is no delegation, there is no vault, so just return the owner
	if len(delegations) == 0 {
		return []common.Address{entry.Owner}, nil
	}

	// Keep a set of vault wallets.
	// Start off the the registered wallet, as this may be a hot wallet which holds the token
	// even though the delegation exists.
	vaultWallets := map[common.Address]bool{
		entry.Owner: true,
	}

	addDelegatesToSet := func(delegates []IDelegationRegistryDelegationInfo) {
		for _, d := range delegates {
			vaultWallets[d.Vault] = true
		}
	}

	// If there are delegations, find the primary delegate in order of specificity
	// starting at the token level and ending at the wallet level
	addDelegatesToSet(filterDelegations(delegations, func(d IDelegationRegistryDelegationInfo) bool {
		return d.Contract.Hex() == daoContractAddress.Hex() && int(d.TokenId.Int64()) == entry.TokenID
	}))

	// If there is no token level delegation, check for collection level delegation
	addDelegatesToSet(filterDelegations(delegations, func(d IDelegationRegistryDelegationInfo) bool {
		return d.Contract.Hex() == daoContractAddress.Hex() && int(d.TokenId.Int64()) == 0
	}))

	// If there is no collection level delegation, check for wallet level delegation
	addDelegatesToSet(filterDelegations(delegations, func(d IDelegationRegistryDelegationInfo) bool {
		return d.Contract.Hex() == zeroAddress.Hex()
	}))

	// Convert the set to a slice
	vaults := []common.Address{}
	for vault := range vaultWallets {
		vaults = append(vaults, vault)
	}

	return vaults, nil
}

func filterDelegations(delegations []IDelegationRegistryDelegationInfo, filter func(d IDelegationRegistryDelegationInfo) bool) []IDelegationRegistryDelegationInfo {
	filtered := []IDelegationRegistryDelegationInfo{}
	for _, d := range delegations {
		if filter(d) {
			filtered = append(filtered, d)
		}
	}
	return filtered
}

func csvLineToRegistration(line []string) registration {
	// Parse the timestamp
	timestamp, err := time.Parse(time.RFC3339, line[3])
	if err != nil {
		log.Fatal(err)
	}

	// Parse the owner address
	owner := common.HexToAddress(line[1])

	col, _ := strconv.Atoi(line[0])
	token, _ := strconv.Atoi(line[2])
	oneOfOne := token >= 870

	return registration{
		CollectionID: col,
		Owner:        owner,
		TokenID:      token,
		IsOneOfOne:   oneOfOne,
		Timestamp:    timestamp,
	}
}

// Find the *latest* entry for each token.
// This removes double entries and handles the case where people entered, sold, and the new owner entered.
func latestRegistrationsByToken(records []registration) []registration {
	idsToEntries := map[int]registration{}

	// Loop through the addresses in the CSV file
	for _, record := range records {
		// Check if entry already exists
		existing, found := idsToEntries[record.TokenID]
		if !found {
			idsToEntries[record.TokenID] = record
			continue
		}
		// Check if the timestamp is newer than the existing entry
		if record.Timestamp.After(existing.Timestamp) {
			idsToEntries[record.TokenID] = record
		}
	}

	// Convert the map to a slice
	latestRegistrations := []registration{}
	for _, entry := range idsToEntries {
		latestRegistrations = append(latestRegistrations, entry)
	}

	return latestRegistrations
}

func writeCSV(file string, validated []validatedRegistration) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, entry := range validated {
		if err := w.Write([]string{entry.Owner.Hex(), strconv.Itoa(entry.TokenID)}); err != nil {
			return err
		}
	}

	return nil
}
