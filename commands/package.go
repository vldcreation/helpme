package commands

type Package struct {
	Package string `short:"p" long:"pkg" description:"Package name (optional)" default:""`
}
