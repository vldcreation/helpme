package commands

import "github.com/spf13/cobra"

type Lang struct {
	Lang string `short:"l" long:"lang" description:"Language" choice:"go" choice:"python" choice:"javascript" default:"go"`
}

func (l *Lang) GetLang() string {
	return l.Lang
}

func NewLang() *cobra.Command {
	return &cobra.Command{
		GroupID: "helpme",
	}
}
