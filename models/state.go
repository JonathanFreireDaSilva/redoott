package models

/*State captura del body, el mensaje que nos llega*/
type State struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
