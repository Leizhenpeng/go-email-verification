package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection

func InitUserCollection() {
	userCollection = client.Database("email-verifications").Collection("users")
	SetEmailUniqueIndex()
}

func GetUserCollection() *mongo.Collection {
	return userCollection
}

func SetEmailUniqueIndex() {
	userCollection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    bson.M{"Email": 1},
		Options: options.Index().SetUnique(true),
	})
}
func AddUser(ctx context.Context, user *UserInfo) (*mongo.InsertOneResult, error) {
	return userCollection.InsertOne(ctx, user)
}

func GetUserByID(ctx context.Context, id interface{}) (*mongo.SingleResult, error) {
	return userCollection.FindOne(ctx, map[string]interface{}{"_id": id}), nil
}

func GetUserByEmail(ctx context.Context, email string) (*mongo.SingleResult, error) {
	return userCollection.FindOne(ctx, map[string]interface{}{"Email": email}), nil
}

func UpdateUser(ctx context.Context, email string, update interface{}) (*mongo.UpdateResult, error) {
	return userCollection.UpdateOne(ctx, map[string]interface{}{"Email": email}, update)
}

func DeleteUser(ctx context.Context, email string) (*mongo.DeleteResult, error) {
	return userCollection.DeleteOne(ctx, map[string]interface{}{"Email": email})
}
