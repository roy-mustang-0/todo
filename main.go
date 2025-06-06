package main

import "fmt"

func main() {
	todo := &Todos{}
	if err := todo.LoadFromFile(); err != nil {
		fmt.Printf("File not loaded %v\n", err)
	}
	cmd := NewCmdFlags()
	cmd.Execute(todo)
	if err := todo.SaveToFile(); err != nil {
		fmt.Printf("error saving todos %v\n", err)
	}
}
