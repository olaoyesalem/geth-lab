package main

import (

	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs(){

	ks:= keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)

	password:= "secret"

	account, err := ks.NewAccount(password)

	if err!= nil{
		log.Fatal(err)
	}

	fmt.Println("just Account", account)
	fmt.Println("just Account with Address", account.Address)
	fmt.Println("just Account with Address and Hex", account.Address.Hex())

	
}

func importKs(){

		file := "./tmp/ this is gotten from the tmp file when we run createKs"
		ks :=keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP);
		jsonBytes, err:=ioutil.ReadFile(file)

		if(err != nil){
			log.Fatal(err)
		}

		password := "secret"

		account,err := ks.Import(jsonBytes, password, password)

		if(err != nil){
			log.Fatal(err)
		}
		fmt.Println(account.Address.Hex())

		if err := os.Remove(file); err != nil{
			log.Fatal(err)
		}

}


func main(){
	createKs()
	//importKs()
}