package cdeps

import (
	"os"
	"path/filepath"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) FileValid(path string) (string, string, []byte) {
	path += "/random"

	err := os.MkdirAll(path, 0700)

	if err != nil {
		panic(err.Error())
	}

	file := m.LoremIpsumWord() + ".random"

	content := []byte(m.Message())

	err = os.WriteFile(filepath.Join(path, file), content, 0600)

	if err != nil {
		panic(err.Error())
	}

	return path, file, content
}

func (m *m) FilesValid(path string, extensions []string) (string, []string) {
	path += "/random"

	err := os.MkdirAll(path, 0700)

	if err != nil {
		panic(err.Error())
	}

	files := make([]string, m.RandomInt([]int{1, 10}))

	for i := range len(files) {
		files[i] = m.LoremIpsumWord() + m.RandomString(extensions)

		err = os.WriteFile(filepath.Join(path, files[i]), []byte{}, 0600)

		if err != nil {
			panic(err.Error())
		}
	}

	return path, files
}

func (m *m) FileInvalid(path string) string {
	return filepath.Join(path, m.LoremIpsumWord())
}

func (m *m) DirectoryInvalid(path string) string {
	return filepath.Join(path, m.LoremIpsumWord())
}

var Mother = mother.New[m]
