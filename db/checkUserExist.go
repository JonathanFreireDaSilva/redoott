package db

import (
	"context"
	"time"

	"github.com/JonathanFreireDaSilva/redoott/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CheckUserExist recibe un emal de parametro y chequea si esta en la abse de datos*/
func CheckUserExist(email string) (models.User, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redoott")
	userCollection := db.Collection("users")

	condicion := bson.M{"email": email}

	var user models.User

	err := userCollection.FindOne(ctx, condicion).Decode(&user) //si no encuentra nada guarda en error y si encuentra algoo lo guarda en user
	ID := user.ID.Hex()

	if err != nil {
		return user, false, ID
	}

	return user, true, ID
}
