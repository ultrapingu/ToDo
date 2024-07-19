package main

import (
	"encoding/json"
	"log"
	"os"

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
	args := os.Args[1:]

	if len(args) != 1 {
		log.Fatal("Unexpected arguments, must provide 1 filename")
	}

	washDishes := Item{ID: uuid.New(), Title: "Wash Dishes", Status: "InProgress"}
	cleanCar := Item{ID: uuid.New(), Title: "Clean Car", Status: "Pending"}
	watchTV := Item{ID: uuid.New(), Title: "Watch TV", Status: "Done"}

	jsonStr, err := toJson(washDishes, cleanCar, watchTV)
	if err != nil {
		log.Fatal("Failed to serialize to json")
	}

	jsonByte := []byte(jsonStr)
	err = os.WriteFile(args[0], jsonByte, 0644)
	if err != nil {
		log.Fatal("Failed to write to file")
	}
}
