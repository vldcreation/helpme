package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vldcreation/helpme-package/pkg/fileserver"
	"github.com/vldcreation/helpme/util"
)

type shareFileCmd struct {
	cmd *cobra.Command

	// flags
	rootDir string
	port    string
}

func NewShareFileCommand() *shareFileCmd {
	apps := &shareFileCmd{}
	cmd := &cobra.Command{
		Use:   "sharefile",
		Short: "Share workspace directory with same network",
		Long:  "Share workspace directory with same network",
		Args:  cobra.ExactArgs(0),
	}

	cmd.PersistentFlags().StringVarP(&apps.rootDir, "dir", "D", "", "Root directory of workspace")
	cmd.PersistentFlags().StringVarP(&apps.port, "port", "P", "9000", "Port of server")
	cmd.MarkPersistentFlagRequired("root")

	cmd.Run = apps.Execute

	apps.cmd = cmd
	return apps
}

func (c *shareFileCmd) Command() *cobra.Command {
	return c.cmd
}

func (c *shareFileCmd) Execute(_ *cobra.Command, args []string) {
	if err := fileserver.New(c.rootDir, util.GetLocalIP(), fileserver.WithPort(c.port)).Run(); err != nil {
		util.PrintlnRed(err.Error())
	}
}
