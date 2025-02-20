package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"crypto/ecdsa"
	"log"


)


func main(){

	privateKey,err :=  crypto.GenerateKey()   // this line generates the privateKey

	if (err != nil){
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Println(" PrivateKey in Hex:  ",hexutil.Encode(privateKeyBytes[2:]))  // we are converting the privateKey in bytes to hexadecimal  and we  chop of the "0x"

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)  // sanity check

	if(!ok){
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	fmt.Println(" Public Key Bytes:  ",hexutil.Encode(publicKeyBytes)[4:])

	address :=  crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	fmt.Println(" Address:  ",address)

	hash := sha3.NewLegacyKeccak256()

	hash.Write(publicKeyBytes[1:])

	fmt.Println(" Hash:  ",hexutil.Encode(hash.Sum(nil)[12:]))

}