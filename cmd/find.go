package cmd

import (
	pkg_generator "github.com/vldcreation/helpme/pkg/generator/pkg"

	"github.com/spf13/cobra"
)

type findCommand struct {
	cmd *cobra.Command

	// flags
	lang string
	pkg  string
	save bool
	exec bool
	dir  string
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
	cmd.PersistentFlags().StringVarP(&apps.dir, "dir", "d", ".", "Directory to save the example file")

	cmd.MarkPersistentFlagRequired("lang")

	cmd.Run = apps.Execute

	apps.cmd = cmd
	return apps
}

func (c *findCommand) Command() *cobra.Command {
	return c.cmd
}

func (c *findCommand) Execute(_ *cobra.Command, args []string) {
	funcName := args[0]

	l := pkg_generator.NewLanguage(c.lang, c.pkg, funcName)
	if err := pkg_generator.NewGenerator(l, c.generateFlagOpts()...).Generate(); err != nil {
		panic(err)
	}
}

func (c *findCommand) generateFlagOpts() (opts []pkg_generator.LangOpt) {
	if c.save {
		opts = append(opts, pkg_generator.WithSave(c.dir))
	}
	if c.exec {
		opts = append(opts, pkg_generator.WithExecute())
	}

	return
}
