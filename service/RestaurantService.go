package service

import (
	"encoding/json"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
	"tattler/model"
	"tattler/repository"
)

type RestaurantService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (service *RestaurantService) CreateRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	// CreateEmployeeHandler creates a new employee
	w.Header().Set("Content-Type", "application/json")

	res := &Response{}

	// Decode the request body
	defer json.NewEncoder(w).Encode(res)

	var restaurant model.Restaurant

	err := json.NewDecoder(r.Body).Decode(&restaurant)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	restaurant.RestaurantId = uuid.NewString()

	repo := repository.RestaurantRepo{MongoCollection: service.MongoCollection}

	insertID, err := repo.InsertRestaurant(&restaurant)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Data = restaurant.RestaurantId
	w.WriteHeader(http.StatusCreated)

	log.Print("Inserted document with ID: ", insertID)

}

func (service *RestaurantService) GetRestaurantHandlerByID(w http.ResponseWriter, r *http.Request) {
	// GetEmployeeHandlerByID gets an employee by ID
	w.Header().Set("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	restaurantId := r.URL.Query().Get("restaurant_id")

	repo := repository.RestaurantRepo{MongoCollection: service.MongoCollection}

	restaurant, err := repo.FindRestaurantById(restaurantId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Data = restaurant
	w.WriteHeader(http.StatusOK)

}
func (service *RestaurantService) GetAllRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	// GetAllEmployeeHandler gets all employees
	w.Header().Set("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	repo := repository.RestaurantRepo{MongoCollection: service.MongoCollection}

	restaurants, err := repo.FindAllRestaurants()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Data = restaurants
	w.WriteHeader(http.StatusOK)
}
func (service *RestaurantService) UpdateRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	// UpdateEmployeeHandler updates an employee
	w.Header().Set("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	var restaurant model.Restaurant

	id := r.PathValue("id")

	err := json.NewDecoder(r.Body).Decode(&restaurant)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	repo := repository.RestaurantRepo{MongoCollection: service.MongoCollection}

	restaurant.RestaurantId = id

	idInt, err := strconv.Atoi(id)
	if err != nil {
		// handle the error
		w.WriteHeader(http.StatusBadRequest)
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	cout, err := repo.UpdateRestaurantByID(idInt, &restaurant)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return

	}

	res.Data = cout

	w.WriteHeader(http.StatusOK)
}
func (service *RestaurantService) DeleteRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	// DeleteEmployeeHandler deletes an employee
	w.Header().Set("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	id := r.PathValue("id")

	repo := repository.RestaurantRepo{MongoCollection: service.MongoCollection}

	count, err := repo.DeleteRestaurant(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
}
