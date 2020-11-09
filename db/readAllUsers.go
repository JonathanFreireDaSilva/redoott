package db

import (
	"context"
	"fmt"
	"time"

	"github.com/JonathanFreireDaSilva/redoott/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ReadAllUsers ______- */
func ReadAllUsers(ID string, page int64, search string, tipe string) ([]*models.User, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redoott")
	collUsers := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	data, err := collUsers.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false

	}

	var found, include bool

	for data.Next(ctx) {
		var s models.User
		err := data.Decode(&s)

		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relation

		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		include = false

		found, err = GetRelation(r)

		if tipe == "new" && found == false {
			include = true
		}

		if tipe == "follow" && found == true {
			include = true
		}

		if r.UserRelationID == ID {
			include = false
		}

		if include == true {
			s.Password = ""
			s.Biografia = ""
			s.Apellidos = ""
			s.SitioWeb = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)

		}

	}

	err = data.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false

	}

	data.Close(ctx)
	return results, true
}
