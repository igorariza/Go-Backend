package routes

import (
	json "encoding/json"
	"net/http"

	"github.com/igorariza/Go-Backend/bd"
)

// VerPerfil commet generic
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro"+err.Error(), 400)
		return
	}

	w.Header().Set("contect-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)

}
