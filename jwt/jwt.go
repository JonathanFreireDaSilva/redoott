package jwt

import (
	"time"

	"github.com/JonathanFreireDaSilva/redoott/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GenerateJWT genero encriptado de datos con JWT*/
func GenerateJWT(t models.User) (string, error) {

	myKey := []byte("myKeySuperSecretisimaaaaaaa")

	payload := jwt.MapClaims{
		"email":           t.Email,
		"nombre":          t.Nombre,
		"apellidos":       t.Apellidos,
		"fechaNacimiento": t.FechaNacimiento,
		"biografia":       t.Biografia,
		"ubicacion":       t.Ubicacion,
		"sitioWeb":        t.SitioWeb,
		"_id":             t.ID.Hex(),
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
