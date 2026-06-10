package main

import "fmt"

func main() {
	wallets := []string{
		"0xE556bF2A4452Dba4B02598c1810900e0DD3C1F58",
		"0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97",
		"0xD62cFd18bdB62D6f2B21CB73E3D8B13359cEa4eA",
	}
	tokenPrices := []float64{0.12, 0.45, 1.25}

	fmt.Println(wallets)
	fmt.Println(tokenPrices[1])
}