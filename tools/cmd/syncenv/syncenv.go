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
	path     string
	template string
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
	flag.StringVar(&path, "p", ".", "Path to \".env\" files (required)")

	flag.StringVar(&template, "t", ".env.example", ".env file template (required)")

	flag.Usage = usage

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()

		println()

		fatal("define required flags")
	}

	log.Println("Starting...")

	if err := syncenv.Up(path, template); err != nil {
		fatal(err.Error())
	}

	log.Println("Completed!")
}
