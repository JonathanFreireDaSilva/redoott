package db

import (
	"context"
	"time"

	"github.com/JonathanFreireDaSilva/redoott/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ReadStatesFollows leo tweets de mis seguidores*/
func ReadStatesFollows(ID string, page int) ([]models.ReturnStatesFollows, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redoott")

	collRelation := db.Collection("relations")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "states",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "states",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$states"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"states.fecha": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	data, err := collRelation.Aggregate(ctx, conditions)
	var result []models.ReturnStatesFollows

	err = data.All(ctx, &result)

	if err != nil {
		return result, false
	}

	return result, true

}
