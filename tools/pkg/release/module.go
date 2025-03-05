package release

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

var (
	Increments = []string{
		"patch",
		"minor",
		"major",
	}
)

type Module struct {
	Name, Increment string
	IsFirstRelease  bool
}

func newModule(name, increment string, isFirstRelease bool) (*Module, error) {
	var err error

	switch {
	case !slices.Contains(Increments, increment):
		err = fmt.Errorf("%q is not a valid increment value, allowed values are \"patch\", \"minor\" or \"major\"", increment)
	case strings.TrimSpace(name) == "":
		err = errors.New("module name is required")
	}

	if err != nil {
		return nil, err
	}

	return &Module{
		Name:           name,
		Increment:      increment,
		IsFirstRelease: isFirstRelease,
	}, nil
}

func NewModuleFirstRelease(name string) (*Module, error) {
	return newModule(name, "minor", true)
}

func NewModuleRelease(name, increment string) (*Module, error) {
	return newModule(name, increment, false)
}
