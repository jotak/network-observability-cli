package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/netobserv/network-observability-cli/cmd"
)

var (
	buildVersion = "unknown"
	buildDate    = "unknown"
)

func main() {
	// Initial log message
	fmt.Printf("Starting %s:\n=====\nBuild version: %s\nBuild date: %s\n\n",
		filepath.Base(os.Args[0]), buildVersion, buildDate)

	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
