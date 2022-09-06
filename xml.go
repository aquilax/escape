package escape

import (
	"bytes"
	"encoding/xml"
)

type XML struct{}

func (h XML) Name() string {
	return "xml"
}

func (h XML) Escape(text string) (string, error) {
	var w bytes.Buffer
	err := xml.EscapeText(&w, []byte(text))
	return w.String(), err
}

func (h XML) UnEscape(text string) (string, error) {
	// FIXME: implement xml unescape
	return text, nil
}
