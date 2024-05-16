package main

import (
	"log"
	"math/big"
	"os"
	"path/filepath"

	rollup "app/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

	sequencerPrivateKey := os.Getenv("SEQUENCER_PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(sequencerPrivateKey)
	if err != nil {
		log.Fatalf("Could not parse private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	contractAddress := os.Getenv("ROLLUP_DATA_LAYER_CONTRACT_ADDRESS")
	address := common.HexToAddress(contractAddress)
	rollupDataLayer, err := rollup.NewRollupDataLayer(address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a RollupDataLayer contract: %v", err)
	}

	// Reading multiple RLP-encoded transactions from a directory
	files, err := filepath.Glob("../transaction-builder/*.rlp")
	if err != nil {
		log.Fatalf("Failed to list transaction files: %v", err)
	}

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Printf("Failed to read transaction data from %s: %v", file, err)
			continue
		}

		log.Printf("Read transaction data from %s: %v", file, data)

		// Submit the transaction to the smart contract
		tx, err := rollupDataLayer.AddRollupTransaction(auth, data)
		if err != nil {
			log.Printf("Failed to send transaction from %s: %v", file, err)
			continue
		}

		log.Printf("Transaction submitted from %s: %s", file, tx.Hash().Hex())
	}
}
