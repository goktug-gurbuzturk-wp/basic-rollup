package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	const receiver = "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC" // wallet address of the third account in Anvil

	client, err := ethclient.Dial("http://localhost:8546")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d") // private key for the second account in Anvil
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	value := big.NewInt(1000000000000000000) // 1 ETH
	gasLimit := uint64(21000)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 3; i++ { // Create 3 transactions
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		toAddress := common.HexToAddress(receiver)
		tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			log.Fatal(err)
		}

		var buf bytes.Buffer
		err = rlp.Encode(&buf, signedTx)
		if err != nil {
			log.Fatal(err)
		}

		rlpEncodedTx := buf.Bytes()
		filename := fmt.Sprintf("transaction_%d.rlp", i+1)
		err = os.WriteFile(filename, rlpEncodedTx, 0644)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Transaction %d stored in '%s'\n", i+1, filename)
		nonce++
	}
}
