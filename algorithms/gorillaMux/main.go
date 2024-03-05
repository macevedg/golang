package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", usersHandler).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}
}

func usersHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "Mauri" {
		fmt.Println("hit")
	}

	data := []User{
		User{
			Name: "Mauricio",
			Age:  43,
		},
		User{
			Name: "Juli",
			Age:  5,
		},
	}
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(&data)

	if err != nil {

	}
}

type User struct {
	Name string
	Age  int64
}
