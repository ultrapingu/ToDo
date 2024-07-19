package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Item struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Status string    `json:"status"`
}

func myVariadic(todos ...Item) {
	for _, todo := range todos {
		fmt.Printf("%#v\n", todo)
	}
}

func main() {
	washDishes := Item{ID: uuid.New(), Title: "Wash Dishes", Status: "InProgress"}
	cleanCar := Item{ID: uuid.New(), Title: "Clean Car", Status: "Pending"}
	watchTV := Item{ID: uuid.New(), Title: "Watch TV", Status: "Done"}

	myVariadic(washDishes, cleanCar, watchTV)
}
