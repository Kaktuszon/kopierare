package main

import (
	"flag"
	"fmt"
	"kopierare/v2/utils"
	"os"
)

func main() {
	configPath := flag.String("path", "", "Path to config file, if empty ~/.kopierare/config.yaml")
	flag.Parse()

	var cfg utils.Config
	cfg, err := utils.LoadConfig(*configPath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var input int

	lines := cfg.Lines
	for i, v := range lines {
		fmt.Printf("[%v] %s\n", i, v)
	}

	fmt.Printf("\nCopy number: ")
	fmt.Scan(&input)

	if err = utils.HandleCopyCommand(lines[input], cfg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Copied %s to clipboard\n", lines[input])
}
