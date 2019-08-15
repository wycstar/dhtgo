package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/config"
	"time"
)

func getClient() {
	client, err := mongo.NewClient(options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%d",
			config.CONFIG.Mongo.Host,
			config.CONFIG.Mongo.Port)))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer func() {
		err := client.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		panic(err)
	}
	collection := client.Database(config.CONFIG.Mongo.Database).Collection(config.CONFIG.Mongo.Collection)
	_, err = collection.InsertOne(ctx, Metadata{
		ID:          "asdfasdf123",
		Name:        "name",
		TotalSize:   13,
		TotalLength: 13,
		Enable:      true,
		FileList:    []File{{Name: "321", Size: 16}},
		Heat:        2,
		CreatedAt:   134,
		UpdatedAt:   135,
	})
}
