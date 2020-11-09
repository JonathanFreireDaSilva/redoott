package routers

import (
	"net/http"

	"github.com/JonathanFreireDaSilva/redoott/db"
	"github.com/JonathanFreireDaSilva/redoott/models"
)

/*DeleteRelation realiza el borrado ded la relacione ntre usuarios*/
func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := db.DeleteRelation(t)
	if err != nil {
		http.Error(w, "Hubo un error al borrar insertar la relacion"+err.Error(), http.StatusBadRequest)
		return

	}

	if status == false {
		http.Error(w, "No se pudo borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
