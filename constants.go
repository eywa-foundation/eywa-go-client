package eywaclient

import (
	"context"
	"time"

	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

const ADDRESS_PREFIX = "cosmos"
const NODE_ADDRESS = "http://localhost:26657"

func createClient(ctx context.Context) (cosmosclient.Client, error) {
	return cosmosclient.New(
		ctx,
		cosmosclient.WithAddressPrefix(ADDRESS_PREFIX),
		cosmosclient.WithNodeAddress(NODE_ADDRESS),
	)
}

func getAccount(client cosmosclient.Client, accountName string) (cosmosaccount.Account, string, error) {
	var addr string

	// Get account from the keyring
	account, err := client.Account(accountName)
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
