package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JonathanFreireDaSilva/redoott/db"
	"github.com/JonathanFreireDaSilva/redoott/models"
)

/*Register es la funcuin para crear en la DB el regitrso de usuario*/
func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "ERROR EN LOS DATOS RECIBIDOS"+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de almenos 6 caracteres", 400)
		return
	}

	_, finded, _ := db.CheckUserExist(user.Email)
	if finded == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := db.InsertRegistry(user)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
