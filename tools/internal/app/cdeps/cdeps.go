package cdeps

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/bastean/x/tools/internal/pkg/cli"
	"github.com/bastean/x/tools/internal/pkg/errs"
	"github.com/bastean/x/tools/internal/pkg/log"
	"github.com/bastean/x/tools/pkg/cdeps"
)

const (
	App = "cDeps"
)

var (
	configFile string
)

type Dependency struct {
	Files          []string
	Source, Target string
}

type Configuration struct {
	Wildcards    map[string]string
	Dependencies []*Dependency
}

func Init() error {
	flag.StringVar(&configFile, "c", "cdeps.json", "cDeps configuration file (required)")

	flag.Usage = func() {
		cli.Usage(App)
	}

	flag.Parse()

	if flag.NFlag() < 1 {
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

	log.Logo(App)

	log.Starting()

	config := new(Configuration)

	data, err := os.ReadFile(configFile) //nolint:gosec

	if err != nil {
		return fmt.Errorf("failure to read configuration file %q [%s]", configFile, err)
	}

	err = json.Unmarshal(data, config)

	if err != nil {
		return fmt.Errorf("failure to decode configuration file %q [%s]", configFile, err)
	}

	var errInterpolation error

	for key, value := range config.Wildcards {
		if HasWildcard(value) {
			config.Wildcards[key], errInterpolation = Interpolate(value, config.Wildcards)
			err = errors.Join(err, errInterpolation)
		}
	}

	if err != nil {
		return err
	}

	explorer := new(cdeps.Explorer)

	var (
		list, copies []string
	)

	for _, dependency := range config.Dependencies {
		for i, file := range dependency.Files {
			if HasWildcard(file) {
				dependency.Files[i], errInterpolation = Interpolate(file, config.Wildcards)
				err = errors.Join(err, errInterpolation)
			}
		}

		if HasWildcard(dependency.Source) {
			dependency.Source, errInterpolation = Interpolate(dependency.Source, config.Wildcards)
			err = errors.Join(err, errInterpolation)
		}

		if HasWildcard(dependency.Target) {
			dependency.Target, errInterpolation = Interpolate(dependency.Target, config.Wildcards)
			err = errors.Join(err, errInterpolation)
		}

		if err != nil {
			return err
		}

		for _, file := range dependency.Files {
			list, err = explorer.CopyDependency(file, dependency.Source, dependency.Target)

			if err != nil {
				return err
			}

			copies = append(copies, list...)
		}
	}

	log.Created(copies...)

	log.Completed()

	return nil
}
