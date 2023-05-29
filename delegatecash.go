// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IDelegationRegistryContractDelegation is an auto generated low-level Go binding around an user-defined struct.
type IDelegationRegistryContractDelegation struct {
	Contract common.Address
	Delegate common.Address
}

// IDelegationRegistryDelegationInfo is an auto generated low-level Go binding around an user-defined struct.
type IDelegationRegistryDelegationInfo struct {
	Type     uint8
	Vault    common.Address
	Delegate common.Address
	Contract common.Address
	TokenId  *big.Int
}

// IDelegationRegistryTokenDelegation is an auto generated low-level Go binding around an user-defined struct.
type IDelegationRegistryTokenDelegation struct {
	Contract common.Address
	TokenId  *big.Int
	Delegate common.Address
}

// DelegateCashMetaData contains all meta data concerning the DelegateCash contract.
var DelegateCashMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"DelegateForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"DelegateForContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"DelegateForToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"RevokeAllDelegates\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"RevokeDelegate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"checkDelegateForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"}],\"name\":\"checkDelegateForContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"checkDelegateForToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"delegateForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"delegateForContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"delegateForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"getContractLevelDelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"internalType\":\"structIDelegationRegistry.ContractDelegation[]\",\"name\":\"contractDelegations\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"getDelegatesForAll\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"delegates\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"}],\"name\":\"getDelegatesForContract\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"delegates\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getDelegatesForToken\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"delegates\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"getDelegationsByDelegate\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIDelegationRegistry.DelegationType\",\"name\":\"type_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structIDelegationRegistry.DelegationInfo[]\",\"name\":\"info\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"getTokenLevelDelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"internalType\":\"structIDelegationRegistry.TokenDelegation[]\",\"name\":\"tokenDelegations\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeAllDelegates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"revokeDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"revokeSelf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DelegateCashABI is the input ABI used to generate the binding from.
// Deprecated: Use DelegateCashMetaData.ABI instead.
var DelegateCashABI = DelegateCashMetaData.ABI

// DelegateCash is an auto generated Go binding around an Ethereum contract.
type DelegateCash struct {
	DelegateCashCaller     // Read-only binding to the contract
	DelegateCashTransactor // Write-only binding to the contract
	DelegateCashFilterer   // Log filterer for contract events
}

// DelegateCashCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegateCashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegateCashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegateCashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegateCashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegateCashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegateCashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegateCashSession struct {
	Contract     *DelegateCash     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelegateCashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegateCashCallerSession struct {
	Contract *DelegateCashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DelegateCashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegateCashTransactorSession struct {
	Contract     *DelegateCashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DelegateCashRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegateCashRaw struct {
	Contract *DelegateCash // Generic contract binding to access the raw methods on
}

// DelegateCashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegateCashCallerRaw struct {
	Contract *DelegateCashCaller // Generic read-only contract binding to access the raw methods on
}

// DelegateCashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegateCashTransactorRaw struct {
	Contract *DelegateCashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegateCash creates a new instance of DelegateCash, bound to a specific deployed contract.
func NewDelegateCash(address common.Address, backend bind.ContractBackend) (*DelegateCash, error) {
	contract, err := bindDelegateCash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelegateCash{DelegateCashCaller: DelegateCashCaller{contract: contract}, DelegateCashTransactor: DelegateCashTransactor{contract: contract}, DelegateCashFilterer: DelegateCashFilterer{contract: contract}}, nil
}

// NewDelegateCashCaller creates a new read-only instance of DelegateCash, bound to a specific deployed contract.
func NewDelegateCashCaller(address common.Address, caller bind.ContractCaller) (*DelegateCashCaller, error) {
	contract, err := bindDelegateCash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegateCashCaller{contract: contract}, nil
}

// NewDelegateCashTransactor creates a new write-only instance of DelegateCash, bound to a specific deployed contract.
func NewDelegateCashTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegateCashTransactor, error) {
	contract, err := bindDelegateCash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegateCashTransactor{contract: contract}, nil
}

// NewDelegateCashFilterer creates a new log filterer instance of DelegateCash, bound to a specific deployed contract.
func NewDelegateCashFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegateCashFilterer, error) {
	contract, err := bindDelegateCash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegateCashFilterer{contract: contract}, nil
}

// bindDelegateCash binds a generic wrapper to an already deployed contract.
func bindDelegateCash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DelegateCashABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegateCash *DelegateCashRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelegateCash.Contract.DelegateCashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegateCash *DelegateCashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegateCash.Contract.DelegateCashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegateCash *DelegateCashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegateCash.Contract.DelegateCashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegateCash *DelegateCashCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelegateCash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegateCash *DelegateCashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegateCash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegateCash *DelegateCashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegateCash.Contract.contract.Transact(opts, method, params...)
}

