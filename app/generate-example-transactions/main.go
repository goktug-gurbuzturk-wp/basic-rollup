package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	const receiver = "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC" // wallet address of the third account in anvil

	client, err := ethclient.Dial("http://localhost:8546")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d") // private key for the second account in anvil
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000) // 1 ETH
	gasLimit := uint64(21000)
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
	fmt.Printf("RLP-encoded transaction: 0x%x\n", rlpEncodedTx)

	// Storing the RLP-encoded transaction in a file
	err = ioutil.WriteFile("transaction.rlp", rlpEncodedTx, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction stored in 'transaction.rlp'")
}
