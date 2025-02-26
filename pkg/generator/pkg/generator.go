package pkg

type generator interface {
	Generate() error
}

func NewGenerator(l Language, opt ...FlagOpt) generator {
	if l.lang == "" {
		return nil
	}

	for _, o := range opt {
		o(&l.f)
	}

	switch l.lang {
	case "go":
		return &goGenerator{
			l: l,
		}
	case "javascript":
		return &javascriptGenerator{
			l: l,
		}
	default:
		return nil
	}
}
