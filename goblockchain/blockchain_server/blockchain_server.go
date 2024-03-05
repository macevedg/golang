package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/macevedg/goblockchain/blockchain"
	"github.com/macevedg/goblockchain/wallet"
)

var cache map[string]*blockchain.Blockchain = make(map[string]*blockchain.Blockchain)

type BlockchainServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{
		port: port,
	}
}

func (bcs *BlockchainServer) Port() uint16 { return bcs.port }

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!")
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", bcs.GetChain)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}

//api

func (bcs *BlockchainServer) GetBlockchain() *blockchain.Blockchain {

	bc, ok := cache["blockchain"]

	if !ok {
		minersWallet := wallet.NewWallet()
		bc = blockchain.NewBlockchain(minersWallet.BlockchainAddress(), bcs.Port())
		cache["blockchain"] = bc
		log.Printf("private key: %v\n", minersWallet.PrivateKeyStr())
		log.Printf("public key: %v\n", minersWallet.PublicKeyStr())
		log.Printf("blockchain address: %v\n", minersWallet.BlockchainAddress())
	}

	return bc
}

func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		{
			w.Header().Add("Content-Type", "application/json")
			bc := bcs.GetBlockchain()
			m, _ := bc.MarshalJSON()
			io.WriteString(w, string(m[:]))
		}
	default:
		log.Printf("Error: Invalid HTTP Method: %v\n", req.Method)
	}

}
