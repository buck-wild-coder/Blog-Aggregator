package main

import (
	"fmt"
	"os"

	"github.com/buck/blog/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name     string
	argument []string
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Can't read")
		return
	}
	stateVar := state{&cfg}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	if len(os.Args) < 3 {
		fmt.Println("Must contain atleast 2 commands")
		os.Exit(1)
	}
	cmdname := os.Args[1]
	arguments := os.Args[2]
	cmds.register(cmdname, handlerLogin)
	cmdVar := command{
		name:     cmdname,
		argument: make([]string, 5),
	}
	cmdVar.argument = []string{arguments}
	cmds.run(&stateVar, cmdVar)
	// fmt.Println(cfg)
}
