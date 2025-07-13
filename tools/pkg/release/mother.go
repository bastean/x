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

func (m *m) RandomModuleRelease() *Module {
	module, err := NewModuleRelease(
		m.LoremIpsumWord(),
		m.RandomString([]string{"patch", "minor", "major"}),
	)

	if err != nil {
		panic(err.Error())
	}

	return module
}

func (m *m) RandomModuleFirstRelease() *Module {
	module, err := NewModuleFirstRelease(
		m.LoremIpsumWord(),
	)

	if err != nil {
		panic(err.Error())
	}

	return module
}

func (m *m) ModuleWithInvalidIncrement() (*Module, string) {
	value := "x"

	module := m.RandomModuleRelease()

	module.Increment = value

	return module, value
}

func (m *m) RandomTag(module *Module) (string, string, int, int, int) {
	version := m.AppVersion()

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
