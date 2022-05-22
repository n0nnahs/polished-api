package main

import (
	"log"
	"net/http"

	"github.com/n0nnahs/polished-api/configs"

	"github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

	configs.ConnectDB()

	log.Fatal(http.ListenAndServe(":6000", router))
}