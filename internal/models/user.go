package models

type User struct {
	Name     string `bson:"name"`
	Password string `bson:"password"`
}
