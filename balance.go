package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"math"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, err := ethclient.Dial("https://cloudflare-eth.com")  // add your own rpc 
	if err != nil {
		log.Fatal(err)
	}

	account  := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance , err :=  client.BalanceAt(context.Background(), account, nil)

	if(err != nil){
		log.Fatal(err)
	}
	fmt.Println(balance)

	blocknumber:= big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blocknumber)
	
	if(err != nil){
		log.Fatal(err)
	}


	fmt.Println(balanceAt)


	 floatBalance := new(big.Float)

	 floatBalance.SetString(balanceAt.String())

	 ethValue := new(big.Float).Quo(floatBalance, big.NewFloat(math.Pow10(18)))
	 
	 fmt.Println(ethValue)

}