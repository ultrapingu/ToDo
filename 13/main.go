package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

type Item struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Status string    `json:"status"`
}

func display(todoList ...Item) {
	for _, todo := range todoList {
		fmt.Printf("%#v\n", todo)
	}
}

func toJson(todoList ...Item) (string, error) {
	result, err := json.Marshal(todoList)
	return string(result), err
}

func readTodoListFromFile(filename string) ([]Item, error) {
	var result []Item

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return result, errors.New(fmt.Sprintf("Failed to read file: %s.  %s", filename, err.Error()))
	}

	err = json.Unmarshal(fileContent, &result)
	if err != nil {
		return result, errors.New(fmt.Sprintf("Failed to parse json from file: %s.  %s", filename, err.Error()))
	}

	return result, nil
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		log.Fatal("Unexpected arguments, must provide 1 filename")
	}

	todoList, err := readTodoListFromFile(args[0])
	if err != nil {
		log.Fatal(err.Error())
	}

	display(todoList...)
}
