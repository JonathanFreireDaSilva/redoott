package models

import "time"

/*SaveState formato que tendra nestro estado*/
type SaveState struct {
	UserID  string    `bson:"userid" json:"userId,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
