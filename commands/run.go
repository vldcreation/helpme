package commands

type Run struct {
	Run bool `short:"r" long:"run" description:"Run the saved example file" default:"false"`
}
