package models

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"leizhenpeng/go-email-verification/config"
)

var client *mongo.Client
var redisClient *redis.Client

func initMongoClient() (err error) {
	ctx := context.TODO()
	conn := options.Client().ApplyURI(config.GetConfig().DbUrI)
	client, err = mongo.Connect(ctx, conn)
	if err != nil {
		return err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB!")
	return err
}

func initRedisClient() (err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.GetConfig().RedisUrI,
		Password: config.GetConfig().RedisPass,
		DB:       config.GetConfig().RedisDb,
	})
	_, err = redisClient.Ping(context.TODO()).Result()
	if err != nil {
		return err
	}
	fmt.Println("Connected to Redis!")
	return err
}

func InitClient() {
	err := initMongoClient()
	if err != nil {
		panic(err)
	}
	err = initRedisClient()
	if err != nil {
		panic(err)
	}
}

func CloseClient() {
	_ = client.Disconnect(context.Background())
	_ = redisClient.Close()
}
