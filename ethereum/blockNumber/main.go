package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

func main() {
	rpcURL := "https://rpc.mevblocker.io"

	client, err := ethclient.Dial(rpcURL)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	ctx := context.Background()

	blockNumber, err := client.BlockNumber(ctx)

	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	timestamp := block.Time()

	fmt.Printf("Block: %d | Date: %s\n",
		blockNumber,
		time.Unix(int64(timestamp), 0).UTC().Format("02 Jan 2006 15:04:05 UTC"))
}
