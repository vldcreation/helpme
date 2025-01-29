package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type SetupCommand struct {
	cmd *cobra.Command

	// flags
	name     string
	template string
}

func NewSetupCommand() *SetupCommand {
	apps := &SetupCommand{}
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "Setup project templates",
	}

	cmd.PersistentFlags().StringVarP(&apps.name, "name", "n", "rest", "Starter template to generate")
	cmd.PersistentFlags().StringVarP(&apps.template, "template", "t", "default", "Template name to use")

	cmd.MarkPersistentFlagRequired("lang")

	apps.cmd = cmd
	return apps
}

func (c *SetupCommand) Command() *cobra.Command {
	c.cmd.Run = c.Execute
	return c.cmd
}

func (c *SetupCommand) Execute(_ *cobra.Command, args []string) {
	fmt.Printf("Creating Project Template")
}
