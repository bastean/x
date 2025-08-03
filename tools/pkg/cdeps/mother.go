package cdeps

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) FileValid(path string) (string, string, []byte) {
	path = filepath.Join(path, m.LoremIpsumWord())

	err := os.MkdirAll(path, 0700)

	if err != nil {
		panic(err.Error())
	}

	file := m.LoremIpsumWord() + m.ID() + "." + m.FileExtension()

	content := []byte(m.Message())

	err = os.WriteFile(filepath.Join(path, file), content, 0600)

	if err != nil {
		panic(err.Error())
	}

	return path, file, content
}

func (m *m) FilesValid(path string, extensions []string) (string, []string) {
	path = filepath.Join(path, m.LoremIpsumWord())

	err := os.MkdirAll(path, 0700)

	if err != nil {
		panic(err.Error())
	}

	files := make([]string, m.RandomInt([]int{1, 10}))

	for i := range len(files) {
		files[i] = m.LoremIpsumWord() + m.ID() + m.RandomString(extensions)

		err = os.WriteFile(filepath.Join(path, files[i]), []byte{}, 0600)

		if err != nil {
			panic(err.Error())
		}
	}

	return path, files
}

func (m *m) FilesFilter(filter string, files []string, path string) []string {
	var filtered []string

	isMatch := regexp.MustCompile(filter).MatchString

	for _, file := range files {
		if isMatch(file) {
			filtered = append(filtered, filepath.Join(path, file))
		}
	}

	return filtered
}

func (m *m) FileInvalid(path string) string {
	return filepath.Join(path, m.LoremIpsumWord()+"."+m.FileExtension())
}

func (m *m) DirectoryInvalid(path string) string {
	return filepath.Join(path, m.LoremIpsumWord())
}

var Mother = mother.New[m]
