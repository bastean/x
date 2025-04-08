package syncenv

import (
	"os"
	"path/filepath"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func RandomFile(path string) (string, string, []byte) {
	path = filepath.Join(path, "random")

	err := os.MkdirAll(path, 0700)

	if err != nil {
		panic(err.Error())
	}

	file := services.Create.LoremIpsumWord() + ".random"

	content := []byte(services.Create.Message())

	err = os.WriteFile(filepath.Join(path, file), content, 0600)

	if err != nil {
		panic(err.Error())
	}

	return path, file, content
}

func RandomFilename() string {
	return services.Create.LoremIpsumWord()
}

func RandomUndefinedFile(path string) string {
	return filepath.Join(path, services.Create.LoremIpsumWord())
}

func RandomUndefinedFileWithExtension(path string) string {
	return filepath.Join(path, services.Create.LoremIpsumWord()+"."+services.Create.FileExtension())
}

func RandomUndefinedPath(path string) string {
	return filepath.Join(path, services.Create.LoremIpsumWord())
}
