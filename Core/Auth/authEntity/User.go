package authEntity

import "go.mongodb.org/mongo-driver/bson/primitive"

//Email is the unique key

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Password    string             `bson:"password"`
	Email       string             `bson:"email"`
	UseProvider bool
	Provider
}
