package cmd

import (
	password_generator "github.com/vldcreation/helpme/pkg/generator/password"

	"github.com/spf13/cobra"
)

type generatePasswordCmd struct {
	cmd *cobra.Command

	// flags
	qty      int
	len      int
	passType int
}

func NewGeneratePasswordCommand() *generatePasswordCmd {
	apps := &generatePasswordCmd{}
	cmd := &cobra.Command{
		Use:   "generate-password",
		Short: "Generate a password",
		Long:  "Generate a password",
		Args:  cobra.ExactArgs(0),
	}

	cmd.PersistentFlags().IntVarP(&apps.qty, "qty", "q", 1, "Quantity of passwords to generate")
	cmd.PersistentFlags().IntVarP(&apps.len, "len", "l", 0, "Password length (words or chars)")
	cmd.PersistentFlags().IntVarP(&apps.passType, "type", "t", 0, "Password type (0: word, 1: phrase, 2: word with special, 3: phrase with special, 4: secure)")

	cmd.Run = apps.Execute
	apps.cmd = cmd
	return apps
}

func (c *generatePasswordCmd) Command() *cobra.Command {
	return c.cmd
}

func (c *generatePasswordCmd) Execute(_ *cobra.Command, args []string) {
	for i := 0; i < c.qty; i++ {
		if out, hint, err := password_generator.GeneratePassword(c.len, c.passType); err != nil {
			println(err.Error())
		} else {
			println("Password: " + out)
			if hint != "" {
				println("Hint: " + hint)
			}
		}
		println()
	}
}
