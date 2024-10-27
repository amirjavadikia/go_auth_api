package routes

import (
	"github.com/amirjavadi/go_auth_api/pkg/controllers"
	"github.com/gorilla/mux"
)

var AuthRouters = func(router *mux.Router) {
	router.HandleFunc("/register/", controllers.Register).Methods("POST")
	router.HandleFunc("/login/", controllers.Login).Methods("POST")
	router.HandleFunc("/logout/", controllers.Logout).Methods("POST")
}
