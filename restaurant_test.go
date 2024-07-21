package main

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"tattler/model"
	"tattler/repository"
	"testing"
	"time"
)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb+srv://tattler-admin:g3Fq8DqFoxnZ1Nrg@cluster0.5bv3dft.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))

	if err != nil {
		log.Fatal("error while connection", err)
	}

	log.Print("connection success")

	err = mongoTestClient.Ping(context.Background(),
		readpref.Primary())

	if err != nil {
		log.Fatal("ping fail")
	}

	log.Print("ping success")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	restaurant1 := uuid.New().String()
	restaurant2 := uuid.New().String()

	connec := mongoTestClient.Database("Tattler").Collection("restaurants_test")

	restRepo := repository.RestaurantRepo{MongoCollection: connec}

	t.Run("Insert Restaurant 1", func(t *testing.T) {
		rest := model.Restaurant{
			Name:    "Restaurante 1",
			Cuisine: "Indian",
			Borough: model.Address{
				Building: "ESCIHU",
				Coord:    []float32{32.4, 23.21},
			},
			Grades: []model.Grade{
				{Date: time.Now()},
				{Score: 4.00},
			},
			Comments: []model.Comment{
				{Date: time.Now()},
				{Content: "Good place"},
			},
			RestaurantId: restaurant1,
		}
		result, err := restRepo.InsertRestaurant(&rest)

		if err != nil {
			t.Fatal("operation failed", err)
		}

		t.Log("operation success", result)
	})

	t.Run("Insert Restaurant 2", func(t *testing.T) {
		rest := model.Restaurant{
			Name:    "Restaurante 2",
			Cuisine: "Mexican",
			Borough: model.Address{
				Building: "ESCIHU",
				Coord:    []float32{32.4, 23.21},
			},
			Grades: []model.Grade{
				{Date: time.Now()},
				{Score: 4.00},
			},
			Comments: []model.Comment{
				{Date: time.Now()},
				{Content: "Good place"},
			},
			RestaurantId: restaurant2,
		}
		result, err := restRepo.InsertRestaurant(&rest)

		if err != nil {
			t.Fatal("operation failed", err)
		}

		t.Log("operation success", result)
	})

	t.Run("Get Restaurant 1", func(t *testing.T) {
		result, err := restRepo.FindRestaurantById(restaurant1)

		if err != nil {
			t.Fatal("get operation failed", err)
		}

		t.Log("Rest 1", result.Name)
	})

	t.Run("Get All Restaurants", func(t *testing.T) {
		result, err := restRepo.FindAllRestaurants()

		if err != nil {
			t.Fatal("get all operation failed")
		}

		t.Log("restaurants", result)
	})

}
