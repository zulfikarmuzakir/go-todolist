package model

type Todo struct {
	Id          string `json:"id" bson:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
