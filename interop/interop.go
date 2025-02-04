package main

import "github.com/vldcreation/go-ressources/interop"

func main() {
	runner := interop.NewInteropRunner(interop.Interop{
		Language: "javascript",
		FilePath: "javascript/ai.js",
	})

	out, err := runner.Run()
	if err != nil {
		panic(err)
	}

	println(out)
}
