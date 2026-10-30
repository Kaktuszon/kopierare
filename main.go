package main

import (
	"fmt"
	"kopierare/v2/internals"
	"os"
	"os/exec"
)

func main() {
	lines := []string{}
	lines, err := internals.LoadLines("")

	var input int

	if err != nil {
		fmt.Println(err)
	}

	for i, v := range lines {
		fmt.Printf("[%v] %s\n", i, v)
	}

	fmt.Printf("\nCopy number: ")
	fmt.Scan(&input)

	// Select which command to use for copying
	cmd := exec.Command("")
	if internals.CheckX11orWayland() == "wayland" {
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
		_, _ = stdin.Write([]byte(lines[input]))
	}()

	if err := cmd.Start(); err != nil {
		fmt.Println("Error running copy command. Is xclip installed?")
		os.Exit(1)
	}

	fmt.Printf("Copied %s to clipboard\n", lines[input])
}
