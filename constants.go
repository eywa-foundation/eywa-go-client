package eywaclient

import (
	"context"
	"time"

	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

const ADDRESS_PREFIX = "cosmos"

func createClient(ctx context.Context, nodeAddress string) (cosmosclient.Client, error) {
	return cosmosclient.New(
		ctx,
		cosmosclient.WithAddressPrefix(ADDRESS_PREFIX),
		cosmosclient.WithNodeAddress(nodeAddress))
}

func getAccount(client cosmosclient.Client, accountName, mnemonic string) (cosmosaccount.Account, string, error) {
	var addr string

	account, err := client.AccountRegistry.Import(accountName, mnemonic, "")
	// Get account from the keyring
	if err != nil {
		return account, addr, err
	}
	addr, err = account.Address(ADDRESS_PREFIX)
	if err != nil {
		return account, addr, err
	}

	return account, addr, nil
}

// get unix timestamp in uint64
func getTimestamp() uint64 {
	return uint64(time.Now().Unix())
}
