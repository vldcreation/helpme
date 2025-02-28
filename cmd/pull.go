package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vldcreation/helpme-package/pkg/pull"
)

type pullCmd struct {
	cmd *cobra.Command

	// flags
	host   string
	user   string
	repo   string
	branch string
}

func NewPullCommand() *pullCmd {
	apps := &pullCmd{}
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Pull depdency golang",
		Long:  "Pull depdency golang",
		Args:  cobra.ExactArgs(0),
	}

	cmd.PersistentFlags().StringVarP(&apps.host, "host", "H", "", "Hostname of the repository (e.g. github.com)")
	cmd.PersistentFlags().StringVarP(&apps.user, "user", "u", "", "Username of the repository")
	cmd.PersistentFlags().StringVarP(&apps.repo, "repo", "r", "", "Repository name")
	cmd.PersistentFlags().StringVarP(&apps.branch, "branch", "b", "", "Branch name of the repository")

	cmd.MarkPersistentFlagRequired("user")
	cmd.MarkPersistentFlagRequired("repo")

	apps.cmd = cmd
	return apps
}

func (c *pullCmd) Command() *cobra.Command {
	c.cmd.Run = c.Execute
	return c.cmd
}

func (c *pullCmd) Execute(_ *cobra.Command, args []string) {
	if err := pull.Pull(c.host, c.user, c.repo, c.branch); err != nil {
		println(err.Error())
	}

	println("Pulling sucessfully")
}
