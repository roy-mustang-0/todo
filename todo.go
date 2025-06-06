package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Text        string    `json:"text"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

type Todos []Todo

const file = "todos.json"

func (t *Todos) LoadFromFile() error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	if len(data) == 0 {
		return nil
	}
	err = json.Unmarshal(data, t)
	if err != nil {
		return fmt.Errorf("error reading file : %v", err)
	}
	return nil

}

func (t *Todos) SaveToFile() error {
	data, err := json.Marshal(*t)
	if err != nil {
		return fmt.Errorf("error marshaling todos %v", err)
	}
	err = os.WriteFile(file, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %v", err)
	}
	return nil
}

func (t *Todos) Add(s string) {
	todo := Todo{
		Text:      s,
		Completed: false,
		CreatedAt: time.Now(),
	}
	*t = append(*t, todo)
	fmt.Printf("%s added successfully\n", s)
}

func (t *Todos) Toggle(index int) error {
	list := *t

	if index <= 0 || index > len(list) {
		return errors.New("invalid index")
	}
	list[index-1].CompletedAt = time.Now()
	list[index-1].Completed = !list[index-1].Completed
	return nil

}

func (t *Todos) Delete(index int) error {
	list := *t
	if index < 1 || index > len(list) {
		return errors.New("invalid index")
	}
	*t = append(list[:index-1], list[index:]...)
	return nil
}

func (t *Todos) Display() {
	todos := *t
	if len(todos) == 0 {
		fmt.Println("Todo list is empty")
		return
	}
	for i := range todos {
		fmt.Printf("%d: %s --> completed=%t \n", i+1, todos[i].Text, todos[i].Completed)
	}
}
