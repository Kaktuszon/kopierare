package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	Software string   `yaml:"software"`
	Lines    []string `yaml:"lines"`
}

func LoadConfig(path string) (Config, error) {
	var cfg Config
	if path == "" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return cfg, err
		}
		path = filepath.Join(homedir, ".kopierare", "config.yaml")
	}
	file, _ := os.ReadFile(path)

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func checkCopySoftware(cfg Config) (*exec.Cmd, error) {
	cmd := &exec.Cmd{}
	switch cfg.Software {
	case "xclip":
		cmd = exec.Command("xclip", "-selection", "clipboard")
	case "wl-copy":
		cmd = exec.Command("wl-copy")
	default:
		return nil, fmt.Errorf("Unrecognised copying Software. Valid: xclip, wl-copy\n")
	}

	return cmd, nil
}

func HandleCopyCommand(input string, cfg Config) error {
	cmd, err := checkCopySoftware(cfg)
	if err != nil {
		return err
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	// Write to clipboard
	go func() {
		defer stdin.Close()
		_, _ = stdin.Write([]byte(input))
	}()

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Error running copy command.\n")
	}

	return nil
}
