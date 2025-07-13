package release_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"

	"github.com/bastean/x/tools/pkg/release"
)

type BumpTestSuite struct {
	suite.Default
}

func (s *BumpTestSuite) TestBumpVersionFirstRelease() {
	module := release.Mother().RandomModuleFirstRelease()

	latest := module.Name + "/v0.0.0"

	expected := "0.1.0"

	actual, err := release.BumpVersion(module, latest)

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *BumpTestSuite) TestBumpVersionRelease() {
	module := release.Mother().RandomModuleRelease()

	latest, _, major, minor, patch := release.Mother().RandomTag(module)

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
	}

	expected := fmt.Sprintf("%d.%d.%d", major, minor, patch)

	actual, err := release.BumpVersion(module, latest)

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *BumpTestSuite) TestBumpVersionErrInvalidVersion() {
	module := release.Mother().RandomModuleRelease()

	latest := module.Name

	_, actual := release.BumpVersion(module, latest)

	s.ErrorIs(actual, release.ErrBumpInvalidVersion)

	expected := errors.New("version value is not valid")

	s.Equal(expected, actual)
}

func (s *BumpTestSuite) TestBumpVersionErrInvalidSemVer() {
	module := release.Mother().RandomModuleRelease()

	latest := module.Name + "/v0"

	_, actual := release.BumpVersion(module, latest)

	s.ErrorIs(actual, release.ErrBumpInvalidSemVer)

	expected := errors.New("version does not follow semver convention")

	s.Equal(expected, actual)
}

func (s *BumpTestSuite) TestBumpVersionErrInvalidIncrement() {
	module, _ := release.Mother().ModuleWithInvalidIncrement()

	latest, _, _, _, _ := release.Mother().RandomTag(module)

	_, actual := release.BumpVersion(module, latest)

	s.ErrorIs(actual, release.ErrBumpInvalidIncrement)

	expected := errors.New("invalid incremental value")

	s.Equal(expected, actual)
}

func TestUnitBumpSuite(t *testing.T) {
	suite.Run(t, new(BumpTestSuite))
}
