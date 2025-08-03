package cdeps

import (
	"fmt"
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

	return nil
}

func (e *Explorer) totalDependencies(files []os.DirEntry, isDependency func(string) bool) int {
	total := 0

	for _, file := range files {
		if isDependency(file.Name()) {
			total++
		}
	}

	return total
}

func (e *Explorer) CopyDependency(dependency string, source, target string) ([]string, error) {
	files, err := os.ReadDir(source)

	if err != nil {
		return nil, fmt.Errorf("failed to copy %q from %q [%s]", dependency, source, err)
	}

	err = e.CreateDirectory(target)

	if err != nil {
		return nil, err
	}

	isDependency := regexp.MustCompile(dependency).MatchString

	var (
		filename string
		copies   = make([]string, 0, e.totalDependencies(files, isDependency))
	)

	for _, file := range files {
		filename = file.Name()

		if !isDependency(filename) {
			continue
		}

		err = e.CopyFile(filename, source, target)

		if err != nil {
			return nil, err
		}

		copies = append(copies, filepath.Join(target, filename))
	}

	return copies, nil
}
