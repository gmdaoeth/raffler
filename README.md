# raffle

1. Download CSV from website
1. Remove header row
1. Run `go run ./... gen-list -input file.csv -output verified.csv`
1. Choose a future block (e.g. 1 hour away), announce we'll use the "gas used" from that block for the raffle.
1. When the block happens, run: `go run ./... raffle -file verified-registrations.csv -seed <seed> > raffled.csv`
1. Use `cat raffled.csv | head -n <num>` to produce a list of N winners.
1. Upload the list to the `raffle-winners` repo.
1. Add 1/1 token holders to raffle output.
1. Use the allowlist generator script in `studio/ethereum` to get signatures.
1. Add signature allowlist file to web repo.
