package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JonathanFreireDaSilva/redoott/db"
	"github.com/JonathanFreireDaSilva/redoott/models"
)

/*SaveState permite grabar el twwt en la base de datos*/
func SaveState(w http.ResponseWriter, r *http.Request) {
	var message models.State
	err := json.NewDecoder(r.Body).Decode(&message)

	registry := models.SaveState{
		UserID:  IDUser,
		Mensaje: message.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertState(registry)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, reintentar"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insetar el estado", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
