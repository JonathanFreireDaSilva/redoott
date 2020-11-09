package db

import (
	"context"
	"time"

	"github.com/JonathanFreireDaSilva/redoott/models"
)

/*InsertRelation graba relacion en la base de datos*/
func InsertRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("redoott")
	collRelations := db.Collection("relations")

	_, err := collRelations.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil

}
