package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ReturnStatesFollows estructura con la que devolvemos de los estados de seguidores*/
type ReturnStatesFollows struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userid" json:"userId,omitempty"`
	UserRelationID string             `bson:"userrelationid" json:"userRelationId,omitempty"`
	States         struct {
		ID      string    `bson:"_id" json:"_id,omitempty"`
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
	}
}