// CheckDelegateForAll is a free data retrieval call binding the contract method 0x9c395bc2.
//
// Solidity: function checkDelegateForAll(address delegate, address vault) view returns(bool)
func (_DelegateCash *DelegateCashCaller) CheckDelegateForAll(opts *bind.CallOpts, delegate common.Address, vault common.Address) (bool, error) {
	var out []interface{}
	err := _DelegateCash.contract.Call(opts, &out, "checkDelegateForAll", delegate, vault)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDelegateForAll is a free data retrieval call binding the contract method 0x9c395bc2.
//
// Solidity: function checkDelegateForAll(address delegate, address vault) view returns(bool)
func (_DelegateCash *DelegateCashSession) CheckDelegateForAll(delegate common.Address, vault common.Address) (bool, error) {
	return _DelegateCash.Contract.CheckDelegateForAll(&_DelegateCash.CallOpts, delegate, vault)
}

// CheckDelegateForAll is a free data retrieval call binding the contract method 0x9c395bc2.
//
// Solidity: function checkDelegateForAll(address delegate, address vault) view returns(bool)
func (_DelegateCash *DelegateCashCallerSession) CheckDelegateForAll(delegate common.Address, vault common.Address) (bool, error) {
	return _DelegateCash.Contract.CheckDelegateForAll(&_DelegateCash.CallOpts, delegate, vault)
}

// CheckDelegateForContract is a free data retrieval call binding the contract method 0x90c9a2d0.
//
// Solidity: function checkDelegateForContract(address delegate, address vault, address contract_) view returns(bool)
func (_DelegateCash *DelegateCashCaller) CheckDelegateForContract(opts *bind.CallOpts, delegate common.Address, vault common.Address, contract_ common.Address) (bool, error) {
	var out []interface{}
	err := _DelegateCash.contract.Call(opts, &out, "checkDelegateForContract", delegate, vault, contract_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDelegateForContract is a free data retrieval call binding the contract method 0x90c9a2d0.
//
// Solidity: function checkDelegateForContract(address delegate, address vault, address contract_) view returns(bool)
func (_DelegateCash *DelegateCashSession) CheckDelegateForContract(delegate common.Address, vault common.Address, contract_ common.Address) (bool, error) {
	return _DelegateCash.Contract.CheckDelegateForContract(&_DelegateCash.CallOpts, delegate, vault, contract_)
}

// CheckDelegateForContract is a free data retrieval call binding the contract method 0x90c9a2d0.
//
// Solidity: function checkDelegateForContract(address delegate, address vault, address contract_) view returns(bool)
func (_DelegateCash *DelegateCashCallerSession) CheckDelegateForContract(delegate common.Address, vault common.Address, contract_ common.Address) (bool, error) {
	return _DelegateCash.Contract.CheckDelegateForContract(&_DelegateCash.CallOpts, delegate, vault, contract_)
}

// CheckDelegateForToken is a free data retrieval call binding the contract method 0xaba69cf8.
//
// Solidity: function checkDelegateForToken(address delegate, address vault, address contract_, uint256 tokenId) view returns(bool)
func (_DelegateCash *DelegateCashCaller) CheckDelegateForToken(opts *bind.CallOpts, delegate common.Address, vault common.Address, contract_ common.Address, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _DelegateCash.contract.Call(opts, &out, "checkDelegateForToken", delegate, vault, contract_, tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDelegateForToken is a free data retrieval call binding the contract method 0xaba69cf8.
//
// Solidity: function checkDelegateForToken(address delegate, address vault, address contract_, uint256 tokenId) view returns(bool)
func (_DelegateCash *DelegateCashSession) CheckDelegateForToken(delegate common.Address, vault common.Address, contract_ common.Address, tokenId *big.Int) (bool, error) {
	return _DelegateCash.Contract.CheckDelegateForToken(&_DelegateCash.CallOpts, delegate, vault, contract_, tokenId)
}

// CheckDelegateForToken is a free data retrieval call binding the contract method 0xaba69cf8.
//
// Solidity: function checkDelegateForToken(address delegate, address vault, address contract_, uint256 tokenId) view returns(bool)
func (_DelegateCash *DelegateCashCallerSession) CheckDelegateForToken(delegate common.Address, vault common.Address, contract_ common.Address, tokenId *big.Int) (bool, error) {
	return _DelegateCash.Contract.CheckDelegateForToken(&_DelegateCash.CallOpts, delegate, vault, contract_, tokenId)
}

// GetContractLevelDelegations is a free data retrieval call binding the contract method 0xf956cf94.
//
// Solidity: function getContractLevelDelegations(address vault) view returns((address,address)[] contractDelegations)
func (_DelegateCash *DelegateCashCaller) GetContractLevelDelegations(opts *bind.CallOpts, vault common.Address) ([]IDelegationRegistryContractDelegation, error) {
	var out []interface{}
	err := _DelegateCash.contract.Call(opts, &out, "getContractLevelDelegations", vault)

	if err != nil {
		return *new([]IDelegationRegistryContractDelegation), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDelegationRegistryContractDelegation)).(*[]IDelegationRegistryContractDelegation)

	return out0, err

}

// GetContractLevelDelegations is a free data retrieval call binding the contract method 0xf956cf94.
//
// Solidity: function getContractLevelDelegations(address vault) view returns((address,address)[] contractDelegations)
func (_DelegateCash *DelegateCashSession) GetContractLevelDelegations(vault common.Address) ([]IDelegationRegistryContractDelegation, error) {
	return _DelegateCash.Contract.GetContractLevelDelegations(&_DelegateCash.CallOpts, vault)
}

// GetContractLevelDelegations is a free data retrieval call binding the contract method 0xf956cf94.
//
// Solidity: function getContractLevelDelegations(address vault) view returns((address,address)[] contractDelegations)
func (_DelegateCash *DelegateCashCallerSession) GetContractLevelDelegations(vault common.Address) ([]IDelegationRegistryContractDelegation, error) {
	return _DelegateCash.Contract.GetContractLevelDelegations(&_DelegateCash.CallOpts, vault)
}

// GetDelegatesForAll is a free data retrieval call binding the contract method 0x1b61f675.
//
// Solidity: function getDelegatesForAll(address vault) view returns(address[] delegates)
func (_DelegateCash *DelegateCashCaller) GetDelegatesForAll(opts *bind.CallOpts, vault common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _DelegateCash.contract.Call(opts, &out, "getDelegatesForAll", vault)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDelegatesForAll is a free data retrieval call binding the contract method 0x1b61f675.
//
// Solidity: function getDelegatesForAll(address vault) view returns(address[] delegates)
func (_DelegateCash *DelegateCashSession) GetDelegatesForAll(vault common.Address) ([]common.Address, error) {
	return _DelegateCash.Contract.GetDelegatesForAll(&_DelegateCash.CallOpts, vault)
}

// GetDelegatesForAll is a free data retrieval call binding the contract method 0x1b61f675.
//
// Solidity: function getDelegatesForAll(address vault) view returns(address[] delegates)
func (_DelegateCash *DelegateCashCallerSession) GetDelegatesForAll(vault common.Address) ([]common.Address, error) {
	return _DelegateCash.Contract.GetDelegatesForAll(&_DelegateCash.CallOpts, vault)
}

// GetDelegatesForContract is a free data retrieval call binding the contract method 0xed4b878e.
//
// Solidity: function getDelegatesForContract(address vault, address contract_) view returns(address[] delegates)
func (_DelegateCash *DelegateCashCaller) GetDelegatesForContract(opts *bind.CallOpts, vault common.Address, contract_ common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _DelegateCash.contract.Call(opts, &out, "getDelegatesForContract", vault, contract_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDelegatesForContract is a free data retrieval call binding the contract method 0xed4b878e.
//
// Solidity: function getDelegatesForContract(address vault, address contract_) view returns(address[] delegates)
func (_DelegateCash *DelegateCashSession) GetDelegatesForContract(vault common.Address, contract_ common.Address) ([]common.Address, error) {
	return _DelegateCash.Contract.GetDelegatesForContract(&_DelegateCash.CallOpts, vault, contract_)
}

// GetDelegatesForContract is a free data retrieval call binding the contract method 0xed4b878e.
//
// Solidity: function getDelegatesForContract(address vault, address contract_) view returns(address[] delegates)
func (_DelegateCash *DelegateCashCallerSession) GetDelegatesForContract(vault common.Address, contract_ common.Address) ([]common.Address, error) {
	return _DelegateCash.Contract.GetDelegatesForContract(&_DelegateCash.CallOpts, vault, contract_)
}

// GetDelegatesForToken is a free data retrieval call binding the contract method 0x1221156b.
//
// Solidity: function getDelegatesForToken(address vault, address contract_, uint256 tokenId) view returns(address[] delegates)
func (_DelegateCash *DelegateCashCaller) GetDelegatesForToken(opts *bind.CallOpts, vault common.Address, contract_ common.Address, tokenId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _DelegateCash.contract.Call(opts, &out, "getDelegatesForToken", vault, contract_, tokenId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDelegatesForToken is a free data retrieval call binding the contract method 0x1221156b.
//
// Solidity: function getDelegatesForToken(address vault, address contract_, uint256 tokenId) view returns(address[] delegates)
func (_DelegateCash *DelegateCashSession) GetDelegatesForToken(vault common.Address, contract_ common.Address, tokenId *big.Int) ([]common.Address, error) {
	return _DelegateCash.Contract.GetDelegatesForToken(&_DelegateCash.CallOpts, vault, contract_, tokenId)
}

// GetDelegatesForToken is a free data retrieval call binding the contract method 0x1221156b.
//
// Solidity: function getDelegatesForToken(address vault, address contract_, uint256 tokenId) view returns(address[] delegates)
func (_DelegateCash *DelegateCashCallerSession) GetDelegatesForToken(vault common.Address, contract_ common.Address, tokenId *big.Int) ([]common.Address, error) {
	return _DelegateCash.Contract.GetDelegatesForToken(&_DelegateCash.CallOpts, vault, contract_, tokenId)
}

// GetDelegationsByDelegate is a free data retrieval call binding the contract method 0x4fc69282.
//
// Solidity: function getDelegationsByDelegate(address delegate) view returns((uint8,address,address,address,uint256)[] info)
func (_DelegateCash *DelegateCashCaller) GetDelegationsByDelegate(opts *bind.CallOpts, delegate common.Address) ([]IDelegationRegistryDelegationInfo, error) {
	var out []interface{}
	err := _DelegateCash.contract.Call(opts, &out, "getDelegationsByDelegate", delegate)

	if err != nil {
		return *new([]IDelegationRegistryDelegationInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDelegationRegistryDelegationInfo)).(*[]IDelegationRegistryDelegationInfo)

	return out0, err

}

// GetDelegationsByDelegate is a free data retrieval call binding the contract method 0x4fc69282.
//
// Solidity: function getDelegationsByDelegate(address delegate) view returns((uint8,address,address,address,uint256)[] info)
func (_DelegateCash *DelegateCashSession) GetDelegationsByDelegate(delegate common.Address) ([]IDelegationRegistryDelegationInfo, error) {
	return _DelegateCash.Contract.GetDelegationsByDelegate(&_DelegateCash.CallOpts, delegate)
}

// GetDelegationsByDelegate is a free data retrieval call binding the contract method 0x4fc69282.
//
// Solidity: function getDelegationsByDelegate(address delegate) view returns((uint8,address,address,address,uint256)[] info)
func (_DelegateCash *DelegateCashCallerSession) GetDelegationsByDelegate(delegate common.Address) ([]IDelegationRegistryDelegationInfo, error) {
	return _DelegateCash.Contract.GetDelegationsByDelegate(&_DelegateCash.CallOpts, delegate)
}

// GetTokenLevelDelegations is a free data retrieval call binding the contract method 0x6f007d87.
//
// Solidity: function getTokenLevelDelegations(address vault) view returns((address,uint256,address)[] tokenDelegations)
func (_DelegateCash *DelegateCashCaller) GetTokenLevelDelegations(opts *bind.CallOpts, vault common.Address) ([]IDelegationRegistryTokenDelegation, error) {
	var out []interface{}
	err := _DelegateCash.contract.Call(opts, &out, "getTokenLevelDelegations", vault)

	if err != nil {
		return *new([]IDelegationRegistryTokenDelegation), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDelegationRegistryTokenDelegation)).(*[]IDelegationRegistryTokenDelegation)

	return out0, err

}

// GetTokenLevelDelegations is a free data retrieval call binding the contract method 0x6f007d87.
//
// Solidity: function getTokenLevelDelegations(address vault) view returns((address,uint256,address)[] tokenDelegations)
func (_DelegateCash *DelegateCashSession) GetTokenLevelDelegations(vault common.Address) ([]IDelegationRegistryTokenDelegation, error) {
	return _DelegateCash.Contract.GetTokenLevelDelegations(&_DelegateCash.CallOpts, vault)
}

// GetTokenLevelDelegations is a free data retrieval call binding the contract method 0x6f007d87.
//
// Solidity: function getTokenLevelDelegations(address vault) view returns((address,uint256,address)[] tokenDelegations)
func (_DelegateCash *DelegateCashCallerSession) GetTokenLevelDelegations(vault common.Address) ([]IDelegationRegistryTokenDelegation, error) {
	return _DelegateCash.Contract.GetTokenLevelDelegations(&_DelegateCash.CallOpts, vault)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DelegateCash *DelegateCashCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DelegateCash.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DelegateCash *DelegateCashSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DelegateCash.Contract.SupportsInterface(&_DelegateCash.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DelegateCash *DelegateCashCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DelegateCash.Contract.SupportsInterface(&_DelegateCash.CallOpts, interfaceId)
}

// DelegateForAll is a paid mutator transaction binding the contract method 0x685ee3e8.
//
// Solidity: function delegateForAll(address delegate, bool value) returns()
func (_DelegateCash *DelegateCashTransactor) DelegateForAll(opts *bind.TransactOpts, delegate common.Address, value bool) (*types.Transaction, error) {
	return _DelegateCash.contract.Transact(opts, "delegateForAll", delegate, value)
}

// DelegateForAll is a paid mutator transaction binding the contract method 0x685ee3e8.
//
// Solidity: function delegateForAll(address delegate, bool value) returns()
func (_DelegateCash *DelegateCashSession) DelegateForAll(delegate common.Address, value bool) (*types.Transaction, error) {
	return _DelegateCash.Contract.DelegateForAll(&_DelegateCash.TransactOpts, delegate, value)
}

// DelegateForAll is a paid mutator transaction binding the contract method 0x685ee3e8.
//
// Solidity: function delegateForAll(address delegate, bool value) returns()
func (_DelegateCash *DelegateCashTransactorSession) DelegateForAll(delegate common.Address, value bool) (*types.Transaction, error) {
	return _DelegateCash.Contract.DelegateForAll(&_DelegateCash.TransactOpts, delegate, value)
}

// DelegateForContract is a paid mutator transaction binding the contract method 0x49c95d29.
//
// Solidity: function delegateForContract(address delegate, address contract_, bool value) returns()
func (_DelegateCash *DelegateCashTransactor) DelegateForContract(opts *bind.TransactOpts, delegate common.Address, contract_ common.Address, value bool) (*types.Transaction, error) {
	return _DelegateCash.contract.Transact(opts, "delegateForContract", delegate, contract_, value)
}

// DelegateForContract is a paid mutator transaction binding the contract method 0x49c95d29.
//
// Solidity: function delegateForContract(address delegate, address contract_, bool value) returns()
func (_DelegateCash *DelegateCashSession) DelegateForContract(delegate common.Address, contract_ common.Address, value bool) (*types.Transaction, error) {
	return _DelegateCash.Contract.DelegateForContract(&_DelegateCash.TransactOpts, delegate, contract_, value)
}

// DelegateForContract is a paid mutator transaction binding the contract method 0x49c95d29.
//
// Solidity: function delegateForContract(address delegate, address contract_, bool value) returns()
func (_DelegateCash *DelegateCashTransactorSession) DelegateForContract(delegate common.Address, contract_ common.Address, value bool) (*types.Transaction, error) {
	return _DelegateCash.Contract.DelegateForContract(&_DelegateCash.TransactOpts, delegate, contract_, value)
}

// DelegateForToken is a paid mutator transaction binding the contract method 0x537a5c3d.
//
// Solidity: function delegateForToken(address delegate, address contract_, uint256 tokenId, bool value) returns()
func (_DelegateCash *DelegateCashTransactor) DelegateForToken(opts *bind.TransactOpts, delegate common.Address, contract_ common.Address, tokenId *big.Int, value bool) (*types.Transaction, error) {
	return _DelegateCash.contract.Transact(opts, "delegateForToken", delegate, contract_, tokenId, value)
}

// DelegateForToken is a paid mutator transaction binding the contract method 0x537a5c3d.
//
// Solidity: function delegateForToken(address delegate, address contract_, uint256 tokenId, bool value) returns()
func (_DelegateCash *DelegateCashSession) DelegateForToken(delegate common.Address, contract_ common.Address, tokenId *big.Int, value bool) (*types.Transaction, error) {
	return _DelegateCash.Contract.DelegateForToken(&_DelegateCash.TransactOpts, delegate, contract_, tokenId, value)
}

// DelegateForToken is a paid mutator transaction binding the contract method 0x537a5c3d.
//
// Solidity: function delegateForToken(address delegate, address contract_, uint256 tokenId, bool value) returns()
func (_DelegateCash *DelegateCashTransactorSession) DelegateForToken(delegate common.Address, contract_ common.Address, tokenId *big.Int, value bool) (*types.Transaction, error) {
	return _DelegateCash.Contract.DelegateForToken(&_DelegateCash.TransactOpts, delegate, contract_, tokenId, value)
}

// RevokeAllDelegates is a paid mutator transaction binding the contract method 0x36137872.
//
// Solidity: function revokeAllDelegates() returns()
func (_DelegateCash *DelegateCashTransactor) RevokeAllDelegates(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegateCash.contract.Transact(opts, "revokeAllDelegates")
}

// RevokeAllDelegates is a paid mutator transaction binding the contract method 0x36137872.
//
// Solidity: function revokeAllDelegates() returns()
func (_DelegateCash *DelegateCashSession) RevokeAllDelegates() (*types.Transaction, error) {
	return _DelegateCash.Contract.RevokeAllDelegates(&_DelegateCash.TransactOpts)
}

// RevokeAllDelegates is a paid mutator transaction binding the contract method 0x36137872.
//
// Solidity: function revokeAllDelegates() returns()
func (_DelegateCash *DelegateCashTransactorSession) RevokeAllDelegates() (*types.Transaction, error) {
	return _DelegateCash.Contract.RevokeAllDelegates(&_DelegateCash.TransactOpts)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0xfa352c00.
//
// Solidity: function revokeDelegate(address delegate) returns()
func (_DelegateCash *DelegateCashTransactor) RevokeDelegate(opts *bind.TransactOpts, delegate common.Address) (*types.Transaction, error) {
	return _DelegateCash.contract.Transact(opts, "revokeDelegate", delegate)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0xfa352c00.
//
// Solidity: function revokeDelegate(address delegate) returns()
func (_DelegateCash *DelegateCashSession) RevokeDelegate(delegate common.Address) (*types.Transaction, error) {
	return _DelegateCash.Contract.RevokeDelegate(&_DelegateCash.TransactOpts, delegate)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0xfa352c00.
//
// Solidity: function revokeDelegate(address delegate) returns()
func (_DelegateCash *DelegateCashTransactorSession) RevokeDelegate(delegate common.Address) (*types.Transaction, error) {
	return _DelegateCash.Contract.RevokeDelegate(&_DelegateCash.TransactOpts, delegate)
}

// RevokeSelf is a paid mutator transaction binding the contract method 0x219044b0.
//
// Solidity: function revokeSelf(address vault) returns()
func (_DelegateCash *DelegateCashTransactor) RevokeSelf(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _DelegateCash.contract.Transact(opts, "revokeSelf", vault)
}

// RevokeSelf is a paid mutator transaction binding the contract method 0x219044b0.
//
// Solidity: function revokeSelf(address vault) returns()
func (_DelegateCash *DelegateCashSession) RevokeSelf(vault common.Address) (*types.Transaction, error) {
	return _DelegateCash.Contract.RevokeSelf(&_DelegateCash.TransactOpts, vault)
}

// RevokeSelf is a paid mutator transaction binding the contract method 0x219044b0.
//
// Solidity: function revokeSelf(address vault) returns()
func (_DelegateCash *DelegateCashTransactorSession) RevokeSelf(vault common.Address) (*types.Transaction, error) {
	return _DelegateCash.Contract.RevokeSelf(&_DelegateCash.TransactOpts, vault)
}

// DelegateCashDelegateForAllIterator is returned from FilterDelegateForAll and is used to iterate over the raw logs and unpacked data for DelegateForAll events raised by the DelegateCash contract.
type DelegateCashDelegateForAllIterator struct {
	Event *DelegateCashDelegateForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegateCashDelegateForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegateCashDelegateForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegateCashDelegateForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegateCashDelegateForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegateCashDelegateForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegateCashDelegateForAll represents a DelegateForAll event raised by the DelegateCash contract.
type DelegateCashDelegateForAll struct {
	Vault    common.Address
	Delegate common.Address
	Value    bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegateForAll is a free log retrieval operation binding the contract event 0x58781eab4a0743ab1c285a238be846a235f06cdb5b968030573a635e5f8c92fa.
//
// Solidity: event DelegateForAll(address vault, address delegate, bool value)
func (_DelegateCash *DelegateCashFilterer) FilterDelegateForAll(opts *bind.FilterOpts) (*DelegateCashDelegateForAllIterator, error) {

	logs, sub, err := _DelegateCash.contract.FilterLogs(opts, "DelegateForAll")
	if err != nil {
		return nil, err
	}
	return &DelegateCashDelegateForAllIterator{contract: _DelegateCash.contract, event: "DelegateForAll", logs: logs, sub: sub}, nil
}

// WatchDelegateForAll is a free log subscription operation binding the contract event 0x58781eab4a0743ab1c285a238be846a235f06cdb5b968030573a635e5f8c92fa.
//
// Solidity: event DelegateForAll(address vault, address delegate, bool value)
func (_DelegateCash *DelegateCashFilterer) WatchDelegateForAll(opts *bind.WatchOpts, sink chan<- *DelegateCashDelegateForAll) (event.Subscription, error) {

	logs, sub, err := _DelegateCash.contract.WatchLogs(opts, "DelegateForAll")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegateCashDelegateForAll)
				if err := _DelegateCash.contract.UnpackLog(event, "DelegateForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateForAll is a log parse operation binding the contract event 0x58781eab4a0743ab1c285a238be846a235f06cdb5b968030573a635e5f8c92fa.
//
// Solidity: event DelegateForAll(address vault, address delegate, bool value)
func (_DelegateCash *DelegateCashFilterer) ParseDelegateForAll(log types.Log) (*DelegateCashDelegateForAll, error) {
	event := new(DelegateCashDelegateForAll)
	if err := _DelegateCash.contract.UnpackLog(event, "DelegateForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegateCashDelegateForContractIterator is returned from FilterDelegateForContract and is used to iterate over the raw logs and unpacked data for DelegateForContract events raised by the DelegateCash contract.
type DelegateCashDelegateForContractIterator struct {
	Event *DelegateCashDelegateForContract // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegateCashDelegateForContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegateCashDelegateForContract)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegateCashDelegateForContract)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegateCashDelegateForContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegateCashDelegateForContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegateCashDelegateForContract represents a DelegateForContract event raised by the DelegateCash contract.
type DelegateCashDelegateForContract struct {
	Vault    common.Address
	Delegate common.Address
	Contract common.Address
	Value    bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegateForContract is a free log retrieval operation binding the contract event 0x8d6b2f5255b8d815cc368855b2251146e003bf4e2fcccaec66145fff5c174b4f.
//
// Solidity: event DelegateForContract(address vault, address delegate, address contract_, bool value)
func (_DelegateCash *DelegateCashFilterer) FilterDelegateForContract(opts *bind.FilterOpts) (*DelegateCashDelegateForContractIterator, error) {

	logs, sub, err := _DelegateCash.contract.FilterLogs(opts, "DelegateForContract")
	if err != nil {
		return nil, err
	}
	return &DelegateCashDelegateForContractIterator{contract: _DelegateCash.contract, event: "DelegateForContract", logs: logs, sub: sub}, nil
}

// WatchDelegateForContract is a free log subscription operation binding the contract event 0x8d6b2f5255b8d815cc368855b2251146e003bf4e2fcccaec66145fff5c174b4f.
//
// Solidity: event DelegateForContract(address vault, address delegate, address contract_, bool value)
func (_DelegateCash *DelegateCashFilterer) WatchDelegateForContract(opts *bind.WatchOpts, sink chan<- *DelegateCashDelegateForContract) (event.Subscription, error) {

	logs, sub, err := _DelegateCash.contract.WatchLogs(opts, "DelegateForContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegateCashDelegateForContract)
				if err := _DelegateCash.contract.UnpackLog(event, "DelegateForContract", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateForContract is a log parse operation binding the contract event 0x8d6b2f5255b8d815cc368855b2251146e003bf4e2fcccaec66145fff5c174b4f.
//
// Solidity: event DelegateForContract(address vault, address delegate, address contract_, bool value)
func (_DelegateCash *DelegateCashFilterer) ParseDelegateForContract(log types.Log) (*DelegateCashDelegateForContract, error) {
	event := new(DelegateCashDelegateForContract)
	if err := _DelegateCash.contract.UnpackLog(event, "DelegateForContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegateCashDelegateForTokenIterator is returned from FilterDelegateForToken and is used to iterate over the raw logs and unpacked data for DelegateForToken events raised by the DelegateCash contract.
type DelegateCashDelegateForTokenIterator struct {
	Event *DelegateCashDelegateForToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegateCashDelegateForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegateCashDelegateForToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegateCashDelegateForToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegateCashDelegateForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegateCashDelegateForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegateCashDelegateForToken represents a DelegateForToken event raised by the DelegateCash contract.
type DelegateCashDelegateForToken struct {
	Vault    common.Address
	Delegate common.Address
	Contract common.Address
	TokenId  *big.Int
	Value    bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegateForToken is a free log retrieval operation binding the contract event 0xe89c6ba1e8957285aed22618f52aa1dcb9d5bb64e1533d8b55136c72fcf5aa5d.
//
// Solidity: event DelegateForToken(address vault, address delegate, address contract_, uint256 tokenId, bool value)
func (_DelegateCash *DelegateCashFilterer) FilterDelegateForToken(opts *bind.FilterOpts) (*DelegateCashDelegateForTokenIterator, error) {

	logs, sub, err := _DelegateCash.contract.FilterLogs(opts, "DelegateForToken")
	if err != nil {
		return nil, err
	}
	return &DelegateCashDelegateForTokenIterator{contract: _DelegateCash.contract, event: "DelegateForToken", logs: logs, sub: sub}, nil
}

// WatchDelegateForToken is a free log subscription operation binding the contract event 0xe89c6ba1e8957285aed22618f52aa1dcb9d5bb64e1533d8b55136c72fcf5aa5d.
//
// Solidity: event DelegateForToken(address vault, address delegate, address contract_, uint256 tokenId, bool value)
func (_DelegateCash *DelegateCashFilterer) WatchDelegateForToken(opts *bind.WatchOpts, sink chan<- *DelegateCashDelegateForToken) (event.Subscription, error) {

	logs, sub, err := _DelegateCash.contract.WatchLogs(opts, "DelegateForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegateCashDelegateForToken)
				if err := _DelegateCash.contract.UnpackLog(event, "DelegateForToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateForToken is a log parse operation binding the contract event 0xe89c6ba1e8957285aed22618f52aa1dcb9d5bb64e1533d8b55136c72fcf5aa5d.
//
// Solidity: event DelegateForToken(address vault, address delegate, address contract_, uint256 tokenId, bool value)
func (_DelegateCash *DelegateCashFilterer) ParseDelegateForToken(log types.Log) (*DelegateCashDelegateForToken, error) {
	event := new(DelegateCashDelegateForToken)
	if err := _DelegateCash.contract.UnpackLog(event, "DelegateForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegateCashRevokeAllDelegatesIterator is returned from FilterRevokeAllDelegates and is used to iterate over the raw logs and unpacked data for RevokeAllDelegates events raised by the DelegateCash contract.
type DelegateCashRevokeAllDelegatesIterator struct {
	Event *DelegateCashRevokeAllDelegates // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegateCashRevokeAllDelegatesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegateCashRevokeAllDelegates)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegateCashRevokeAllDelegates)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegateCashRevokeAllDelegatesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegateCashRevokeAllDelegatesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegateCashRevokeAllDelegates represents a RevokeAllDelegates event raised by the DelegateCash contract.
type DelegateCashRevokeAllDelegates struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRevokeAllDelegates is a free log retrieval operation binding the contract event 0x32d74befd0b842e19694e3e3af46263e18bcce41352c8b600ff0002b49edf662.
//
// Solidity: event RevokeAllDelegates(address vault)
func (_DelegateCash *DelegateCashFilterer) FilterRevokeAllDelegates(opts *bind.FilterOpts) (*DelegateCashRevokeAllDelegatesIterator, error) {

	logs, sub, err := _DelegateCash.contract.FilterLogs(opts, "RevokeAllDelegates")
	if err != nil {
		return nil, err
	}
	return &DelegateCashRevokeAllDelegatesIterator{contract: _DelegateCash.contract, event: "RevokeAllDelegates", logs: logs, sub: sub}, nil
}

// WatchRevokeAllDelegates is a free log subscription operation binding the contract event 0x32d74befd0b842e19694e3e3af46263e18bcce41352c8b600ff0002b49edf662.
//
// Solidity: event RevokeAllDelegates(address vault)
func (_DelegateCash *DelegateCashFilterer) WatchRevokeAllDelegates(opts *bind.WatchOpts, sink chan<- *DelegateCashRevokeAllDelegates) (event.Subscription, error) {

	logs, sub, err := _DelegateCash.contract.WatchLogs(opts, "RevokeAllDelegates")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegateCashRevokeAllDelegates)
				if err := _DelegateCash.contract.UnpackLog(event, "RevokeAllDelegates", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokeAllDelegates is a log parse operation binding the contract event 0x32d74befd0b842e19694e3e3af46263e18bcce41352c8b600ff0002b49edf662.
//
// Solidity: event RevokeAllDelegates(address vault)
func (_DelegateCash *DelegateCashFilterer) ParseRevokeAllDelegates(log types.Log) (*DelegateCashRevokeAllDelegates, error) {
	event := new(DelegateCashRevokeAllDelegates)
	if err := _DelegateCash.contract.UnpackLog(event, "RevokeAllDelegates", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegateCashRevokeDelegateIterator is returned from FilterRevokeDelegate and is used to iterate over the raw logs and unpacked data for RevokeDelegate events raised by the DelegateCash contract.
type DelegateCashRevokeDelegateIterator struct {
	Event *DelegateCashRevokeDelegate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelegateCashRevokeDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegateCashRevokeDelegate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelegateCashRevokeDelegate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelegateCashRevokeDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegateCashRevokeDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegateCashRevokeDelegate represents a RevokeDelegate event raised by the DelegateCash contract.
type DelegateCashRevokeDelegate struct {
	Vault    common.Address
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRevokeDelegate is a free log retrieval operation binding the contract event 0x3e34a3ee53064fb79c0ee57448f03774a627a9270b0c41286efb7d8e32dcde93.
//
// Solidity: event RevokeDelegate(address vault, address delegate)
func (_DelegateCash *DelegateCashFilterer) FilterRevokeDelegate(opts *bind.FilterOpts) (*DelegateCashRevokeDelegateIterator, error) {

	logs, sub, err := _DelegateCash.contract.FilterLogs(opts, "RevokeDelegate")
	if err != nil {
		return nil, err
	}
	return &DelegateCashRevokeDelegateIterator{contract: _DelegateCash.contract, event: "RevokeDelegate", logs: logs, sub: sub}, nil
}

// WatchRevokeDelegate is a free log subscription operation binding the contract event 0x3e34a3ee53064fb79c0ee57448f03774a627a9270b0c41286efb7d8e32dcde93.
//
// Solidity: event RevokeDelegate(address vault, address delegate)
func (_DelegateCash *DelegateCashFilterer) WatchRevokeDelegate(opts *bind.WatchOpts, sink chan<- *DelegateCashRevokeDelegate) (event.Subscription, error) {

	logs, sub, err := _DelegateCash.contract.WatchLogs(opts, "RevokeDelegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegateCashRevokeDelegate)
				if err := _DelegateCash.contract.UnpackLog(event, "RevokeDelegate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevokeDelegate is a log parse operation binding the contract event 0x3e34a3ee53064fb79c0ee57448f03774a627a9270b0c41286efb7d8e32dcde93.
//
// Solidity: event RevokeDelegate(address vault, address delegate)
func (_DelegateCash *DelegateCashFilterer) ParseRevokeDelegate(log types.Log) (*DelegateCashRevokeDelegate, error) {
	event := new(DelegateCashRevokeDelegate)
	if err := _DelegateCash.contract.UnpackLog(event, "RevokeDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
