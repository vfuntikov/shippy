package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	mongoDB, err := NewMongoRepository()
	if err != nil {
		log.Fatal(err)
	}
	defer mongoDB.client.Disconnect(ctx)

	createDummyData(mongoDB)

	service := NewService(mongoDB)
	if err := service.Serve(); err != nil {
		log.Fatal(err)
	}
}
