package syncenv

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	ExtBackup = ".bak"
	RExBackup = `^.+\.bak$`
)

type Backup struct{}

func (*Backup) File(name, source, target string) error {
	data, err := os.ReadFile(filepath.Join(source, name)) //nolint:gosec

	if err != nil {
		return fmt.Errorf("failed to read %q from %q [%s]", name, source, err)
	}

	err = os.WriteFile(filepath.Join(target, name+ExtBackup), data, 0600)

	if err != nil {
		return fmt.Errorf("failed to write %q on %q [%s]", name+ExtBackup, target, err)
	}

	return nil
}

func (*Backup) Restore(backup, source string) error {
	original := strings.TrimSuffix(backup, ExtBackup)

	err := os.Rename(filepath.Join(source, backup), filepath.Join(source, original))

	if err != nil {
		return fmt.Errorf("failure to restore file %q from %q [%s]", original, backup, err)
	}

	return nil
}

func (*Backup) Remove(backup, source string) error {
	isBackup := regexp.MustCompile(RExBackup).MatchString

	if isBackup(backup) {
		err := os.Remove(filepath.Join(source, backup))

		if err != nil {
			return fmt.Errorf("failure to remove backup %q from %q [%s]", backup, source, err)
		}
	}

	return nil
}
