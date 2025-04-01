package cdeps

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	EveryMinFile   = `^.+\.min\.(js|css)$`
	EveryWoff2File = `^.+\.woff2$`
)

type Explorer struct{}

func (*Explorer) CreateDirectory(path string) error {
	err := os.MkdirAll(path, 0700)

	if err != nil {
		return fmt.Errorf("failed to create \"%s\"", path)
	}

	return nil
}

func (*Explorer) CopyFile(filename, sourcePath, targetPath string) error {
	data, err := os.ReadFile(filepath.Join(sourcePath, filepath.Base(filename))) //nolint:gosec

	if err != nil {
		return fmt.Errorf("failed to read \"%s\" from \"%s\"", filename, sourcePath)
	}

	err = os.WriteFile(filepath.Join(targetPath, filepath.Base(filename)), data, 0600)

	if err != nil {
		return fmt.Errorf("failed to write \"%s\" on \"%s\"", filename, targetPath)
	}

	return nil
}

func (e *Explorer) CopyDeps(filenames []string, sourcePath, targetPath string) error {
	files, err := os.ReadDir(sourcePath)

	if err != nil {
		return fmt.Errorf("failed to copy \"%s\" from \"%s\"", filenames, sourcePath)
	}

	err = e.CreateDirectory(targetPath)

	if err != nil {
		return err
	}

	if strings.HasPrefix(filenames[0], "^") && strings.HasSuffix(filenames[0], "$") {
		isMinFile := regexp.MustCompile(filenames[0]).MatchString

		for _, file := range files {
			if isMinFile(file.Name()) {
				err = errors.Join(err, e.CopyFile(file.Name(), sourcePath, targetPath))
			}
		}

		if err != nil {
			return err
		}

		return nil
	}

	for _, filename := range filenames {
		for _, file := range files {
			if filepath.Base(filename) == file.Name() {
				err = errors.Join(err, e.CopyFile(filename, sourcePath, targetPath))
			}
		}
	}

	if err != nil {
		return err
	}

	return nil
}
