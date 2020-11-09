package routers

import (
	"net/http"

	"github.com/JonathanFreireDaSilva/redoott/db"
	"github.com/JonathanFreireDaSilva/redoott/models"
)

/*UploadRelation realiza registro de relacion entre users */
func UploadRelation(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID
	status, err := db.InsertRelation(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se logro insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
