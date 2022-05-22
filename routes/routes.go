package routes

import (
	"github.com/gorilla/mux"
	"github.com/n0nnahs/polished-api/controllers"
)

func PolishRoute(router *mux.Router){
    router.HandleFunc("/user", controllers.CreatePolish()).Methods("POST")
}