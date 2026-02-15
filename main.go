package main

import (
	"fmt"
	"github.com/vetal-bla/bootdev-gorat/internal/config"
	"os"
)

func main() {
	cfg := config.Config{}
	cfg, err := config.Read()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfg.SetUser("testing5432")

	fmt.Println(cfg)

}
