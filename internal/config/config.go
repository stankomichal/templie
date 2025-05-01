package config

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	TemplateFolder  string                   `yaml:"template_folder"`
	DefaultCategory string                   `yaml:"default_category"`
	FileDecorators  map[string]FileDecorator `yaml:"file_decorators"`
}

func DefaultConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not get user home dir: %w", err)
	}

	return &Config{
		TemplateFolder:  filepath.Join(homeDir, ".config", "templie", "templates"),
		DefaultCategory: "general",
		FileDecorators: map[string]FileDecorator{
			".js":   {Icon: "📜", Hex: "#007acc", Color: color.New(color.FgCyan)},
			".ts":   {Icon: "📘", Hex: "#007acc", Color: color.New(color.FgCyan)},
			".json": {Icon: "🧾", Hex: "#ff7b72", Color: color.New(color.FgMagenta)},
			".go":   {Icon: "🐹", Hex: "#00add8", Color: color.New(color.FgBlue)},
			".yml":  {Icon: "⚙️", Hex: "#ffffff", Color: color.New(color.FgWhite)},
			".yaml": {Icon: "⚙️", Hex: "#ffffff", Color: color.New(color.FgWhite)},
			".cpp":  {Icon: "💻", Hex: "#ff7b72", Color: color.New(color.FgMagenta)},
		},
	}, nil
}

func (c *Config) Show(varName string) (string, error) {
	varPointer := c.getVar(varName)
	if varPointer == nil {
		return "", fmt.Errorf("could not find the variable %s", varName)
	}
	return *(varPointer.(*string)), nil
}

func (c *Config) Update(varName string, varNewValue string) (string, error) {
	varPointer := c.getVar(varName)
	if varPointer == nil {
		return "", fmt.Errorf("could not find the variable %s", varName)
	}
	*(varPointer.(*string)) = varNewValue

	if err := writeConfig(c); err != nil {
		return "", fmt.Errorf("could not write config file: %w", err)
	}

	return varNewValue, nil
}

func (c *Config) Reset(varName string) (string, error) {
	varPointer := c.getVar(varName)
	if varPointer == nil {
		return "", fmt.Errorf("could not find the variable %s", varName)
	}

	defaultConfig, err := DefaultConfig()
	if err != nil {
		return "", fmt.Errorf("could not get default config: %w", err)
	}
	varPointerDefault := defaultConfig.getVar(varName)
	*(varPointer.(*string)) = *(varPointerDefault.(*string))
	if err := writeConfig(c); err != nil {
		return "", fmt.Errorf("could not write config file: %w", err)
	}

	return *(varPointer.(*string)), nil
}

func Load(cmd *cobra.Command) (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not get user home dir: %w", err)
	}
	configDir := filepath.Join(homeDir, ".config", "templie")
	configPath := filepath.Join(configDir, "settings.yaml")

	// Check if the config file exists
	if _, err = os.Stat(configPath); os.IsNotExist(err) {
		// Folder
		if err = os.MkdirAll(configDir, 0755); err != nil {
			return nil, fmt.Errorf("could not create config directory: %w", err)
		}

		// Config file
		defaultConfig, err := DefaultConfig()
		if err != nil {
			return nil, fmt.Errorf("could not get default config: %w", err)
		}

		if err = writeConfig(defaultConfig); err != nil {
			return nil, err
		}

		cmd.Printf("Created default config at %s\n", configPath)
		return defaultConfig, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("could not read config: %w", err)
	}

	var cfg Config
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("could not parse YAML: %w", err)
	}

	return &cfg, nil
}

func (c *Config) getVar(varName string) interface{} {
	switch varName {
	case "templateFolder":
		return &c.TemplateFolder
	case "defaultCategory":
		return &c.DefaultCategory
	default:
		return nil
	}
}

func writeConfig(config *Config) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not get user home dir: %w", err)
	}
	configPath := filepath.Join(homeDir, ".config", "templie", "settings.yaml")

	out, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("could not marshal default config: %w", err)
	}

	if err = os.WriteFile(configPath, out, 0644); err != nil {
		return fmt.Errorf("could not write config file: %w", err)
	}
	return nil
}
