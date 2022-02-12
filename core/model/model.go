package model

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//type DataBaseSessionManager struct{
//	DataBaseSession []*DataBaseSession
//}

//Il pourrait y avoir plusieurs sessions avec un session manager qui serait dans le Global
type DataBase struct {
	DB_URI  string
	DB_NAME string
}

//Faire une fonction de connexion qui retoure l'accès au methode de la DB ( mongo / sql / ... )

func Connexion(ctx context.Context, DB_NAME string, DB_URI string) *mongo.Database {
	clientOptions := options.Client().ApplyURI(DB_URI)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}
	return client.Database(DB_NAME)
}

type Entity struct {
	DataBase   *mongo.Database
	AppContext context.Context
}
