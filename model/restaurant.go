package model

import "time"

type Restaurant struct {
	Name         string    `json:"name" bson:"name"`
	Cuisine      string    `json:"cuisine" bson:"cuisine"`
	Borough      Address   `json:"borough" bson:"borough"`
	Grades       []Grade   `json:"grades" bson:"grades"`
	Comments     []Comment `json:"comments" bson:"comments"`
	RestaurantId string    `json:"restaurant_id" bson:"restaurant_id"`
}

type Address struct {
	Building string    `json:"building" bson:"building"`
	Coord    []float32 `json:"coord" bson:"coord"`
}

type Grade struct {
	Date  time.Time `json:"date" bson:"date"`
	Score float32   `json:"score" bson:"score"`
}

type Comment struct {
	Date    time.Time `json:"date" bson:"date"`
	Content string    `json:"content" bson:"content"`
}
