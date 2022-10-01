package todo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {

	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {

	list := *t
	if index <= 0 || index > len(list) {
		return errors.New("invalid index")
	}
	list[index-1].Done = true
	list[index-1].CompletedAt = time.Now()
	return nil
}

func (t *Todos) Delete(index int) error {

	list := *t
	if index <= 0 || index > len(list) {
		return errors.New("invalid index")
	}
	*t = append(list[:index-1], list[index:]...)
	return nil
}

func (t *Todos) Load(filename string) error {

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todos) Store(filename string) error {

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
