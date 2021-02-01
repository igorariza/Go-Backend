package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/igorariza/Go-Backend/middlew"
	"github.com/igorariza/Go-Backend/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores comment generic
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routes.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routes.Login)).Methods("POST")
	// router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routes.VerPerfil))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
