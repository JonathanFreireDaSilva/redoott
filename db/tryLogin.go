package db

import (
	"github.com/JonathanFreireDaSilva/redoott/models"
	"golang.org/x/crypto/bcrypt"
)

/*TryLogin realiza el chqueo de login ala DB*/
func TryLogin(email string, password string) (models.User, bool) {

	user, found, _ := CheckUserExist(email)
	if found == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}

	return user, true
}
