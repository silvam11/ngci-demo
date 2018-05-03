package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Test)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Test(w http.ResponseWriter, r *http.Request) {
	response := HipChatResponse{Color: "yellow", Message: "This is a Test 22", Notify: "false", MessageFormat: "text"}
	json.NewEncoder(w).Encode(response)
}

type HipChatResponse struct {
	Color         string `json:"color"`
	Message       string `json:"message"`
	Notify        string `json:"notify"`
	MessageFormat string `json:"message_format"`
}
