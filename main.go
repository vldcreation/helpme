package main

import (
	"log"

	"github.com/vldcreation/helpme/cmd"
)

const (
	RED_COLOR    = "\033[0;31m"
	BLUE_COLOR   = "\033[0;34m"
	YELLOW_COLOR = "\033[0;33m"
	RESET_COLOR  = "\033[0m"
)

var Version string

func main() {
	log.Printf("Version: %s", Version)
	if err := cmd.NewApp().Execute(); err != nil {
		println(string(RED_COLOR), err.Error(), string(RESET_COLOR))
	}
}
