package syncenv

import (
	"os"
	"path/filepath"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) RandomFile(path string) (string, string, []byte) {
	file := m.LoremIpsumWord()

	content := []byte(m.Message())

	err := os.WriteFile(filepath.Join(path, file), content, 0600)

	if err != nil {
		panic(err.Error())
	}

	return path, file, content
}

func (m *m) RandomFilename() string {
	return m.LoremIpsumWord()
}

func (m *m) RandomUndefinedFile(path string) string {
	return filepath.Join(path, m.LoremIpsumWord())
}

func (m *m) RandomUndefinedDir(path string) string {
	return filepath.Join(path, m.LoremIpsumWord())
}

func (m *m) RandomEnvValues(isTemplate bool) []string {
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

func (m *m) RandomEnvs() []string {
	return m.RandomEnvValues(m.Bool())
}

func (m *m) RandomTemplateEnvs() []string {
	return m.RandomEnvValues(true)
}

func (m *m) RandomFileEnvs() []string {
	return m.RandomEnvValues(false)
}

func (m *m) EnvsWithEmptyValues() []string {
	return []string{}
}

func (m *m) RandomEnvFile(envs, path string) (string, string) {
	file := m.LoremIpsumWord()

	err := os.WriteFile(filepath.Join(path, file), []byte(envs), 0600)

	if err != nil {
		panic(err.Error())
	}

	return path, file
}

var Mother = mother.New[m]
