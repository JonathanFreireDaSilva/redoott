package routers

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/JonathanFreireDaSilva/redoott/db"
)

/*GetAllRelations leo la lsita de todos los usuauios con relacion*/
func GetAllRelations(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("tipe")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina como entero maypr a 0"+err.Error(), http.StatusBadRequest)
		return

	}

	pag := int64(pagTemp)

	result, status := db.ReadAllUsers(IDUser, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al leer los users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}
