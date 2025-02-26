package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)



func main() {
	
	re:= regexp.MustCompile("^0x[0-9a-fA-F]{40}$")  // we are trying to check if the ethereum addresses are valid

	fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid true
	fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid false


	client, err := ethclient.Dial("https://mainnet.infura.io/v3/f2e799e62655472cbaf43f0cea4d765e") 

	if err != nil{
		log.Fatal(err)
	}

	 // 0x Protocol Token (ZRX) smart contract address

	 address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")  // we are converting this because this is case sensitive

	 bytecode, err  := client.CodeAt(context.Background(),address,nil)


	 isContract :=  len(bytecode) > 0

	 fmt.Println("isContract: ", isContract)  // true


	     // a random user account address


		 _address := common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")  // we are converting this because this is case sensitive

		 _bytecode, err  := client.CodeAt(context.Background(),_address,nil)
	
	
		 _isContract :=  len(_bytecode) > 0
	
		 fmt.Println("isContract: ", _isContract)  // false
}