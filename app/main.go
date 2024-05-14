package main

import (
	"log"
	"math/big"

	rollup "app/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80") // sequencer private key - first account in anvil
	if err != nil {
		log.Fatalf("Could not parse private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	address := common.HexToAddress("0x5fbdb2315678afecb367f032d93f642f64180aa3")
	rollupDataLayer, err := rollup.NewRollupDataLayer(address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a RollupDataLayer contract: %v", err)
	}

	tx, err := rollupDataLayer.AddRollupTransaction(auth, []byte("example transaction data2"))
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	log.Printf("Transaction submitted: %s", tx.Hash().Hex())
}
