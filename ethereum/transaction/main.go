
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	ctx := context.Background()

	rpcURL := os.Getenv("RPC_URL")
	privateKeyHex := strings.TrimPrefix(os.Getenv("PRIVATE_KEY"), "0x")
	recipient := os.Getenv("TO_ADDRESS")
	amountETH := os.Getenv("AMOUNT_ETH")

	if rpcURL == "" || privateKeyHex == "" {
		log.Fatal("RPC_URL and PRIVATE_KEY must be set in .env")
	}

	if !common.IsHexAddress(recipient) {
		log.Fatal("Invalid TO_ADDRESS in .env")
	}

	if amountETH == "" {
		log.Fatal("AMOUNT_ETH must be set in .env")
	}

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatal("RPC connection failed: ", err)
	}
	defer client.Close()

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal("Failed to get chain ID: ", err)
	}

	if chainID.Cmp(big.NewInt(11155111)) != 0 {
		log.Fatal("Wrong network: expected Sepolia")
	}

	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatal("Failed to get block number: ", err)
	}

	block, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatal("Failed to get latest block: ", err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("Invalid private key: ", err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	to := common.HexToAddress(recipient)

	balance, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		log.Fatal("Balance query failed: ", err)
	}

	fmt.Println("Connected to Ethereum Sepolia")
	fmt.Println("Chain ID:", chainID)
	fmt.Printf("Block Number: %d\n", blockNumber)
	fmt.Printf("Date: %s\n", time.Unix(
		int64(block.Time()), 0,
	).UTC().Format("02 Jan 2006 15:04:05 UTC"))
	fmt.Println("Wallet:", address.Hex())
	fmt.Println("Balance:", weiToETH(balance), "ETH")

	amount, err := parseETH(amountETH)
	if err != nil {
		log.Fatal("Invalid AMOUNT_ETH: ", err)
	}
	if amount.Sign() <= 0 {
		log.Fatal("AMOUNT_ETH must be greater than zero")
	}

	nonce, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		log.Fatal("Failed to get nonce: ", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal("Failed to get gas price: ", err)
	}

	gasLimit, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From:  address,
		To:    &to,
		Value: amount,
	})
	if err != nil {
		log.Fatal("Gas estimation failed: ", err)
	}

	gasCost := new(big.Int).Mul(
		new(big.Int).SetUint64(gasLimit),
		gasPrice,
	)
	required := new(big.Int).Add(
		new(big.Int).Set(amount),
		gasCost,
	)

	if balance.Cmp(required) < 0 {
		log.Fatal("Insufficient Sepolia ETH for amount and gas")
	}

	fmt.Println("Recipient:", to.Hex())
	fmt.Println("Amount:", weiToETH(amount), "ETH")
	fmt.Println("Estimated gas limit:", gasLimit)
	fmt.Println("Estimated gas price:", gasPrice.String(), "wei")

	// Build and sign the transaction.
	tx := types.NewTransaction(
		nonce,
		to,
		amount,
		gasLimit,
		gasPrice,
		nil,
	)

	signedTx, err := types.SignTx(
		tx,
		types.LatestSignerForChainID(chainID),
		privateKey,
	)
	if err != nil {
		log.Fatal("Transaction signing failed: ", err)
	}

	fmt.Println("Transaction hash:", signedTx.Hash().Hex())

	// Broadcast the transaction.
	if err := client.SendTransaction(ctx, signedTx); err != nil {
		log.Fatal("Transaction submission failed: ", err)
	}

	fmt.Println("Transaction submitted. Waiting for confirmation...")

	receipt, err := bind.WaitMined(ctx, client, signedTx)
	if err != nil {
		log.Fatal("Waiting for confirmation failed: ", err)
	}

	fmt.Println("Mined in block:", receipt.BlockNumber)
	fmt.Println("Gas used:", receipt.GasUsed)

	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Fatal("Transaction was mined but failed")
	}

	fmt.Println("Transaction successful!")
	fmt.Printf(
		"View transaction: https://sepolia.etherscan.io/tx/%s\n",
		signedTx.Hash().Hex(),
	)
}

func weiToETH(wei *big.Int) string {
	eth := new(big.Float).SetInt(wei)
	eth.Quo(eth, big.NewFloat(1e18))
	return eth.Text('f', 8)
}

// Convert an ETH decimal string to wei without floating-point rounding.
func parseETH(value string) (*big.Int, error) {
	value = strings.TrimSpace(value)

	if value == "" || strings.HasPrefix(value, "-") ||
		strings.HasPrefix(value, "+") {
		return nil, fmt.Errorf("invalid ETH amount")
	}

	parts := strings.Split(value, ".")
	if len(parts) > 2 {
		return nil, fmt.Errorf("invalid ETH amount")
	}

	whole := parts[0]
	fraction := ""
	if len(parts) == 2 {
		fraction = parts[1]
	}

	if whole == "" {
		whole = "0"
	}
	if len(fraction) > 18 {
		return nil, fmt.Errorf("maximum 18 decimal places allowed")
	}

	for _, part := range []string{whole, fraction} {
		for _, ch := range part {
			if ch < '0' || ch > '9' {
				return nil, fmt.Errorf("invalid ETH amount")
			}
		}
	}

	fraction += strings.Repeat("0", 18-len(fraction))

	wholeWei, ok := new(big.Int).SetString(whole, 10)
	if !ok {
		return nil, fmt.Errorf("invalid ETH amount")
	}

	wholeWei.Mul(
		wholeWei,
		new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil),
	)

	fractionWei := new(big.Int)
	if fraction != "" {
		fractionWei.SetString(fraction, 10)
	}

	return wholeWei.Add(wholeWei, fractionWei), nil
}

// https://sepolia.etherscan.io/tx/0xf9b5e63b92683bd295225e1b02b0fbc6a3f6ded6edd28c912d256149bcced9ff
