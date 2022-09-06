package escape

import (
	"encoding/json"
)

type JSON struct{}

func (h JSON) Name() string {
	return "json"
}

func (h JSON) Escape(text string) (string, error) {
	b, err := json.Marshal(text)
	return string(b), err
}

func (h JSON) UnEscape(text string) (string, error) {
	var s string
	err := json.Unmarshal([]byte(text), &s)
	return s, err
}
