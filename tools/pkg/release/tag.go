package release

import (
	"fmt"
	"strings"
)

type Tag struct {
	Doer
}

func (t *Tag) Latest(module *Module) (string, error) {
	if module.IsFirstRelease {
		return module.Name + "/v0.0.0", nil
	}

	output, err := t.Do("bash", "-c", fmt.Sprintf("git tag --sort -v:refname | grep %s | head -n 1", module.Name))

	switch {
	case err != nil:
		return "", err
	case len(output) == 0:
		return "", fmt.Errorf("no previous release found for %q", module.Name)
	}

	return strings.TrimRight(output, "\n"), nil
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
