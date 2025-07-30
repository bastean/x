package main

import (
	"flag"
	"log"

	"github.com/bastean/x/tools/internal/app/syncenv"
	"github.com/bastean/x/tools/internal/pkg/cli"
	"github.com/bastean/x/tools/internal/pkg/errs"
)

var (
	template, envs string
)

func main() {
	flag.StringVar(&template, "t", "", "Path to \".env\" file template (required)")

	flag.StringVar(&envs, "e", "", "Path to \".env\" files directory (required)")

	flag.Usage = func() {
		cli.Usage("syncenv")
	}

	flag.Parse()

	if flag.NFlag() < 2 {
		flag.Usage()

		println()

		errs.Fatal("Please define required flags")
	}

	log.Println("Starting synchronization...")

	if err := syncenv.Up(template, envs); err != nil {
		errs.Fatal(err.Error())
	}

	log.Println("Synchronization completed!")
}
