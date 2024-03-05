package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/macevedg/goblockchain/utils"
	"github.com/macevedg/goblockchain/wallet"
)

const (
	tempDir = "templates"
)

type WalletServer struct {
	port    uint16 //
	gateway string //
}

func NewWalletServer(port uint16, gateway string) *WalletServer {

	return &WalletServer{port: port, gateway: gateway}
}

func (s *WalletServer) Port() uint16 { return s.port }

func (s *WalletServer) Gateway() string { return s.gateway }

func (s *WalletServer) Index(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		{
			t, _ := template.ParseFiles(path.Join(tempDir, "index.html"))
			log.Printf("template object %v\n", path.Join(tempDir, "index.html"))
			t.Execute(w, "")
		}
	default:
		log.Printf("ERROR: Invalid HTTP method: %s", r.Method)
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}

func (s *WalletServer) Wallet(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		{
			w.Header().Add("Content-Type", "application/json")
			mywallet := wallet.NewWallet()
			m, _ := mywallet.MarshalJSON()

			io.WriteString(w, string(m[:]))
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERROR: Invalid HTTP method: %s", r.Method)
	}
}

func (s *WalletServer) CreateTransaction(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		{
			io.WriteString(w, string(utils.JsonStatus("success!")))
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERROR: Invalid HTTP method: %s", r.Method)
	}
}
