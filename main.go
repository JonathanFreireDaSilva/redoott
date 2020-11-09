package main

import (
	"log"

	"github.com/JonathanFreireDaSilva/redoott/db"
	"github.com/JonathanFreireDaSilva/redoott/handlers"
)

//fmt.Println(splitToken)
func main() {
	if db.CheckConnection() == false {
		log.Fatal("SIN CONEXION A LA DB ")
		return
	}
	handlers.Manejadores()
}
