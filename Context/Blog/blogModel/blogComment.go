package blogModel

import (
	"root/core/model"

	"go.mongodb.org/mongo-driver/mongo"
)

//IMPORTANT to create a struct on Entity => to add custom method for this model
type BlogCommentEntity model.Entity

type Comment struct {
	User    string
	Message string
}

var collectionName string = "blogComment"
var collectionComment *mongo.Collection

func (e *BlogCommentEntity) CreateComment(comment Comment) *mongo.InsertOneResult {
	collectionComment := e.DataBase.Collection(collectionName)
	result, err := collectionComment.InsertOne(e.AppContext, comment)
	if err != nil {
		panic(err)
	}
	return result
}
