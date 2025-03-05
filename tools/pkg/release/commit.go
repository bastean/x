package release

import (
	"fmt"
)

type Commit struct {
	Doer
}

func (c *Commit) Create(message string) error {
	if _, err := c.Do("git", "commit", "--allow-empty", "-m", message); err != nil {
		return err
	}

	return nil
}

func (c *Commit) CreateStd(module *Module, version string) error {
	return c.Create(fmt.Sprintf("chore(release): %s/v%s", module.Name, version))
}

func (c *Commit) Reset() error {
	if _, err := c.Do("git", "reset", "--hard", "HEAD^"); err != nil {
		return err
	}

	return nil
}
