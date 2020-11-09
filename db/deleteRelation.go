package db

import (
	"context"
	"time"

	"github.com/JonathanFreireDaSilva/redoott/models"
)

/*DeleteRelation borra una relacion de la base de deadatos*/
func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redoott")
	collRelation := db.Collection("relations")

	_, err := collRelation.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
