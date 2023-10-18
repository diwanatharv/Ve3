package config

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Makemongoserver() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Error(err.Error())
		log.Error("error in connect to localhost mongodb")
	}
	db := client.Database("Practice")
	return db
}

func Makeredisserever() *redis.Client {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	return rdb
}
