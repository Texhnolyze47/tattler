package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"tattler/model"
)

type RestaurantRepo struct {
	MongoCollection *mongo.Collection
}

// InsertRestaurant CRUD
// Create
func (r *RestaurantRepo) InsertRestaurant(modelRestaurant *model.Restaurant) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), modelRestaurant)

	if err != nil {
		return nil, err
	}

	return result, nil

}

// READ
func (rs *RestaurantRepo) FindRestaurantById(id string) (*model.Restaurant, error) {
	var result model.Restaurant
	err := rs.MongoCollection.FindOne(context.Background(),
		bson.D{{Key: "restaurant_id", Value: id}}).Decode(&result)

	if err != nil {
		return nil, err

	}

	return &result, nil
}

// READ ALL

func (rs *RestaurantRepo) FindAllRestaurants() (interface{}, error) {
	results, err := rs.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var restaurants []model.Restaurant

	err = results.All(context.Background(), &restaurants)
	if err != nil {
		return nil, err
	}

	return restaurants, nil

}

// Update

func (r *RestaurantRepo) UpdateRestaurantByID(restaurantId int, modelRestaurant *model.Restaurant) (int64, error) {
	result, err := r.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "restaurant_id", Value: restaurantId}},
		bson.D{{Key: "$set", Value: modelRestaurant}})
	if err != nil {
		return 0, nil
	}

	return result.ModifiedCount, nil
}

// delete
func (r *RestaurantRepo) DeleteRestaurant(restaurant_id string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "restaurant_id", Value: restaurant_id}})

	if err != nil {
		return 0, nil
	}

	return result.DeletedCount, nil
}
