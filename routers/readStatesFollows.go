package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JonathanFreireDaSilva/redoott/db"
)

func ReadStatesFollows(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro page", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "Debe enviar el aprametro apgina como entero mayor a 0", http.StatusBadRequest)
		return
	}

	response, correct := db.ReadStatesFollows(IDUser, page)
	if correct == false {
		http.Error(w, "Error al leer los estados", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
