package escape

type Escaper interface {
	Name() string
	Escape(text string) (string, error)
	UnEscape(text string) (string, error)
}

func Escape(text string, escaper Escaper) (string, error) {
	return escaper.Escape(text)
}

func UnEscape(text string, escaper Escaper) (string, error) {
	return escaper.UnEscape(text)
}

func AllEscapers() []Escaper {
	return []Escaper{
		HTML{},
		YAML{},
		JSON{},
	}
}
