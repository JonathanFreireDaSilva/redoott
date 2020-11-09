package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JonathanFreireDaSilva/redoott/db"
	"github.com/JonathanFreireDaSilva/redoott/models"
)

/*GetRelation chequqea si hay relacion entre dos users*/
func GetRelation(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	var resp models.ResponseGetRelation

	status, err := db.GetRelation(t)

	if err != nil || status == false {

		resp.Status = false

	} else {

		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
