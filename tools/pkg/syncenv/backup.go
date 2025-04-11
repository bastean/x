package syncenv

import (
	"fmt"
	"os"
)

const (
	ExtBackup = ".syncenv.bak"
)

type Backup struct{}

func (*Backup) Create(file string) error {
	data, err := os.ReadFile(file) //nolint:gosec

	if err != nil {
		return fmt.Errorf("failed to read %q [%s]", file, err)
	}

	file += ExtBackup

	err = os.WriteFile(file, data, 0600)

	if err != nil {
		return fmt.Errorf("failed to write %q [%s]", file, err)
	}

	return nil
}

func (*Backup) Restore(file string) error {
	err := os.Rename(file+ExtBackup, file)

	if err != nil {
		return fmt.Errorf("failure to restore file %q [%s]", file, err)
	}

	return nil
}

func (*Backup) Remove(file string) error {
	file += ExtBackup

	err := os.Remove(file)

	if err != nil {
		return fmt.Errorf("failure to remove backup %q [%s]", file, err)
	}

	return nil
}
