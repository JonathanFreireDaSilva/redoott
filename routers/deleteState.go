package routers

import (
	"net/http"

	"github.com/JonathanFreireDaSilva/redoott/db"
)

/*DeleteState borra un estado determinado*/
func DeleteState(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, " Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	err := db.DeleteState(ID, IDUser)
	if err != nil {
		http.Error(w, "Ocurrio un error al intetnar borrar el twteet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
