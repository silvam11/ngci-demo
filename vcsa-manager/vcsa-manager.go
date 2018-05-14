package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/data", getData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Enpoint Test called.")

	repoManager := os.Getenv("REPO_MANAGER")
	fmt.Println("REPO_MANAGER ", repoManager)

	resp, err := http.Get(repoManager + "colors")
	if err != nil {
		fmt.Println("HTTP ERROR: %s ", err)
		panic(err)
        }
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("BODY: ", string(body))
	json.NewEncoder(w).Encode(string(body))
}

