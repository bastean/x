package cdeps

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type Explorer struct{}

func (*Explorer) CreateDirectory(path string) error {
	err := os.MkdirAll(path, 0700)

	if err != nil {
		return fmt.Errorf("failed to create %q [%s]", path, err)
	}

	return nil
}

func (*Explorer) CopyFile(file, source, target string) error {
	data, err := os.ReadFile(filepath.Join(source, filepath.Base(file))) //nolint:gosec

	if err != nil {
		return fmt.Errorf("failed to read %q from %q [%s]", file, source, err)
	}

	err = os.WriteFile(filepath.Join(target, filepath.Base(file)), data, 0600)

	if err != nil {
		return fmt.Errorf("failed to write %q on %q [%s]", file, target, err)
	}

	log.Printf("Created: %q", filepath.Join(target, file))

	return nil
}

func (e *Explorer) CopyDependency(dependency string, source, target string) error {
	files, err := os.ReadDir(source)

	if err != nil {
		return fmt.Errorf("failed to copy %q from %q [%s]", dependency, source, err)
	}

	err = e.CreateDirectory(target)

	if err != nil {
		return err
	}

	isDependency := regexp.MustCompile(dependency).MatchString

	for _, file := range files {
		if isDependency(file.Name()) {
			err = errors.Join(err, e.CopyFile(file.Name(), source, target))
		}
	}

	if err != nil {
		return err
	}

	return nil
}
