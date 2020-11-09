package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JonathanFreireDaSilva/redoott/db"
	"github.com/JonathanFreireDaSilva/redoott/models"
)

/*ModifyProfile modifica el perfil*/
func ModifyProfile(w http.ResponseWriter, r *http.Request) {

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
	}

	var status bool
	status, err = db.ModifyRegister(t, IDUser)

	if err != nil {
		http.Error(w, "Ocurrio un error al modifiar el registro.Reintentar"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modifica el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
