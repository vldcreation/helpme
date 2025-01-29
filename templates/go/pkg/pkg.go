package pkg

import _ "embed"

var (
	// go:embed pkg.tmpl
	DefaultPackage string
)
