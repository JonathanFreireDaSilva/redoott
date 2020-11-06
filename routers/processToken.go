package routers

import (
	"errors"
	"strings"

	"github.com/JonathanFreireDaSilva/redoott/db"
	"github.com/JonathanFreireDaSilva/redoott/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email valor de email usadi e todos los endponts*/
var Email string

/*IDUser es el id devuelto del modelo, que se usara en todos los endpoints*/
var IDUser string

/*ProcessToken proce token para extraer sus valores*/
func ProcessToken(myToken string) (*models.Claim, bool, string, error) {
	myKey := []byte("myKeySuperSecretisimaaaaaaa")
	claims := &models.Claim{}

	splitToken := strings.Split(myToken, "Bearer")
	if len(myToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	myToken = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(myToken, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, find, _ := db.CheckUserExist(claims.Email)
		if find == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, find, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
