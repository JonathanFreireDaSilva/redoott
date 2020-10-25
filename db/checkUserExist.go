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
	col := db.Collection("users")

	condicion := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condicion).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
