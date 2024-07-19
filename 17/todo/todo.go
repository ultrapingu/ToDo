package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

const (
	Pending    = "pending"
	InProgress = "inprogress"
	Done       = "done"
)

type Item struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Status string    `json:"status"`
}

type List struct {
	items []Item
}

func NewList() List {
	return List{items: make([]Item, 0)}
}

func Add(list List, item Item) (List, Item) {
	item = Item{ID: uuid.New(), Title: item.Title, Status: item.Status}
	list.items = append(list.items, item)

	return list, item
}

func Update(list List, item Item) (List, error) {
	for i, curr := range list.items {
		if curr.ID == item.ID {
			list.items[i] = item
			return list, nil
		}
	}

	return list, errors.New("Unable to find item with id: " + item.ID.String())
}

func GetIdx(list List, id uuid.UUID) int {
	for i, item := range list.items {
		if item.ID == id {
			return i
		}
	}

	return -1
}

func Delete(list List, id uuid.UUID) (List, error) {

	if idx := GetIdx(list, id); idx != -1 {
		list.items = append(list.items[:idx], list.items[idx+1:]...)
		return list, nil
	}

	return list, fmt.Errorf("unable to find item with id: %s", id.String())
}

func Save(list List, filename string) error {
	data, err := json.Marshal(list.items)
	if err != nil {
		return fmt.Errorf("failed to serialize to json: %s", err.Error())
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %s.  %s", filename, err.Error())
	}

	return nil
}

func Load(filename string) (List, error) {
	list := List{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return list, fmt.Errorf("failed to read file: %s.  %s", filename, err.Error())
	}

	err = json.Unmarshal(data, &list.items)
	if err != nil {
		return list, fmt.Errorf("failed to parse json from file: %s.  %s", filename, err.Error())
	}

	return list, nil
}

func (l *List) GetItems() []Item {
	return l.items
}

func ParseStatus(s string) (string, error) {
	switch strings.ToLower(s) {
	case "1", Pending:
		return Pending, nil
	case "2", InProgress:
		return InProgress, nil
	case "3", Done:
		return Done, nil
	default:
		return "", fmt.Errorf("invalid input for todo item status: %s\n", s)
	}
}
