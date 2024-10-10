package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Edit   string
	Del    int
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title.")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Specify a todo by index to delete.")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle.")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
		todos.print()
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid format for edit. Use id:new_title")
			os.Exit(1)
		}

		idx, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index provided.")
			os.Exit(1)
		}

		error := todos.edit(idx, parts[1])
		if error != nil {
			fmt.Println("Error: Invalid index. Please provide an index of a valid todo item.")
			os.Exit(1)
		}
		todos.print()
	case cf.Toggle != -1:
		err := todos.toggle(cf.Toggle)
		if err != nil {
			fmt.Println("Error: Invalid index. Please provide an index of a valid todo item.")
			os.Exit(1)
		}
		todos.print()
	case cf.Del != -1:
		err := todos.delete(cf.Del)
		if err != nil {
			fmt.Println("Error: Invalid index. Please provide an index of a valid todo item.")
			os.Exit(1)
		}
		todos.print()
	default:
		todos.print()
	}
}
