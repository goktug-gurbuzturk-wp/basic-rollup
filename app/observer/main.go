package main

import (
	"context"
	"log"
	"math/big"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	rollup "app/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	layer1NodeUrl := os.Getenv("LAYER1_NODE_URL")
	client, err := ethclient.Dial(layer1NodeUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contractAddressInHex := os.Getenv("ROLLUP_DATA_LAYER_CONTRACT_ADDRESS")
	contractAddress := common.HexToAddress(contractAddressInHex)
	rollupDataLayer, err := rollup.NewRollupDataLayer(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a RollupDataLayer contract: %v", err)
	}

	opts := &bind.CallOpts{} // Read-only call options
	genesisBlock, err := rollupDataLayer.GenesisBlock(opts)
	if err != nil {
		log.Fatalf("Failed to fetch genesis block from contract: %v", err)
	}

	// Write the genesis block to a JSON file
	err = os.WriteFile("genesis.json", []byte(genesisBlock), 0644)
	if err != nil {
		log.Fatalf("Failed to write genesis block to file: %v", err)
	}

	// Start Anvil with the genesis block
	l2PortNumber := os.Getenv("LAYER2_NODE_PORT")
	cmd := exec.Command("anvil", "--init", "genesis.json", "--port", l2PortNumber)
	err = cmd.Start()
	if err != nil {
		log.Fatalf("Failed to start Anvil: %v", err)
	}

	time.Sleep(5 * time.Second)

	// Create a new Ethereum client connected to the initialized Anvil instance
	layer2NodeUrl := os.Getenv("LAYER2_NODE_URL")
	execClient, err := ethclient.Dial(layer2NodeUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the initialized Anvil instance: %v", err)
	}

	go executeTransactions(rollupDataLayer, execClient)

	// Handle shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan // Wait for termination signal

	// Cleanly shut down Anvil
	err = cmd.Process.Signal(syscall.SIGTERM)
	if err != nil {
		log.Printf("Failed to send SIGTERM to Anvil process: %v", err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Printf("Anvil stopped with error: %v", err)
	}
}

func executeTransactions(rollupDataLayer *rollup.RollupDataLayer, execClient *ethclient.Client) {
	opts := &bind.CallOpts{}
	// Loop and execute transactions
	for i := big.NewInt(0); ; i.Add(i, big.NewInt(1)) {
		txData, err := rollupDataLayer.Transactions(opts, i)
		if err != nil {
			log.Printf("Failed to retrieve transaction at index %d: %v", i, err)
			break
		}

		tx := new(types.Transaction)
		err = rlp.DecodeBytes(txData, &tx)
		if err != nil {
			log.Printf("Failed to decode transaction: %v", err)
			continue
		}

		err = execClient.SendTransaction(context.Background(), tx)
		if err != nil {
			log.Printf("Failed to send transaction: %v", err)
			continue
		}

		log.Printf("Transaction successfully sent: %s", tx.Hash().Hex())
	}
}
