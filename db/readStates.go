package db

import (
	"context"
	"log"
	"time"

	"github.com/JonathanFreireDaSilva/redoott/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ReadStates leo estados de un perfil */
func ReadStates(ID string, page int64) ([]*models.ReturnStates, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redoott")
	collStates := db.Collection("states")

	var statesSlice []*models.ReturnStates

	condition := bson.M{
		"userid": ID,
	}

	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "fecha", Value: -1}})
	options.SetSkip((page - 1) * 20)

	data, err := collStates.Find(ctx, condition, options)

	if err != nil {
		log.Fatal(err.Error())
		return statesSlice, false
	}

	for data.Next(context.TODO()) {

		var registry models.ReturnStates

		err := data.Decode(&registry)

		if err != nil {
			return statesSlice, false
		}

		statesSlice = append(statesSlice, &registry)

	}
	return statesSlice, true

}
