package cmd

import (
	pkg_generator "vldcreation/github.com/helpme/pkg/generator/pkg"

	"github.com/spf13/cobra"
)

type findCommand struct {
	cmd *cobra.Command

	// flags
	lang string
	pkg  string
	save bool
	exec bool
}

func NewFindCommand() *findCommand {
	apps := &findCommand{}
	cmd := &cobra.Command{
		Use:   "find",
		Short: "Find an example for a given function",
		Long:  "Find an example for a given function in a given language",
		Args:  cobra.ExactArgs(1),
	}

	cmd.PersistentFlags().StringVarP(&apps.lang, "lang", "l", "", "Language to search (go/javascript)")
	cmd.PersistentFlags().StringVarP(&apps.pkg, "pkg", "p", "", "Package name (optional)")
	cmd.PersistentFlags().BoolVarP(&apps.save, "save", "s", false, "Save example to a file")
	cmd.PersistentFlags().BoolVarP(&apps.exec, "exec", "e", false, "Run the saved example file")

	cmd.MarkPersistentFlagRequired("lang")

	apps.cmd = cmd
	return apps
}

func (c *findCommand) Command() *cobra.Command {
	c.cmd.Run = c.Execute
	return c.cmd
}

func (c *findCommand) Execute(_ *cobra.Command, args []string) {
	funcName := args[0]

	l := pkg_generator.NewLanguage(c.lang, c.pkg, funcName)
	if err := pkg_generator.NewGenerator(l, c.generateFlagOpts()...).Generate(); err != nil {
		panic(err)
	}
}

func (c *findCommand) generateFlagOpts() (opts []pkg_generator.FlagOpt) {
	if c.save {
		opts = append(opts, pkg_generator.WithSave())
	}
	if c.exec {
		opts = append(opts, pkg_generator.WithExecute())
	}

	return
}
