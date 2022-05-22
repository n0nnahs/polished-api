package main

import (
	"log"
	"net/http"

	"github.com/n0nnahs/polished-api/configs"
	"github.com/n0nnahs/polished-api/routes"

	"github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

	configs.ConnectDB()

	routes.PolishRoute(router)

	log.Fatal(http.ListenAndServe(":6000", router))
}