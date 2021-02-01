package main

import (
	"log"

	"github.com/igorariza/Go-Backend/bd"
	"github.com/igorariza/Go-Backend/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la bd")
		return
	}
	handlers.Manejadores()
}
