package release

import (
	"errors"
	"os/exec"
)

type Exec struct{}

func (*Exec) Do(cmds ...string) (string, error) {
	output, err := exec.Command(cmds[0], cmds[1:]...).CombinedOutput() //nolint:gosec

	if err != nil {
		return "", errors.New(string(output))
	}

	return string(output), nil
}
