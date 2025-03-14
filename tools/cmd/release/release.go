package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bastean/x/tools/pkg/release"
)

const (
	cli = "release"
)

var (
	err            error
	name           string
	increment      string
	isFirstRelease bool
)

func usage() {
	fmt.Printf("Usage: %s [flags]\n\n", cli)
	flag.PrintDefaults()
}

func fatal(what string) {
	fmt.Printf("Error: %s\n", what)
	os.Exit(1)
}

func main() {
	flag.StringVar(&name, "m", "", "Module name (required)")

	flag.StringVar(&increment, "i", "", "Increment \"patch\", \"minor\" or \"major\" (optional: if \"-f\" is used)")

	flag.BoolVar(&isFirstRelease, "f", false, "First Release (default: false)")

	flag.Usage = usage

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()

		println()

		fatal("define required flags")
	}

	var module *release.Module

	switch {
	case isFirstRelease:
		module, err = release.NewModuleFirstRelease(name)
	default:
		module, err = release.NewModuleRelease(name, increment)
	}

	if err != nil {
		fatal(err.Error())
	}

	exec := new(release.Exec)

	tag := &release.Tag{
		Doer: exec,
	}

	latest, err := tag.Latest(module)

	if err != nil {
		fatal(err.Error())
	}

	version, err := release.BumpVersion(module, latest)

	if err != nil {
		fatal(err.Error())
	}

	commit := &release.Commit{
		Doer: exec,
	}

	err = commit.CreateStd(module, version)

	if err != nil {
		fatal(err.Error())
	}

	err = tag.CreateStd(module, version)

	if err != nil {
		if errReset := commit.Reset(); errReset != nil {
			fatal(fmt.Sprintf("\n\n%s\n%s", errReset.Error(), err.Error()))
		}

		fatal(err.Error())
	}

	fmt.Printf("Successfully released \"%s %s\"\n", module.Name, version)
}
