package main

import (
	"flag"

	"github.com/vldcreation/helpme/cmd"
	"github.com/vldcreation/helpme/pkg/profiler"
	"github.com/vldcreation/helpme/util"
)

func main() {
	// Setup profiling flags
	cpuProfile := flag.String("cpuprofile", "", "write cpu profile to file")
	memProfile := flag.String("memprofile", "", "write memory profile to file")
	flag.Parse()

	// Initialize profiler if profiling is enabled
	if *cpuProfile != "" || *memProfile != "" {
		prof, err := profiler.New("./profiles")
		if err != nil {
			util.PrintlnRed(err.Error())
			return
		}

		// Start CPU profiling if requested
		if *cpuProfile != "" {
			if err := prof.StartCPUProfile(); err != nil {
				util.PrintlnRed(err.Error())
				return
			}
			defer prof.StopCPUProfile()
		}

		// Write memory profile at exit if requested
		if *memProfile != "" {
			defer func() {
				if err := prof.WriteHeapProfile(); err != nil {
					util.PrintlnRed(err.Error())
				}
			}()
		}
	}

	if err := cmd.NewApp().Execute(); err != nil {
		util.PrintlnRed(err.Error())
	}
}
