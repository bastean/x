package release

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func RandomTag(module *Module) (string, string, int, int, int) {
	version := services.Create.AppVersion()

	semver := strings.Split(version, ".")

	major, errMajor := strconv.Atoi(semver[0])
	minor, errMinor := strconv.Atoi(semver[1])
	patch, errPatch := strconv.Atoi(semver[2])

	if err := errors.Join(errMajor, errMinor, errPatch); err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/v%s", module.Name, version), version, major, minor, patch
}
