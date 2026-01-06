package syncenv

import (
	"fmt"
	"os"
	"strings"
)

type Env struct{}

func (*Env) Dump(file string) ([]string, error) {
	data, err := os.ReadFile(file) //nolint:gosec

	if err != nil {
		return nil, fmt.Errorf("failed to read %q [%s]", file, err)
	}

	envs := strings.Split(string(data), "\n")

	return envs, nil
}

func (e *Env) Sync(envs []string, file string) error {
	fileEnvs, err := e.Dump(file)

	if err != nil {
		return err
	}

	if len(envs) > 0 && envs[len(envs)-1] == "" {
		envs = envs[:len(envs)-1]
	}

	var syncEnvs strings.Builder
	var isNotUpdated bool

	for _, env := range envs {
		isNotUpdated = true

		if env == "" {
			syncEnvs.WriteString("\n")
			continue
		}

		for _, fileEnv := range fileEnvs {
			if strings.Contains(fileEnv, env) {
				syncEnvs.WriteString(fileEnv + "\n")
				isNotUpdated = false
				break
			}
		}

		if isNotUpdated {
			syncEnvs.WriteString(env + "\n")
		}
	}

	err = os.WriteFile(file, []byte(syncEnvs.String()), 0600)

	if err != nil {
		return fmt.Errorf("failure to overwrite %q [%s]", file, err)
	}

	return nil
}
