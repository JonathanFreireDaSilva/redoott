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

	router := mux.NewRouter() //captura el http y le da manejo al response y al request que viene en el llamado de la api

	router.HandleFunc("/register", middleware.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/viewProfile", middleware.CheckDB(middleware.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modifyProfile", middleware.CheckDB(middleware.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/state", middleware.CheckDB(middleware.ValidateJWT(routers.SaveState))).Methods("POST")
	router.HandleFunc("/readStates", middleware.CheckDB(middleware.ValidateJWT(routers.ReadStates))).Methods("GET")
	router.HandleFunc("/deleteState", middleware.CheckDB(middleware.ValidateJWT(routers.DeleteState))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middleware.CheckDB(middleware.ValidateJWT(routers.UploadAvatar))).Methods("GET")
	router.HandleFunc("/uploadBanner", middleware.CheckDB(middleware.ValidateJWT(routers.UploadBanner))).Methods("GET")
	router.HandleFunc("/getAvatar", middleware.CheckDB(middleware.ValidateJWT(routers.GetAvatar))).Methods("GET")
	router.HandleFunc("/getAvatar", middleware.CheckDB(middleware.ValidateJWT(routers.GetBanner))).Methods("GET")

	router.HandleFunc("/uploadRelation", middleware.CheckDB(middleware.ValidateJWT(routers.UploadRelation))).Methods("POST")
	router.HandleFunc("/deleteRelation", middleware.CheckDB(middleware.ValidateJWT(routers.DeleteRelation))).Methods("DELETE")

	router.HandleFunc("/getRelation", middleware.CheckDB(middleware.ValidateJWT(routers.GetRelation))).Methods("GET")
	router.HandleFunc("/getAllRelations", middleware.CheckDB(middleware.ValidateJWT(routers.GetAllRelations))).Methods("GET")

	router.HandleFunc("/readStatesFollows", middleware.CheckDB(middleware.ValidateJWT(routers.ReadStatesFollows))).Methods("GET")

	PORT := os.Getenv("PORT") // me fijo si tengo una variable de entorno con el nombre port definida
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) //damos permiso a cualquiera para usar
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
