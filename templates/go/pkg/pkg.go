package pkg

import _ "embed"

var (
	//go:embed default.tmpl
	DefaultPackage string
)
