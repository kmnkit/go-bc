package main

import (
	"fmt"
	"log"

	"github.com/kmnkit/go-bc/wallet"
)

func init() {
	log.SetPrefix("BlockChain: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	fmt.Println(w.BlockchainAddress())
}
