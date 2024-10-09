package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
)

func processLogIndexer(client *ethclient.Client, processor adapters.TransferLogsIndexer, blockNumber int64) error {
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch the chain ID: %v", err)
		return err
	}
	query := ethereum.FilterQuery{
		Addresses: processor.Addresses(),
		FromBlock: big.NewInt(blockNumber),
		ToBlock:   big.NewInt(blockNumber),
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to fetch the logs: %v", err)
		return err
	}
	senders, err := processor.IndexLogs(context.Background(), chainID, client, logs)
	if err != nil {
		log.Fatalf("Failed to process the logs: %v", err)
		return err
	}

	fmt.Printf("Senders: %v\n", senders)
	return nil
}

func processLPLogIndexer(client *ethclient.Client, processor adapters.LPLogsIndexer, blockNumber int64) error {
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch the chain ID: %v", err)
		return err
	}
	query := ethereum.FilterQuery{
		Addresses: processor.Address(),
		FromBlock: big.NewInt(blockNumber),
		ToBlock:   big.NewInt(blockNumber),
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to fetch the logs: %v", err)
		return err
	}
	senders, err := processor.IndexLogs(context.Background(), chainID, client, logs)
	if err != nil {
		log.Fatalf("Failed to process the logs: %v", err)
		return err
	}

	fmt.Printf("Senders: %v\n", senders)
	return nil
}

func processLockIndexer(client *ethclient.Client, processor adapters.LockLogsIndexer, blockNumber int64) error {
	ctx := context.Background()
	query := ethereum.FilterQuery{
		Addresses: processor.Address(),
		FromBlock: big.NewInt(blockNumber),
		ToBlock:   big.NewInt(blockNumber),
	}
	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		log.Fatalf("Failed to fetch the logs: %v", err)
		return err
	}
	senders, err := processor.IndexLogs(ctx, logs)
	if err != nil {
		log.Fatalf("Failed to process the logs: %v", err)
		return err
	}

	fmt.Printf("Senders: %v\n", senders)
	return nil
}
