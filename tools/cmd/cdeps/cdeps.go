package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bastean/x/tools/internal/app/cdeps"
)

const (
	cli = "cdeps"
)

var (
	configFile string
)

func usage() {
	fmt.Printf("Usage: %s [flags]\n\n", cli)
	flag.PrintDefaults()
}

func fatal(what string) {
	fmt.Printf("%s\n", what)
	os.Exit(1)
}

func main() {
	flag.StringVar(&configFile, "c", "cdeps.json", "cDeps configuration file (required)")

	flag.Usage = usage

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()

		println()

		fatal("define required flags")
	}

	log.Println("Starting...")

	if err := cdeps.Up(configFile); err != nil {
		fatal(err.Error())
	}

	log.Println("Completed!")
}
