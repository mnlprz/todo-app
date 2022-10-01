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

	default:
		log.Fatal("invalid param")
	}
}
