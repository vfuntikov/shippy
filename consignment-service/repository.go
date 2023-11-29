package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/vfuntikov/shippy/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoRepository implementation
type MongoRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
	// ctx        context.Context
}

// CreateClient -
func NewMongoRepository() (*MongoRepository, error) {
	// Database host from the environment variables
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost:27017"
	}

	uri := fmt.Sprintf("mongodb://admin:mongo@%s", host)
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cl, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	if err = cl.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("connected to the database")

	col := cl.Database("shippy").Collection("consignment")
	return &MongoRepository{client: cl, collection: col}, nil
}

// Create -
func (repository *MongoRepository) Create(ctx context.Context, consignment *pb.Consignment) error {
	_, err := repository.collection.InsertOne(ctx, consignment)
	return err
}

// GetAll -
func (repository *MongoRepository) GetAll(ctx context.Context) ([]*pb.Consignment, error) {
	cur, err := repository.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var consignments []*pb.Consignment
	for cur.Next(ctx) {
		var consignment *pb.Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}
