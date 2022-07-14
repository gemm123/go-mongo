package models

type User struct {
	Name string `json:"name" bson:"name"`
	Age  int    `json:"age" bson:"age"`
	City string `json:"city" bson:"city"`
}
