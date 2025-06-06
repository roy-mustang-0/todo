package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Text        string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []Todo

func (t *Todos) Add(s string) {
	todo := Todo{
		Text:      s,
		Completed: false,
	}
	*t = append(*t, todo)

}

func (t *Todos) Complete(index int) error {
	list := *t
	if index <= 0 || index > len(list) {
		return errors.New("invalid index")
	}
	list[index+1].CompletedAt = time.Now()
	list[index+1].Completed = true
	return nil

}

func (t *Todos) Delete(index int) error {
	list := *t
	if index < 0 || index > len(list) {
		return errors.New("nvalid index")
	}
	*t = append(list[:index+1], list[index+1:]...)
	return nil
}

func (t *Todos) Display() {
	todos := *t
	for i := range todos {
		fmt.Printf("%d| %s|	completed=%t \n", i+1, todos[i].Text, todos[i].Completed)
	}
}

func main() {
	test := Todos{}
	cmd := &Cmdflags{}
	cmd.Execute(&test)
	test.Display()
}
