package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"leizhenpeng/go-email-verification/initialize"
	"leizhenpeng/go-email-verification/models"
)

var userCollection *mongo.Collection

func InitUserCollection() {
	userCollection = initialize.Client.Database("email-verifications").Collection("users")
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
func AddUser(ctx context.Context, user *models.UserInfo) (*mongo.InsertOneResult, error) {
	return userCollection.InsertOne(ctx, user)
}

func GetUserByID(ctx context.Context, id interface{}) (*mongo.SingleResult, error) {
	return userCollection.FindOne(ctx, map[string]interface{}{"_id": id}), nil
}

func GetUserByEmail(ctx context.Context, email string) (*mongo.SingleResult, error) {
	return userCollection.FindOne(ctx, map[string]interface{}{"Email": email}), nil
}

func UpdateUserFieldById(ctx context.Context, id interface{}, field string, value interface{}) (*mongo.UpdateResult, error) {
	return userCollection.UpdateOne(ctx, map[string]interface{}{"_id": id}, map[string]interface{}{"$set": map[string]interface{}{field: value}})
}

func UpdateUserFieldByEmail(ctx context.Context, email string, field string, value interface{}) (*mongo.UpdateResult, error) {
	return userCollection.UpdateOne(ctx, map[string]interface{}{"Email": email}, map[string]interface{}{"$set": map[string]interface{}{field: value}})
}
func DeleteUser(ctx context.Context, email string) (*mongo.DeleteResult, error) {
	return userCollection.DeleteOne(ctx, map[string]interface{}{"Email": email})
}
