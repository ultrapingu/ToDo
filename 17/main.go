package main

import (
	"bufio"
	"fmt"
	"log"
	"main/todo"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

var todoList todo.List
var reader *bufio.Reader
var running bool

type cmd func()
type Command struct {
	name string
	fn   cmd
}

func readTrimmedStr(msg string) string {
	log.Println(msg)
	command, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read input!")
	}

	return strings.TrimSpace(command)
}

func readCommand() string {
	command := readTrimmedStr("")
	command = strings.Split(command, " ")[0]
	command = strings.Replace(command, ")", "", -1)
	command = strings.ToLower(command)

	return command
}

func readUUID(msg string) uuid.UUID {
	for {
		input := readTrimmedStr(msg)
		id, err := uuid.Parse(input)
		if err == nil {
			return id
		}

		log.Println(err.Error())
	}
}

func handleCommandInput(cmds []Command) {
	fmt.Println("What would you like to do?")
	for i, cmd := range cmds {
		fmt.Printf("\t%d) %s\n", i+1, cmd.name)
	}

	inputCmd := readCommand()
	for i, cmd := range cmds {
		if inputCmd == strconv.Itoa(i+1) || inputCmd == strings.Split(cmd.name, " ")[0] {
			cmd.fn()
			return
		}
	}

	fmt.Printf("Unknown command: '%s'\n", inputCmd)
}

func loadFile() {
	filename := readTrimmedStr("Please enter the filename")

	result, err := todo.Load(filename)
	if err != nil {
		log.Printf("Failed to load file: %s.  %s\n", filename, err.Error())
		return
	}

	todoList = result
}

func saveFile() {
	filename := readTrimmedStr("Please enter the filename")

	if err := todo.Save(todoList, filename); err != nil {
		log.Printf("Failed to save file: %s.  %s\n", filename, err.Error())
		return
	}
}

func listItems() {
	fmt.Printf("%d items in todo list\n", len(todoList.GetItems()))
	log.Println()

	for _, item := range todoList.GetItems() {
		fmt.Println("========================================")
		fmt.Printf("ID: %s\n", item.ID.String())
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Status: %s\n", item.Status)
		fmt.Println()
	}
}

func createItem() {
	title := readTrimmedStr("Please enter the title for the todo item")

	statusStr := readTrimmedStr("Please enter the status for the todo item")
	status, err := todo.ParseStatus(statusStr)
	if err != nil {
		log.Println("Failed to parse input!")
		return
	}

	updatedList, newItem := todo.Add(todoList, todo.Item{ID: uuid.New(), Title: title, Status: status})
	todoList = updatedList

	fmt.Printf("Added new item with id %s", newItem.ID.String())
}

func updateItem() {
	id := readUUID("Please enter the item's id")

	if idx := todo.GetIdx(todoList, id); idx == -1 {
		log.Printf("Unknown item with id: %s\n", id.String())
		return
	}

	title := readTrimmedStr("Please enter the updated title for the todo item")

	statusStr := readTrimmedStr("Please enter the updated status for the todo item")
	status, err := todo.ParseStatus(statusStr)
	if err != nil {
		log.Println("Failed to parse input!")
		return
	}

	updatedList, err := todo.Update(todoList, todo.Item{ID: id, Title: title, Status: status})
	if err != nil {
		log.Printf("Failed to update list: %s\n", err.Error())
		return
	}

	todoList = updatedList
	fmt.Printf("Updated item with id %s", id.String())
}

func deleteItem() {
	id := readUUID("Please enter the item's id")

	newList, err := todo.Delete(todoList, id)
	if err != nil {
		log.Println(err.Error())
		return
	}

	todoList = newList
}

func exit() {
	running = false
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	todoList = todo.NewList()

	cmds := []Command{
		{name: "Load file", fn: loadFile},
		{name: "Save file", fn: saveFile},
		{name: "List Items", fn: listItems},
		{name: "Create Item", fn: createItem},
		{name: "Update Item", fn: updateItem},
		{name: "Delete Item", fn: deleteItem},
		{name: "Exit", fn: exit},
	}

	for running = true; running; {
		handleCommandInput(cmds)
	}
}
