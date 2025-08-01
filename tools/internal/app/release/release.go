package release

import (
	"flag"
	"fmt"
	"log"

	"github.com/bastean/x/tools/internal/pkg/cli"
	"github.com/bastean/x/tools/internal/pkg/errs"
	"github.com/bastean/x/tools/pkg/release"
)

var (
	name           string
	increment      string
	isFirstRelease bool
)

func Init() error {
	flag.StringVar(&name, "m", "", "Module name (required)")

	flag.StringVar(&increment, "i", "", "Increment \"patch\", \"minor\" or \"major\" (optional: if \"-f\" is used)")

	flag.BoolVar(&isFirstRelease, "f", false, "First Release (default: false)")

	flag.Usage = func() {
		cli.Usage("release")
	}

	flag.Parse()

	if flag.NFlag() < 2 {
		flag.Usage()

		println()

		return errs.ErrRequiredFlags
	}

	return nil
}

func Up() error {
	err := Init()

	if err != nil {
		return err
	}

	log.Println("Starting...")

	var module *release.Module

	switch {
	case isFirstRelease:
		module, err = release.NewModuleFirstRelease(name)
	default:
		module, err = release.NewModuleRelease(name, increment)
	}

	if err != nil {
		return err
	}

	exec := new(release.Exec)

	tag := &release.Tag{
		Doer: exec,
	}

	latest, err := tag.Latest(module)

	if err != nil {
		return err
	}

	version, err := release.BumpVersion(module, latest)

	if err != nil {
		return err
	}

	commit := &release.Commit{
		Doer: exec,
	}

	err = commit.CreateStd(module, version)

	if err != nil {
		return err
	}

	err = tag.CreateStd(module, version)

	if err != nil {
		if errReset := commit.Reset(); errReset != nil {
			return fmt.Errorf("\n\n%s\n%s", errReset, err)
		}

		return err
	}

	fmt.Printf("Successfully released \"%s %s\"\n", module.Name, version)

	log.Println("Completed!")

	return nil
}
