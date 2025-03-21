package cmd

import (
	"fmt"
	"log/slog"
	"runtime/debug"

	"github.com/spf13/cobra"
	"github.com/vldcreation/helpme/pkg/profiler"
	"github.com/vldcreation/helpme/util"
)

type CMD interface {
	Execute() error
}

type App struct {
	root *cobra.Command
	prof *profiler.Profiler
}

var (
	Version string
	Commit  string
)

func printVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "(unknown)"
	}

	Version = info.Main.Version
	Commit = info.Main.Sum

	return fmt.Sprintf("%s\nCommit: %s", Version, Commit)
}

func NewApp() CMD {
	app := &App{}
	rootCmd := &cobra.Command{
		Use:   "helpme",
		Short: "Helpme CLI tool",
		Long:  `A CLI tool for finding and setting up development templates`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cpuProfile, _ := cmd.Flags().GetBool("cpuprofile")
			memProfile, _ := cmd.Flags().GetBool("memprofile")

			if cpuProfile || memProfile {
				prof, err := profiler.New("./profiles")
				if err != nil {
					util.PrintlnRed(err.Error())
					return
				}
				app.prof = prof

				if cpuProfile {
					if err := prof.StartCPUProfile(); err != nil {
						util.PrintlnRed(err.Error())
						return
					}
				}
			}
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if app.prof != nil {
				app.prof.StopCPUProfile()
				memProfile, _ := cmd.Flags().GetBool("memprofile")
				if memProfile {
					if err := app.prof.WriteHeapProfile(); err != nil {
						util.PrintlnRed(err.Error())
					}
				}
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			slog.Info("Root command executed successfully")
		},
	}

	rootCmd.Version = printVersion()
	rootCmd.PersistentFlags().BoolP("memprofile", "m", false, "enable memory profiling")
	rootCmd.PersistentFlags().BoolP("cpuprofile", "c", false, "enable cpu profiling")

	// Register subcommands
	rootCmd.AddCommand(
		NewFindCommand().Command(),
		NewSetupCommand().Command(),
		NewGeneratePasswordCommand().Command(),
		NewPullCommand().Command(),
		NewRunTestCommand().Command(),
		NewEncodeCommand().Command(),
	)

	app.root = rootCmd
	return app
}

func (app *App) Execute() error {
	if err := app.root.Execute(); err != nil {
		slog.Error("Error executing root command", "error", err)
		return err
	}
	return nil
}
