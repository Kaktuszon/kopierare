package internals

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func LoadLines(path string) ([]string, error) {
	lines := []string{}

	if path == "" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(homedir, ".kopierare", "config.txt")
	}
	data, err := os.Open(path)

	if err != nil {
		fmt.Println("File not found")
		return nil, err
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
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
