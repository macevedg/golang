package main

import (
	//. "github.com/macevedg/goblockchain/blockchain"
	"fmt"

	"github.com/macevedg/goblockchain/blockchain"
	"github.com/macevedg/goblockchain/wallet"
)

func main() {
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()
	walletM := wallet.NewWallet()

	//wallet
	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)

	//blockchain
	blockchain := blockchain.NewBlockchain(walletM.BlockchainAddress(), 8080)

	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0,
		walletA.PublicKey(), t.GenerateSignature())

	fmt.Println("Added?", isAdded)
	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))
}

/*

func main2() {
	myblockchainAddress := "my_block_chain_address"
	blockChain := NewBlockchain(myblockchainAddress)

	fmt.Println(blockChain)
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)
	blockChain.Mining()
	blockChain.Print()

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)
	blockChain.Mining()
	blockChain.Print()

	fmt.Printf("my %.1f\n", blockChain.CalculateTotalAmount("my_block_chain_address"))
	fmt.Printf("C %.1f\n", blockChain.CalculateTotalAmount("C"))
	fmt.Printf("D %.1f\n", blockChain.CalculateTotalAmount("D"))

}

func NewBlockchain(myblockchainAddress string) {
	panic("unimplemented")
}
*/
