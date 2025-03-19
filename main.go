package main

import (
	"github.com/vldcreation/helpme/cmd"
	"github.com/vldcreation/helpme/util"
)

func main() {
	if err := cmd.NewApp().Execute(); err != nil {
		util.PrintlnRed(err.Error())
	}
}
