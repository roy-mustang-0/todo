package main

import (
	"flag"
	"fmt"
)

type Cmdflags struct {
	Add    string
	Toggle int
	List   bool
	Delete int
}

func NewCmdFlags() *Cmdflags {
	cf := Cmdflags{}
	flag.StringVar(&cf.Add, "add", "", "add a new todo")
	flag.IntVar(&cf.Toggle, "toggle", -1, "mark as completed")
	flag.BoolVar(&cf.List, "list", false, "List all the todos")
	flag.IntVar(&cf.Delete, "del", -1, "Delete a todo entry")

	flag.Parse()

	return &cf
}
func (cf *Cmdflags) Execute(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	case cf.List:
		todos.Display()
	case cf.Delete != -1:
		todos.Delete(cf.Delete)
	default:
		fmt.Println("Usage: todo [options...]\n--add	To add todo\n--toggle To mark todo as completed\n--list	To list all the current todos\n--del To delete a todo entry")
	}

}
