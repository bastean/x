package syncenv

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bastean/x/tools/pkg/syncenv"
)

const (
	EnvFile = ".env"
)

func Up(templatePath, envsDir string) error {
	env := new(syncenv.Env)

	templateEnvs, err := env.Dump(templatePath)

	if err != nil {
		return err
	}

	log.Printf("Template: %q", templatePath)

	files, err := os.ReadDir(envsDir)

	if err != nil {
		return fmt.Errorf("failure to read files from %q [%s]", envsDir, err)
	}

	backup := new(syncenv.Backup)

	var filename, filePath string

	for _, file := range files {
		filename = file.Name()

		if filename == filepath.Base(templatePath) || !strings.Contains(filename, EnvFile) {
			continue
		}

		filePath = filepath.Join(envsDir, filename)

		err = backup.Create(filePath)

		if err != nil {
			return err
		}

		err = env.Sync(templateEnvs, filePath)

		if err != nil {
			return errors.Join(err, backup.Restore(filePath))
		}

		err = backup.Remove(filePath)

		if err != nil {
			return errors.Join(err, backup.Restore(filePath))
		}

		log.Printf("Synchronized: %q", filePath)
	}

	return nil
}
