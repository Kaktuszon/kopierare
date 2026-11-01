package main

import (
	"flag"
	"fmt"
	"kopierare/v2/utils"
	"os"
)

func main() {
	configPath := flag.String("path", "", "Path to config file, if empty ~/.kopierare/config.txt")
	flag.Parse()

	var cfg utils.Config
	cfg, err := utils.LoadConfig(*configPath)

	var input int

	lines := cfg.Lines
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, v := range lines {
		fmt.Printf("[%v] %s\n", i, v)
	}

	fmt.Printf("\nCopy number: ")
	fmt.Scan(&input)

	utils.HandleCopyCommand(lines[input])

	fmt.Printf("Copied %s to clipboard\n", lines[input])
}
