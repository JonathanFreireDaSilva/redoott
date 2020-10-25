package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/JonathanFreireDaSilva/redoott/middleware"
	"github.com/JonathanFreireDaSilva/redoott/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo el puerto, pongo a escuchar al servidor*/
func Manejadores() {

	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT") // me fijo si tengo una variable de entorno con el nombre port definida
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) //damos permiso a cualquiera para usar
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}