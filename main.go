package main

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
	"tattler/service"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading file")
	}

	connect, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONG0_URL")))
	if err != nil {
		log.Print("connection error", err)
	}

	err = connect.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Print("ping fail", err)

	}

	defer func(connect *mongo.Client, ctx context.Context) {
		err := connect.Disconnect(ctx)
		if err != nil {

		}
	}(connect, context.Background())

	db := connect.Database("Tattler").Collection("restaurants")
	restaurantService := service.RestaurantService{MongoCollection: db}

	r := http.NewServeMux()

	r.HandleFunc("GET /health", healthHandler)

	r.HandleFunc("POST /api/restaurant", restaurantService.CreateRestaurantHandler)

	r.HandleFunc("GET /api/restaurant", restaurantService.GetAllRestaurantHandler)

	r.HandleFunc("GET /api/restaurant/{id}", restaurantService.GetRestaurantHandlerByID)

	r.HandleFunc("PUT /api/restaurant/{id}", restaurantService.UpdateRestaurantHandler)

	r.HandleFunc("DELETE /api/restaurant/{id}", restaurantService.DeleteRestaurantHandler)

	http.ListenAndServe(":4444", r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("running"))

}
