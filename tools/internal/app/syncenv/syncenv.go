package syncenv

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bastean/x/tools/internal/pkg/cli"
	"github.com/bastean/x/tools/internal/pkg/errs"
	"github.com/bastean/x/tools/pkg/syncenv"
)

const (
	EnvFile = ".env"
)

var (
	templateFile, envsDir string
)

func Init() error {
	flag.StringVar(&templateFile, "t", "", "Path to \".env\" file template (required)")

	flag.StringVar(&envsDir, "e", "", "Path to \".env\" files directory (required)")

	flag.Usage = func() {
		cli.Usage("syncenv")
	}

	flag.Parse()

	if flag.NFlag() < 2 {
		flag.Usage()

		println()

		return errs.ErrRequiredFlags
	}

	return nil
}

func Up() error {
	err := Init()

	if err != nil {
		return err
	}

	log.Println("Starting...")

	env := new(syncenv.Env)

	templateEnvs, err := env.Dump(templateFile)

	if err != nil {
		return err
	}

	log.Printf("Template: %q", templateFile)

	files, err := os.ReadDir(envsDir)

	if err != nil {
		return fmt.Errorf("failure to read files from %q [%s]", envsDir, err)
	}

	backup := new(syncenv.Backup)

	var filename, filePath string

	for _, file := range files {
		filename = file.Name()

		if filename == filepath.Base(templateFile) || !strings.Contains(filename, EnvFile) {
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

	log.Println("Completed!")

	return nil
}
