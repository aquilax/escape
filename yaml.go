package escape

import (
	yaml "gopkg.in/yaml.v2"
)

type YAML struct{}

func (h YAML) Name() string {
	return "yaml"
}

func (h YAML) Escape(text string) (string, error) {
	b, err := yaml.Marshal(text)
	return string(b), err
}

func (h YAML) UnEscape(text string) (string, error) {
	var s string
	err := yaml.Unmarshal([]byte(text), &s)
	return s, err
}
