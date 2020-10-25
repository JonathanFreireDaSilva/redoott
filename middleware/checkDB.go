package middleware

import (
	"net/http"

	"github.com/JonathanFreireDaSilva/redoott/db"
)

/*CheckDB es el middlew que me permite conocer el estado de la DB */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == false {
			http.Error(w, "Conexion perdida con la DB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}
