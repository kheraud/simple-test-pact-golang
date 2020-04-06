package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MyResponse struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Age         int    `json:"age"`
	Catchphrase string `json:"catchphrase"`
}

// startServer starts a real server
func startServer() {
	myResponse := MyResponse{
		Firstname:   "K",
		Lastname:    "Heraud",
		Age:         35,
		Catchphrase: "Oooh ooh ooh aha ooh ooh aha",
	}

	jsonResponse, _ := json.Marshal(myResponse)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintln(w, string(jsonResponse))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	startServer()
}
