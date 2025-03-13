package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vldcreation/helpme-package/pkg/runtest"
)

type runtestCmd struct {
	cmd *cobra.Command

	// flags
	fpath            string
	funcName         string
	mustReturnOutput bool
	inPath           string
	outPath          string
}

func NewRunTestCommand() *runtestCmd {
	apps := &runtestCmd{}
	cmd := &cobra.Command{
		Use:   "runtest",
		Short: "Run Test sample with sample output",
		Long:  "Run test sample with sample output",
		Args:  cobra.ExactArgs(0),
	}

	cmd.PersistentFlags().StringVarP(&apps.fpath, "file", "F", "", "Filepath of file to execute (e.g. mypackage/a.go)")
	cmd.PersistentFlags().StringVarP(&apps.funcName, "func", "f", "", "Function name to invoke (e.g: MyFunc)")
	cmd.PersistentFlags().BoolVarP(&apps.mustReturnOutput, "debug_out", "D", false, "Print debug output")
	cmd.PersistentFlags().StringVarP(&apps.inPath, "input", "i", "", "Input path sample")
	cmd.PersistentFlags().StringVarP(&apps.outPath, "output", "o", "", "Output path sample")

	cmd.MarkPersistentFlagRequired("file")
	cmd.MarkPersistentFlagRequired("input")
	cmd.MarkPersistentFlagRequired("output")

	apps.cmd = cmd
	return apps
}

func (c *runtestCmd) Command() *cobra.Command {
	c.cmd.Run = c.Execute
	return c.cmd
}

func (c *runtestCmd) Execute(_ *cobra.Command, args []string) {
	if res, _, err := runtest.RunTest(c.fpath, c.funcName, c.mustReturnOutput, c.inPath, c.outPath); err != nil {
		println(err.Error())
	} else {
		println(res)
	}
}
