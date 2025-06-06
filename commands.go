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
	cmd := Cmdflags{}
	flag.StringVar(&cmd.Add, "add", "", "add a new todo")
	flag.IntVar(&cmd.Toggle, "toggle", -1, "add a new todo")
	flag.BoolVar(&cmd.List, "list", false, "add a new todo")
	flag.IntVar(&cmd.Delete, "delete", -1, "add a new todo")

	flag.Parse()

	return &cmd
}
func (cf *Cmdflags) Execute(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Toggle != -1:
		todos.Complete(cf.Toggle)
	case cf.List:
		todos.Display()
	case cf.Delete != -1:
		todos.Delete(cf.Delete)
	default:
		fmt.Println("invalid Flag")
	}

}
