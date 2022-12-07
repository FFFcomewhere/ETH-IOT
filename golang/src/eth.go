package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://goerli.infura.io/v3/132ff1db81e74017bd120ce624a02d1f")
	if err != nil {
		panic(err)
	}

	account := common.HexToAddress("0x58B17C0F317324cd804Ee9bf910D018fd3A91844")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(balance)
}
