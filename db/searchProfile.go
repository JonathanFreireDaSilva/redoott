package db

import (
	"context"
	"fmt"
	"time"

	"github.com/JonathanFreireDaSilva/redoott/models"
	"go.mongo.org/mongo-driver/bson"
	"go.mongo.org/mongo-driver/bson/primitive"
)

/**SeachProfile busca un perfil en la base de daos*/
func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	db := MongoCN.Database("redoot")
	collUsers := db.Colection("users")

	var profile models.User

	objID, _ := primitive.ObjectIDFromHEx(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := collUsers.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado" + err.Error())
		return profile, err
	}

	return profile, nil
}
