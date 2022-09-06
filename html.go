package escape

import "html"

type HTML struct{}

func (h HTML) Escape(text string) (string, error) {
	return html.EscapeString(text), nil
}

func (h HTML) UnEscape(text string) (string, error) {
	return html.UnescapeString(text), nil
}
