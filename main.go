package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/mrdhat/eth-txns/api"
	"github.com/mrdhat/eth-txns/config"
	"github.com/mrdhat/eth-txns/jsonrpc"
	"github.com/mrdhat/eth-txns/logger"
	"github.com/mrdhat/eth-txns/node"
	"github.com/mrdhat/eth-txns/store"
)

func main() {
	configData := config.NewConfig()

	rpcClient := jsonrpc.NewJSONRPCClient(configData.NodeRPCUrl, &http.Client{})

	blockStore := store.NewBlockStore(store.StoreTypeMemory)
	addressSubscriptionStore := store.NewAddressSubscriptionStore(store.StoreTypeMemory)
	transactionStore := store.NewTransactionStore(store.StoreTypeMemory, addressSubscriptionStore)
	parser := api.NewParser(blockStore, addressSubscriptionStore, transactionStore)

	commander := api.NewCommander(parser)

	listener := node.NewListener(configData.BlockDuration, rpcClient, blockStore, transactionStore)

	stop := make(chan bool, 1)
	// stop on process kill
	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		logger.Log(fmt.Sprintf("caught sig: %+v", sig))
		fmt.Println("Wait for 5 second to finish processing")
		time.Sleep(5 * time.Second)
		fmt.Println("Stopping listener")
		stop <- true
		transactions := transactionStore.GetAll()
		logger.Log("Transactions: ", transactions)
		os.Exit(0)
	}()

	// create a wait group & start listener & commander in separate go routines
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		logger.Log("Starting listener")
		err := listener.Start(stop)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		defer wg.Done()
		logger.Log("Starting commander")
		err := commander.Start(stop)
		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}
