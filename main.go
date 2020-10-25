package main

import (
	"log"

	"github.com/JonathanFreireDaSilva/redoott/db"
	"github.com/JonathanFreireDaSilva/redoott/handlers"
)

func main() {
	if db.CheckConnection() == false {
		log.Fatal("SIN CONEXION ALA DB ")
		return
	}
	handlers.Manejadores()
}
