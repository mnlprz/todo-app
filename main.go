package main

import (
	"flag"
	"log"

	"github.com/mnlprz/todo-app/todo"
)

const (
	todoFile = ".todos.json"
)

func main() {

	add := flag.Bool("add", false, "add a new todo task")
	complete := flag.Int("complete", 0, "complete a task")
	delete := flag.Int("delete", 0, "delete a task")
	list := flag.Bool("list", false, "list tasks")

	flag.Parse()

	todos := &todo.Todos{}

	err := todos.Load(todoFile)
	if err != nil {
		log.Fatal(err)
	}

	switch {

	case *add:
		todos.Add("example1")
		err := todos.Store(todoFile)
		if err != nil {
			log.Fatal(err)
		}

	case *complete > 0:
		err = todos.Complete(*complete)
		if err != nil {
			log.Fatal(err)
		}
		err := todos.Store(todoFile)
		if err != nil {
			log.Fatal(err)
		}

	case *delete > 0:
		err = todos.Delete(*delete)
		if err != nil {
			log.Fatal(err)
		}
		err := todos.Store(todoFile)
		if err != nil {
			log.Fatal(err)
		}

	case *list:
		todos.List()

	default:
		log.Fatal("invalid param")
	}
}
