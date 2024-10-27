package main

import (
	"github.com/amirjavadi/go_auth_api/pkg/config"
	"github.com/amirjavadi/go_auth_api/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const port = ":9090"

func main() {
	r := mux.NewRouter()
	routes.AuthRouters(r)
	log.Printf("Server is running on port %s", port)

	handler := config.ConfigureCORS(r)

	if err := http.ListenAndServe(port, handler); err != nil {
		log.Printf("Faild to start server : %v", err)
	}
}
