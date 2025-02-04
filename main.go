package main

import (
	"vldcreation/github.com/helpme/cmd"
)

func main() {
	if err := cmd.NewApp().Execute(); err != nil {
		panic(err)
	}

	// interopOb := interop.NewInterop("javascript", "interop/javascript/ai.js")
	// output, err := interopOb.Run()
	// if err != nil {
	// 	panic(err)
	// }

	// println("Output:")
	// println(output.String())
}
