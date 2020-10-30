package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN  es el objeto de conexion a la DB*/
var MongoCN = ConnectDB() //devuelve la coneccion
// variable local    //me permite modificar la url de la base de datos
var clientOptions = options.Client().ApplyURI("mongodb+srv://redoott:redoottsecret@redoott.fb9md.mongodb.net/<dbname>?retryWrites=true&w=majority")

/*ConnectDB es la funcion que me permite conectar*/
func ConnectDB() *mongo.Client {
	//toma la coneccion de clientOption
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("conexion exitosa con la DB")
	return client
}

/*CheckConnection es el ping ala DB*/
func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil) //mi objetod e conecxion es mongoCN
	if err != nil {
		return false
	}
	return true
}
