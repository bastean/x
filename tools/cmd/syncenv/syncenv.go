package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bastean/x/tools/internal/app/syncenv"
)

const (
	cli = "syncenv"
)

var (
	template, envs string
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
	flag.StringVar(&template, "t", "", "Path to \".env\" file template (required)")

	flag.StringVar(&envs, "e", "", "Path to \".env\" files directory (required)")

	flag.Usage = usage

	flag.Parse()

	if flag.NFlag() < 2 {
		flag.Usage()

		println()

		fatal("Please define required flags")
	}

	log.Println("Starting synchronization...")

	if err := syncenv.Up(template, envs); err != nil {
		fatal(err.Error())
	}

	log.Println("Synchronization completed!")
}
