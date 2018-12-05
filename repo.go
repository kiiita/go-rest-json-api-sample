package main

import "fmt"

var (
	todos     Todos
	currentID int
)

func init() {
	RepoCreateTodo(Todo{Name: "Write Presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) Todo {
	// for _, t := range todos {
	// 	if t.ID == id {
	// 		return t
	// 	}
	// }
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	// db.Create(&Todo{Name: "hoge"})
	return t
}

func RepoDestroyTodo(id int) error {
	// for i, t := range todos {
	// 	if t.ID == id {
	// 		todos = append(todos[:i], todos[i+1:]...)
	// 		return nil
	// 	}
	// }

	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
