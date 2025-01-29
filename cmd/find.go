package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type FindCommand struct {
	cmd *cobra.Command

	// flags
	lang string
	pkg  string
	save bool
	exec bool
}

func NewFindCommand() *FindCommand {
	apps := &FindCommand{}
	cmd := &cobra.Command{
		Use:   "find",
		Short: "Find an example for a given function",
		Long:  "Find an example for a given function in a given language",
		Args:  cobra.ExactArgs(1),
	}

	cmd.PersistentFlags().StringVarP(&apps.lang, "lang", "l", "", "Language to search (go/javascript)")
	cmd.PersistentFlags().StringVarP(&apps.pkg, "pkg", "p", "", "Package name (optional)")
	cmd.PersistentFlags().BoolVarP(&apps.save, "save", "s", false, "Save example to a file")
	cmd.PersistentFlags().BoolVarP(&apps.exec, "run", "r", false, "Run the saved example file")

	cmd.MarkPersistentFlagRequired("lang")

	apps.cmd = cmd
	return apps
}

func (c *FindCommand) Command() *cobra.Command {
	c.cmd.Run = c.Execute
	return c.cmd
}

func (c *FindCommand) Execute(_ *cobra.Command, args []string) {
	fmt.Printf("Please implement me\n")
}
