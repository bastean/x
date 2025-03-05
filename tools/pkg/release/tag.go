package release

import (
	"errors"
	"fmt"
	"strings"
)

type Tag struct {
	Doer
}

func (t *Tag) Latest(module *Module) (string, error) {
	if module.IsFirstRelease {
		return fmt.Sprintf("%s/v0.0.0", module.Name), nil
	}

	output, err := t.Do("bash", "-c", fmt.Sprintf("git tag --sort -v:refname | grep %s | head -n 1", module.Name))

	switch {
	case err != nil:
		return "", errors.New(string(output))
	case len(output) == 0:
		return "", fmt.Errorf("no previous release found for %q", module.Name)
	}

	return strings.TrimRight(string(output), "\n"), nil
}

func (t *Tag) Create(annotate, message string) error {
	if _, err := t.Do("git", "tag", "-a", annotate, "-m", message); err != nil {
		return err
	}

	return nil
}

func (t *Tag) CreateStd(module *Module, version string) error {
	return t.Create(fmt.Sprintf("%s/v%s", module.Name, version), fmt.Sprintf("%s %s", module.Name, version))
}
