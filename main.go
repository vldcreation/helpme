package main

import (
	"vldcreation/github.com/helpme/cmd"
)

func main() {
	if err := cmd.NewApp().Execute(); err != nil {
		panic(err)
	}
}
