package release

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrBumpInvalidVersion   = errors.New("version value is not valid")
	ErrBumpInvalidSemVer    = errors.New("version does not follow semver convention")
	ErrBumpInvalidIncrement = errors.New("invalid incremental value")
)

func BumpVersion(module *Module, latest string) (string, error) {
	version := strings.Split(latest, "v")

	if len(version) != 2 {
		return "", ErrBumpInvalidVersion
	}

	semver := strings.Split(version[1], ".")

	if len(semver) != 3 {
		return "", ErrBumpInvalidSemVer
	}

	major, errMajor := strconv.Atoi(semver[0])
	minor, errMinor := strconv.Atoi(semver[1])
	patch, errPatch := strconv.Atoi(semver[2])

	if err := errors.Join(errMajor, errMinor, errPatch); err != nil {
		return "", err
	}

	switch module.Increment {
	case "patch":
		patch++
	case "minor":
		patch = 0
		minor++
	case "major":
		patch = 0
		minor = 0
		major++
	default:
		return "", ErrBumpInvalidIncrement
	}

	return fmt.Sprintf("%d.%d.%d", major, minor, patch), nil
}
