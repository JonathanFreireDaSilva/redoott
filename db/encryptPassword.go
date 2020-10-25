package db

import "golang.org/x/crypto/bcrypt"

/*EncryptPassword es la rutina que me permite encriptar la password de lapos usuarios*/
func EncryptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err // es la cantidad de pasadas que va a ser para encriptar el password

}
