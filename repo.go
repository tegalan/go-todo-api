package main

import "fmt"

var currentID int

var todos Todos

func init() {
	RepoCreateTodo(Todo{Name: "Write Golang REST API"})
	RepoCreateTodo(Todo{Name: "Upload the code to Github"})
}

func RepoCreateTodo(t Todo) Todo {
	currentID += 1
	t.ID = currentID
	todos = append(todos, t)

	return t
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}

	// Return empty todo if not found
	return Todo{}
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Couldn't find todo with id %d", id)
}
