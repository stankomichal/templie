package config

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

type FileDecorator struct {
	Icon  string       `yaml:"icon"`
	Hex   string       `yaml:"hex"`
	Color *color.Color `yaml:"-"`
}

func (fd *FileDecorator) UnmarshalYAML(value *yaml.Node) error {
	type rawDecorator FileDecorator
	var raw rawDecorator
	if err := value.Decode(&raw); err != nil {
		return err
	}

	r, g, b, err := hexToRgb(raw.Hex)
	if err != nil {
		return err
	}
	raw.Color = color.RGB(r, g, b)
	*fd = FileDecorator(raw)
	return nil
}

func hexToRgb(hex string) (int, int, int, error) {
	var r, g, b int
	_, err := fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	if err != nil {
		return 0, 0, 0, err
	}
	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		return 0, 0, 0, fmt.Errorf("invalid RGB value")
	}
	return r, g, b, nil
}
