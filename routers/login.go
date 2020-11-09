package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JonathanFreireDaSilva/redoott/db"
	"github.com/JonathanFreireDaSilva/redoott/jwt"
	"github.com/JonathanFreireDaSilva/redoott/models"
)

/*Login mvalidaciones de ruta login*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "aplication/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "usuarios y/o contraseña invalidos"+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	document, exist := db.TryLogin(user.Email, user.Password)
	if exist == false {
		http.Error(w, "Usuario y/o Contraseña invalidos", 400)
		return

	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el token correspondiente"+err.Error(), 400)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
