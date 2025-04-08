package syncenv

import (
	"fmt"
	"os"
	"strings"
)

type Env struct{}

func (*Env) Dump(path string) ([]string, error) {
	data, err := os.ReadFile(path) //nolint:gosec

	if err != nil {
		return nil, fmt.Errorf("failed to read %q [%s]", path, err)
	}

	envs := strings.Split(string(data), "\n")

	return envs, nil
}

func (e *Env) Sync(templateEnvs, targetEnvs []string, targetPath string) error {
	var syncEnvs string
	var isNotUpdated bool

	for _, templateEnv := range templateEnvs {
		isNotUpdated = true

		if templateEnv == "" {
			syncEnvs += "\n"
			continue
		}

		for _, targetEnv := range targetEnvs {
			if strings.Contains(targetEnv, templateEnv) {
				syncEnvs += targetEnv + "\n"
				isNotUpdated = false
				break
			}
		}

		if isNotUpdated {
			syncEnvs += templateEnv + "\n"
		}
	}

	err := os.WriteFile(targetPath, []byte(syncEnvs), 0600)

	if err != nil {
		return fmt.Errorf("failure to overwrite %q [%s]", targetPath, err)
	}

	return nil
}
