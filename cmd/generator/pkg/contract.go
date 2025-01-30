package pkg

// Language
type Language struct {
	lang     string
	pkg      string
	funcName string

	// flag
	f flag
}

func NewLanguage(lang, pkg, funcName string) Language {
	return Language{
		lang:     lang,
		pkg:      pkg,
		funcName: funcName,
	}
}

// apply flag
func (l *Language) Apply(opt ...FlagOpt) {
	for _, o := range opt {
		o(&l.f)
	}
}

// flag
type flag struct {
	execute bool
	save    bool
}

type FlagOpt func(*flag)

func WithExecute() FlagOpt {
	return func(f *flag) {
		f.execute = true
	}
}

func WithSave() FlagOpt {
	return func(f *flag) {
		f.save = true
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
