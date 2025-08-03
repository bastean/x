package release

import (
	"errors"
	"os/exec"
	"strings"
)

type Exec struct{}

func (*Exec) Do(cmds ...string) (string, error) {
	output, err := exec.Command(cmds[0], cmds[1:]...).CombinedOutput() //nolint:gosec

	if err != nil {
		return "", errors.New(strings.TrimSuffix(string(output), "\n"))
	}

	return string(output), nil
}
