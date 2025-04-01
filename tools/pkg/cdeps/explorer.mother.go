package cdeps

import (
	"os"
	"path/filepath"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func RandomFile(path string) (string, string, []byte) {
	path += "/random"

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

func RandomFiles(path string, extensions []string) (string, []string) {
	path += "/random"

	err := os.MkdirAll(path, 0700)

	if err != nil {
		panic(err.Error())
	}

	files := make([]string, services.Create.RandomInt([]int{1, 10}))

	for i := range len(files) {
		files[i] = services.Create.LoremIpsumWord() + services.Create.RandomString(extensions)

		err = os.WriteFile(filepath.Join(path, files[i]), []byte{}, 0600)

		if err != nil {
			panic(err.Error())
		}
	}

	return path, files
}
