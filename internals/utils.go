package internals

import (
	"bufio"
	"fmt"
	"os"
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

func CheckX11orWayland() string {
	session := os.Getenv("XDG_SESSION_TYPE")
	wayland := os.Getenv("WAYLAND_DISPLAY")

	if session == "wayland" && wayland != "" {
		return "wayland"
	} else {
		return "x11"
	}
}
