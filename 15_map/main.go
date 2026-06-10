package main

import "fmt"

func main() {
	// A map stores data as key-value pairs
	balances := map[string]float64{
		"ETH": 3,
		"SOL": 15,
		"BNB": 4,
	}
	wallets := make(map[string]string)

	wallets["Bob"] = "0xD62cFd18bdB62D6f2B21CB73E3D8B13359cEa4eA"
	wallets["Alice"] = "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"

	fmt.Println(balances) // map[BNB:4 ETH:3 SOL:15]
	delete(balances, "BNB")
	fmt.Println(balances)
	fmt.Println(wallets)
}