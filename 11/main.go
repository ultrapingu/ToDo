package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

type Item struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Status string    `json:"status"`
}

func toJson(todoList ...Item) (string, error) {
	result, err := json.Marshal(todoList)
	return string(result), err
}

func main() {
	washDishes := Item{ID: uuid.New(), Title: "Wash Dishes", Status: "InProgress"}
	cleanCar := Item{ID: uuid.New(), Title: "Clean Car", Status: "Pending"}
	watchTV := Item{ID: uuid.New(), Title: "Watch TV", Status: "Done"}

	result, err := toJson(washDishes, cleanCar, watchTV)
	if err != nil {
		log.Fatal("Failed to serialize to json")
	}

	log.Println(result)
}
