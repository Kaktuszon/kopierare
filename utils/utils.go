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

func LoadLines(path string) ([]string, error) {
	if path == "" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(homedir, ".kopierare", "config.yaml")
	}
	file, _ := os.ReadFile(path)

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}

	return cfg.Lines, nil
}

func checkX11orWayland() string {
	session := os.Getenv("XDG_SESSION_TYPE")
	wayland := os.Getenv("WAYLAND_DISPLAY")

	if session == "wayland" && wayland != "" {
		return "wayland"
	} else {
		return "x11"
	}
}

func HandleCopyCommand(input string) {
	cmd := exec.Command("")
	if checkX11orWayland() == "wayland" {
		cmd = exec.Command("wl-copy")
	} else {
		cmd = exec.Command("xclip", "-selection", "clipboard")
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Failed creating the pipe")
		os.Exit(1)
	}

	// Write to clipboard
	go func() {
		defer stdin.Close()
		_, _ = stdin.Write([]byte(input))
	}()

	if err := cmd.Start(); err != nil {
		fmt.Println("Error running copy command. Is xclip installed?")
		os.Exit(1)
	}
}
