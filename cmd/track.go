package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vldcreation/helpme-package/pkg/configurator"
	"github.com/vldcreation/helpme-package/pkg/trackclipboard"
)

type trackCmd struct {
	cmd *cobra.Command

	// flags
	cfgPath string
}

func NewTrackCommand() *trackCmd {
	apps := &trackCmd{}
	cmd := &cobra.Command{
		Use:   "trackclipboard",
		Short: "Track data from clipboard and send to your channel",
		Long:  "Track data from clipboard and send to your channel",
		Args:  cobra.ExactArgs(0),
	}

	cmd.PersistentFlags().StringVarP(&apps.cfgPath, "config", "C", "", "Config Filepath to use (e.g. myconfig/config.yaml)")

	cmd.MarkPersistentFlagRequired("config")

	apps.cmd = cmd
	return apps
}

func (c *trackCmd) Command() *cobra.Command {
	c.cmd.Run = c.Execute
	return c.cmd
}

func (c *trackCmd) Execute(_ *cobra.Command, args []string) {
	config := trackclipboard.Config{}
	if err := configurator.LoadFromYaml(c.cfgPath, &config); err != nil {
		panic(err)
	}
	trackclipboard.NewTrackClipboard(&config).Track()
}
