package main

import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.GenerateKey()

	if err != nil {
		log.Fatal("Failed to generate private key: ", err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	privateKeyString := fmt.Sprintf("%x", crypto.FromECDSA(privateKey))

	fmt.Println("==================================== Ethereum Wallet ==================================")
	fmt.Println("Address: ", address)
	fmt.Println("Private Key: 0x" + privateKeyString)
	fmt.Println("=======================================================================================")
}

// Generate ethereum wallet and private key
