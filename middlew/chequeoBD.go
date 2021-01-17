package middlew

import (
	"net/http"

	"github.com/igorariza/Go-Backend/bd"
)

//ChequeoBD es el middlew y permite conocer el estado de la BD
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)

	}
}
