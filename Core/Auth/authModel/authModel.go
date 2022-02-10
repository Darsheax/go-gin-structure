package authModel

import (
	"context"
	"errors"
	"fmt"
	"root/core/auth/authEntity"
	"root/core/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthEntity model.Entity

func (e *AuthEntity) AuthLogin(email string) authEntity.User {

	var user authEntity.User

	collectionUser := e.DataBase.Collection("user")
	err := collectionUser.FindOne(e.AppContext, bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user
		}
		panic(err)
	}

	return user

}

func (e *AuthEntity) Register(user authEntity.User) (*mongo.InsertOneResult, error) {

	collectionUser := e.DataBase.Collection("user")

	var existingUser authEntity.User
	collectionUser.FindOne(e.AppContext, bson.D{{"email", user.Email}}).Decode(&existingUser)

	if (existingUser != authEntity.User{}) {
		fmt.Println("[Auth Register] Email deja enregistré")
		return &mongo.InsertOneResult{}, errors.New("email deja enregistré")
	}

	fmt.Println("[Auth Register] Création du nouvel utilisateur")
	result, err := collectionUser.InsertOne(e.AppContext, user)
	if err != nil {
		panic(err)
	}
	return result, nil

}

func (e *AuthEntity) ResetPassword(email string, passwordHashed string) error {

	collectionUser := e.DataBase.Collection("user")

	var user authEntity.User
	collectionUser.FindOne(e.AppContext, bson.D{{"email", email}}).Decode(&user)

	if (user == authEntity.User{}) {
		return errors.New("utilisateur inexistant")
	}

	filter := bson.D{{"_id", user.ID}, {"useprovider", false}}
	update := bson.D{{"$set", bson.D{{"password", passwordHashed}}}}

	_, err := collectionUser.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	//if exist => continue else => ERROR
	//continue => change password

	return nil

}
