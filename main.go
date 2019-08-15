package main

import (
	"container/list"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/model"
	"time"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.111.115:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Mongodb Connect ERROR!")
		panic(err)
	}
	collection := client.Database("test").Collection("test")
	l := list.New()
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	_, err = collection.InsertOne(ctx, model.Metadata{
		ID: "asdfasdf123",
		Name: "name",
		TotalSize: 13,
		TotalLength: 13,
		Enable: true,
		FileList: *l,
		Heat: 2,
		CreatedAt: 134,
		UpdatedAt: 135,
	})
	fmt.Println("hehe")
}
