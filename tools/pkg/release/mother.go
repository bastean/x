package release

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) ModuleReleaseValid() *Module {
	module, err := NewModuleRelease(
		m.LoremIpsumWord(),
		m.RandomString([]string{"patch", "minor", "major"}),
	)

	if err != nil {
		panic(err.Error())
	}

	return module
}

func (m *m) ModuleFirstReleaseValid() *Module {
	module, err := NewModuleFirstRelease(
		m.LoremIpsumWord(),
	)

	if err != nil {
		panic(err.Error())
	}

	return module
}

func (m *m) ModuleInvalidIncrement() (*Module, string) {
	value := "x"

	module := m.ModuleReleaseValid()

	module.Increment = value

	return module, value
}

func (m *m) TagValid(module *Module) (latest string, version string, major int, minor int, patch int) {
	version = m.AppVersion()

	semver := strings.Split(version, ".")

	major, errMajor := strconv.Atoi(semver[0])
	minor, errMinor := strconv.Atoi(semver[1])
	patch, errPatch := strconv.Atoi(semver[2])

	if err := errors.Join(errMajor, errMinor, errPatch); err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/v%s", module.Name, version), version, major, minor, patch
}

var Mother = mother.New[m]
