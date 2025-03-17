package cmd

import (
	"fmt"
	"log/slog"
	"runtime/debug"

	"github.com/spf13/cobra"
)

type CMD interface {
	Execute() error
}

type App struct {
	root *cobra.Command
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
	rootCmd := &cobra.Command{
		Use:   "helpme",
		Short: "Helpme CLI tool",
		Long:  `A CLI tool for finding and setting up development templates`,
		Run: func(cmd *cobra.Command, args []string) {
			slog.Info("Root command executed successfully")
		},
	}

	rootCmd.Version = printVersion()

	// Register subcommands
	rootCmd.AddCommand(
		NewFindCommand().Command(),
		NewSetupCommand().Command(),
		NewGeneratePasswordCommand().Command(),
		NewPullCommand().Command(),
		NewRunTestCommand().Command(),
		NewEncodeCommand().Command(),
	)

	return &App{
		root: rootCmd,
	}
}

func (app *App) Execute() error {
	if err := app.root.Execute(); err != nil {
		slog.Error("Error executing root command", "error", err)
		return err
	}
	return nil
}
