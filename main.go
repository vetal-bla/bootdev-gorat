package main

import (
	"fmt"
	"github.com/vetal-bla/bootdev-gorat/internal/config"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading file %w\n", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("testing some user")

	if err != nil {
		log.Fatalf("Cant set current_user: %w\n", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading file %w\n", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)

}
