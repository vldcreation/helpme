package pkg

// Language
type Language struct {
	lang     string
	pkg      string
	funcName string
	dir      string

	// opt
	execute bool
	save    bool
}

func NewLanguage(lang, pkg, funcName string) Language {
	return Language{
		lang:     lang,
		pkg:      pkg,
		funcName: funcName,
	}
}

// apply flag
func (l *Language) Apply(opt ...LangOpt) {
	for _, o := range opt {
		o(l)
	}
}

type LangOpt func(*Language)

func WithExecute() LangOpt {
	return func(f *Language) {
		f.execute = true
	}
}

func WithSave(dir string) LangOpt {
	return func(f *Language) {
		f.save = true
		f.dir = dir
	}
}

// global config
var (
	docBaseUrl = map[string]string{
		"go":         "https://pkg.go.dev/",
		"javascript": "https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/",
	}

	exampleCodeBaseUrl = map[string]string{
		"go": "https://pkg.go.dev/%s@go1.23.5#example-%s",
	}
)
