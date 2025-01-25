package commands

type Save struct {
	Save bool `short:"s" long:"save" description:"Save example to a file" default:"false"`
}
