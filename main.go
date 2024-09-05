package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mrdhat/eth-txns/config"
	"github.com/mrdhat/eth-txns/entity"
	"github.com/mrdhat/eth-txns/jsonrpc"
	"github.com/mrdhat/eth-txns/node"
	"github.com/mrdhat/eth-txns/store"
)

func main() {
	configData := config.NewConfig()

	rpcClient := jsonrpc.NewJSONRPCClient(configData.NodeRPCUrl, &http.Client{})

	blockStore := store.NewBlockStore(store.StoreTypeMemory)
	addressSubscriptionStore := store.NewAddressSubscriptionStore(store.StoreTypeMemory)
	transactionStore := store.NewTransactionStore(store.StoreTypeMemory, addressSubscriptionStore)

	addresses := []string{"0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5", "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97", "0x8E1fDdf3b59BeA59a3B963Cc38719eb1dc7b6ee2"}

	for _, address := range addresses {
		addressSubscriptionStore.Save(entity.AddressSubscription{Address: address, IsActive: true})
	}

	listener := node.NewListener(configData.BlockDuration, rpcClient, blockStore, transactionStore)

	fmt.Println("Starting listener")
	stop := make(chan bool)

	// TODO: stop on signal
	go func() {
		time.Sleep(50 * time.Second)
		fmt.Println("Stopping listener")
		transactions := transactionStore.GetAll()
		fmt.Println("Transactions: ", transactions)
		stop <- true
	}()

	err := listener.Start(stop)
	if err != nil {
		fmt.Println("Listener stopped: ", err)
	}
}
