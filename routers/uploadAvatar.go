package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/JonathanFreireDaSilva/redoott/db"
	"github.com/JonathanFreireDaSilva/redoott/models"
)

/*UploadAvatar subir avatar al servidro*/
func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	var exten = strings.Split(handler.Filename, ".")[1]
	var archive string = "uploads/avatars/" + IDUser + "." + exten

	f, err := os.OpenFile(archive, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir imagen"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen!"+err.Error(), http.StatusBadRequest)
		return
	}
	var user models.User
	var status bool

	user.Avatar = IDUser + "." + exten
	status, err = db.ModifyRegister(user, IDUser)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatar en la BD!"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
