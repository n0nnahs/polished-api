package routes

import (
	"github.com/gorilla/mux"
	"github.com/n0nnahs/polished-api/controllers"
)

func PolishRoute(router *mux.Router){
	router.HandleFunc("/polish", controllers.CreatePolish()).Methods("POST")
	router.HandleFunc("/polishes", controllers.GetAllPolish()).Methods("GET") //add this
	router.HandleFunc("/polish/id/{polishId}", controllers.GetPolishId()).Methods("GET")
	router.HandleFunc("/polish/{name}", controllers.GetPolishName()).Methods("GET")
	router.HandleFunc("/polish/{polishId}", controllers.EditPolish()).Methods("PUT") 
	router.HandleFunc("/user/{userId}", controllers.DeletePolish()).Methods("DELETE")
}