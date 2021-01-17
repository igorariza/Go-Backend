package routes

import (
	"encoding/json"
	"net/http"

	"github.com/igorariza/Go-Backend/bd"
	"github.com/igorariza/Go-Backend/models"
)

//Registro es la función para crear el usuario en la BD
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Contraseña muy corta ", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe usuario ", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio error al intentar registrar el usuario "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "no se inserto el registro", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
