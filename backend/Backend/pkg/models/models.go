package models

type Task struct {
	Id     int    `json:"Id" bson:"Id"`
	Title  string `json:"Title" bson:"Title" validate:"required"`
	Status string `json:"Status" bson:"Status" `
}
