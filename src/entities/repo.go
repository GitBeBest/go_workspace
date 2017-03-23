package entities

import (
	"fmt"
)

var currentId int
var todos Todos

func init()  {
	RepoCreateTodo(Todo{Name:"write presentation"})
	RepoCreateTodo(Todo{Name:"Host meetup"})
}

func RepoFindTodo(id int) Todo {
	for _,t := range todos {
		if t.Id == id {
			return t
		}
	}

	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error  {
	for id,t := range todos{
		if t.Id == id {
			todos = append(todos[:id], todos[id+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}

func GetTodo() Todos {
	return todos
}
