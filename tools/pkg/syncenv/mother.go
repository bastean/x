package syncenv

import (
	"os"
	"path/filepath"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) FileValid(path string) (string, string, []byte) {
	file := m.LoremIpsumWord()

	content := []byte(m.Message())

	err := os.WriteFile(filepath.Join(path, file), content, 0600)

	if err != nil {
		panic(err.Error())
	}

	return path, file, content
}

func (m *m) FilenameValid() string {
	return m.LoremIpsumWord()
}

func (m *m) FileInvalid(path string) string {
	return filepath.Join(path, m.LoremIpsumWord())
}

func (m *m) DirectoryInvalid(path string) string {
	return filepath.Join(path, m.LoremIpsumWord())
}

func (m *m) EnvValuesValid(isTemplate bool) []string {
	var key, value string

	envs := make([]string, 0, 12)

	for range m.IntRange(1, 12) {
		if m.IntRange(1, 6) == 3 {
			envs = append(envs, "\n")
			continue
		}

		key = m.LoremIpsumWord()

		if isTemplate {
			envs = append(envs, key+"="+"\n")
			continue
		}

		value = m.LoremIpsumWord()

		envs = append(envs, key+"="+value+"\n")
	}

	return envs
}

func (m *m) EnvsValuesValid() []string {
	return m.EnvValuesValid(m.Bool())
}

func (m *m) EnvsValuesTemplateValid() []string {
	return m.EnvValuesValid(true)
}

func (m *m) EnvsValuesFileValid() []string {
	return m.EnvValuesValid(false)
}

func (m *m) EnvsValuesInvalid() []string {
	return []string{}
}

func (m *m) EnvFileValid(envs, path string) (string, string) {
	file := m.LoremIpsumWord()

	err := os.WriteFile(filepath.Join(path, file), []byte(envs), 0600)

	if err != nil {
		panic(err.Error())
	}

	return path, file
}

var Mother = mother.New[m]
