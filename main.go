package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/vetal-bla/bootdev-gorat/internal/config"
	"github.com/vetal-bla/bootdev-gorat/internal/database"
	"log"
	"os"
)

type state struct {
	db    *database.Queries
	state *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading file %w\n", err)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("Cant open database", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db:    dbQueries,
		state: &cfg,
	}

	c := commands{
		registeredCmds: make(map[string]func(*state, command) error),
	}

	// register command by name with handler function
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerGetUsres)
	c.register("agg", handlerAgg)
	c.register("addfeed", handlerAddFeed)

	args := os.Args

	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	// now you can run command which you registered above
	err = c.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

}
