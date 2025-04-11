package syncenv

import (
	"os"
	"path/filepath"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func RandomEnvValues(isTemplate bool) []string {
	var key, value string

	envs := make([]string, 0, 12)

	for range services.Create.IntRange(1, 12) {
		if services.Create.IntRange(1, 6) == 3 {
			envs = append(envs, "\n")
			continue
		}

		key = services.Create.LoremIpsumWord()

		if isTemplate {
			envs = append(envs, key+"="+"\n")
			continue
		}

		value = services.Create.LoremIpsumWord()

		envs = append(envs, key+"="+value+"\n")
	}

	return envs
}

func RandomEnvs() []string {
	return RandomEnvValues(services.Create.Bool())
}

func RandomTemplateEnvs() []string {
	return RandomEnvValues(true)
}

func RandomFileEnvs() []string {
	return RandomEnvValues(false)
}

func EnvsWithEmptyValues() []string {
	return []string{}
}

func RandomEnvFile(envs, path string) (string, string) {
	file := services.Create.LoremIpsumWord()

	err := os.WriteFile(filepath.Join(path, file), []byte(envs), 0600)

	if err != nil {
		panic(err.Error())
	}

	return path, file
}

func RandomUndefinedEnvFile(path string) string {
	return filepath.Join(path, services.Create.LoremIpsumWord())
}

func RandomUndefinedEnvPath(path string) string {
	return filepath.Join(path, services.Create.LoremIpsumWord())
}
