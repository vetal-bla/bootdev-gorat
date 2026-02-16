package main

import (
	"github.com/vetal-bla/bootdev-gorat/internal/config"
	"log"
	"os"
)

type state struct {
	state *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading file %w\n", err)
	}
	programState := &state{
		state: &cfg,
	}

	c := commands{
		registeredCmds: make(map[string]func(*state, command) error),
	}

	c.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	err = c.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

}
