package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JonathanFreireDaSilva/redoott/db"
)

/*ReadStates leo estados */
func ReadStates(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro page", http.StatusBadRequest)
		return

	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "Debe enviar el parametri pagina con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	pageForUse := int64(page)

	res, correct := db.ReadStates(ID, pageForUse)
	if correct == false {
		http.Error(w, "Error al leer los estados", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
