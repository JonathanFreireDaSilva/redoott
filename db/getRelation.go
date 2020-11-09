package db

import (
	"context"
	"fmt"
	"time"

	"github.com/JonathanFreireDaSilva/redoott/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*GetRelation consulta relaciones de 2 users*/
func GetRelation(t models.Relation) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) //le decimos al context.TODO que teniamos que espere 15 segundos
	defer cancel()

	db := MongoCN.Database("redoott")
	collRelation := db.Collection("relations")

	condition := bson.M{
		"userid":         t.UserID,
		"userrelationid": t.UserRelationID,
	}

	var result models.Relation

	fmt.Print(result)
	err := collRelation.FindOne(ctx, condition).Decode(&result)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil

}
