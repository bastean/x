package cdeps

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/bastean/x/tools/pkg/cdeps"
)

type Dependency struct {
	Files          []string
	Source, Target string
}

type Configuration struct {
	Wildcards    map[string]string
	Dependencies []*Dependency
}

func Up(configFile string) error {
	config := new(Configuration)

	data, err := os.ReadFile(configFile) //nolint:gosec

	if err != nil {
		return fmt.Errorf("failure to read configuration file %q [%s]", configFile, err.Error())
	}

	err = json.Unmarshal(data, config)

	if err != nil {
		return fmt.Errorf("failure to decode configuration file %q [%s]", configFile, err.Error())
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
			if err = explorer.CopyDependency(file, dependency.Source, dependency.Target); err != nil {
				return err
			}
		}
	}

	return nil
}
