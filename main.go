package main

import (
	"net/http"
	"os"
	"time"

	"github.com/mrdhat/eth-txns/config"
	"github.com/mrdhat/eth-txns/entity"
	"github.com/mrdhat/eth-txns/jsonrpc"
	"github.com/mrdhat/eth-txns/logger"
	"github.com/mrdhat/eth-txns/node"
	"github.com/mrdhat/eth-txns/store"
)

func main() {
	addressesToSubscribe := os.Args[1:]
	if len(addressesToSubscribe) == 0 {
		logger.Log("No addresses to subscribe")
		logger.Log("Usage: go run main.go <address1> <address2> <address3> ...")
		os.Exit(1)
	}

	configData := config.NewConfig()

	rpcClient := jsonrpc.NewJSONRPCClient(configData.NodeRPCUrl, &http.Client{})

	blockStore := store.NewBlockStore(store.StoreTypeMemory)
	addressSubscriptionStore := store.NewAddressSubscriptionStore(store.StoreTypeMemory)
	transactionStore := store.NewTransactionStore(store.StoreTypeMemory, addressSubscriptionStore)

	for _, address := range addressesToSubscribe {
		addressSubscriptionStore.Save(entity.AddressSubscription{Address: address, IsActive: true})
	}

	listener := node.NewListener(configData.BlockDuration, rpcClient, blockStore, transactionStore)

	logger.Log("Starting listener")
	stop := make(chan bool)

	// TODO: stop on signal
	go func() {
		time.Sleep(50 * time.Second)
		logger.Log("Stopping listener")
		transactions := transactionStore.GetAll()
		logger.Log("Transactions: ", transactions)
		stop <- true
	}()

	err := listener.Start(stop)
	if err != nil {
		logger.Log("Listener stopped: ", err)
	}
}
