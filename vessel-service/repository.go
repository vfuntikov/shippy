package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/vfuntikov/shippy/vessel-service/proto/vessel"
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
	// ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cl, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	if err = cl.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("connected to the database")

	col := cl.Database("shippy").Collection("vessels")
	return &MongoRepository{client: cl, collection: col}, nil
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repo *MongoRepository) FindAvailable(ctx context.Context, spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	// Here we define a more complex query than our consignment-service's
	// GetAll function. Here we're asking for a vessel who's max weight and
	// capacity are greater than and equal to the given capacity and weight.
	// We're also using the `One` function here as that's all we want.
	err := repo.collection.Find(ctx, bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (repo *MongoRepository) Create(ctx context.Context, vessel *pb.Vessel) error {
	return repo.collection.Insert(ctx, vessel)
}
