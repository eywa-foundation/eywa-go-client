package eywaclient

import (
	"context"
	"fmt"
	"log"

	// Importing the general purpose Cosmos blockchain client
	"github.com/ignite/cli/ignite/pkg/cosmosclient"

	// Importing the types package of your blog blockchain
	"github.com/eywa-foundation/eywaclient/types"
)

func CreateChatTx(accountName, roomID, from, to, message string) error {
	ctx := context.Background()
	client, err := createClient(ctx)
	if err != nil {
		return err
	}

	relayAccount, _, err := getAccount(client, accountName)
	if err != nil {
		return err
	}

	msg := &types.MsgCreateChat{
		Creator:  from,
		RoomId:   roomID,
		Receiver: to,
		Message:  message,
		Time:     getTimestamp(),
	}
	txResp, err := client.BroadcastTx(ctx, relayAccount, msg)
	if err != nil {
		return err
	}
	log.Println("MsgCreateChat ->")
	log.Println(txResp)

	return nil
}

func PostRegister() {
	ctx := context.Background()

	// Create a Cosmos client instance
	client, err := cosmosclient.New(
		ctx,
		cosmosclient.WithAddressPrefix(ADDRESS_PREFIX),
		cosmosclient.WithNodeAddress(NODE_ADDRESS),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Account `alice` was initialized during `ignite chain serve`
	accountName := "bob"

	// Get account from the keyring
	account, err := client.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := account.Address(ADDRESS_PREFIX)
	if err != nil {
		log.Fatal(err)
	}

	// Define a message to create a post
	msg := &types.MsgRegisterUser{
		Creator: addr,
		Pubkey:  "pubkey",
	}
	fmt.Println("Addr: ", addr)

	// Broadcast a transaction from account `alice` with the message
	// to create a post store response in txResp
	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		log.Fatal(err)
	}

	// Print response from broadcasting a transaction
	fmt.Print("MsgCreateUser:\n\n")
	fmt.Println(txResp)

	// Instantiate a query client for your `blog` blockchain
	queryClient := types.NewQueryClient(client.Context())

	// Query the blockchain using the client's `PostAll` method
	// to get all posts store all posts in queryResp
	queryResp, err := queryClient.GetUser(ctx, &types.QueryGetUserRequest{Submitter: addr})
	if err != nil {
		log.Fatal(err)
	}

	// Print response from querying all the posts
	fmt.Print("\n\nAll user:\n\n")
	fmt.Println(queryResp)
}
