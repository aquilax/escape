package escape

import "testing"

var escapers = []Escaper{
	HTML{},
	// XML{},
	YAML{},
	JSON{},
}

func TestEscape(t *testing.T) {
	want := "potato"
	for _, escaper := range escapers {
		t.Run(escaper.Name(), func(t *testing.T) {
			escaped, err := Escape("potato", escaper)
			if err != nil {
				t.Errorf("Escape() error = %v", err)
				return
			}
			got, err := UnEscape(escaped, escaper)
			if err != nil {
				t.Errorf("Escape() error = %v", err)
				return
			}
			if got != want {
				t.Errorf("Escape() = %v, want %v", got, want)
			}
		})
	}
}
