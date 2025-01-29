package main

import (
	"vldcreation/github.com/helpme/cmd"
)

func main() {
	app := cmd.NewApp()

	if err := app.Execute(); err != nil {
		panic(err)
	}
}
