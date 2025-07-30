package main

import (
	"flag"
	"fmt"

	"github.com/bastean/x/tools/internal/pkg/cli"
	"github.com/bastean/x/tools/internal/pkg/errs"
	"github.com/bastean/x/tools/pkg/release"
)

var (
	err            error
	name           string
	increment      string
	isFirstRelease bool
)

func main() {
	flag.StringVar(&name, "m", "", "Module name (required)")

	flag.StringVar(&increment, "i", "", "Increment \"patch\", \"minor\" or \"major\" (optional: if \"-f\" is used)")

	flag.BoolVar(&isFirstRelease, "f", false, "First Release (default: false)")

	flag.Usage = func() {
		cli.Usage("release")
	}

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()

		println()

		errs.Fatal("define required flags")
	}

	var module *release.Module

	switch {
	case isFirstRelease:
		module, err = release.NewModuleFirstRelease(name)
	default:
		module, err = release.NewModuleRelease(name, increment)
	}

	if err != nil {
		errs.Fatal(err.Error())
	}

	exec := new(release.Exec)

	tag := &release.Tag{
		Doer: exec,
	}

	latest, err := tag.Latest(module)

	if err != nil {
		errs.Fatal(err.Error())
	}

	version, err := release.BumpVersion(module, latest)

	if err != nil {
		errs.Fatal(err.Error())
	}

	commit := &release.Commit{
		Doer: exec,
	}

	err = commit.CreateStd(module, version)

	if err != nil {
		errs.Fatal(err.Error())
	}

	err = tag.CreateStd(module, version)

	if err != nil {
		if errReset := commit.Reset(); errReset != nil {
			errs.Fatal(fmt.Sprintf("\n\n%s\n%s", errReset.Error(), err.Error()))
		}

		errs.Fatal(err.Error())
	}

	fmt.Printf("Successfully released \"%s %s\"\n", module.Name, version)
}
