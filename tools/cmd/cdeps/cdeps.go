package main

import (
	"flag"
	"log"

	"github.com/bastean/x/tools/internal/app/cdeps"
	"github.com/bastean/x/tools/internal/pkg/cli"
	"github.com/bastean/x/tools/internal/pkg/errs"
)

var (
	configFile string
)

func main() {
	flag.StringVar(&configFile, "c", "cdeps.json", "cDeps configuration file (required)")

	flag.Usage = func() {
		cli.Usage("cdeps")
	}

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()

		println()

		errs.Fatal("define required flags")
	}

	log.Println("Starting...")

	if err := cdeps.Up(configFile); err != nil {
		errs.Fatal(err.Error())
	}

	log.Println("Completed!")
}
