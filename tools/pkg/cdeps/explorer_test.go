package cdeps_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/embed"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"

	"github.com/bastean/x/tools/pkg/cdeps"
)

type ExplorerTestSuite struct {
	suite.Default
	SUT                   *cdeps.Explorer
	targetDirectory, file string
	extensions            []string
}

func (s *ExplorerTestSuite) SetupSuite() {
	s.targetDirectory = "ignore"

	s.file = "ignore.test"

	s.extensions = []string{".min.js", ".min.css", ".woff2"}

	s.SUT = new(cdeps.Explorer)
}

func (s *ExplorerTestSuite) SetupTest() {
	_ = os.Remove(s.file)
	_ = os.RemoveAll(s.targetDirectory)
}

func (s *ExplorerTestSuite) TestCreateDirectory() {
	s.NoError(s.SUT.CreateDirectory(s.targetDirectory))
	s.DirExists(s.targetDirectory)
}

func (s *ExplorerTestSuite) TestCreateDirectoryErrFailedCreation() {
	directory := ""

	actual := s.SUT.CreateDirectory(directory)

	expected := fmt.Errorf("failed to create %q [%s]", directory, embed.Extract(actual.Error()))

	s.Equal(expected, actual)
}

func (s *ExplorerTestSuite) TestCopyFile() {
	sourceDirectory, file, expected := cdeps.Mother().FileValid(s.targetDirectory)

	s.NoError(s.SUT.CopyFile(file, sourceDirectory, s.targetDirectory))

	targetFile := filepath.Join(s.targetDirectory, file)

	s.FileExists(targetFile)

	actual, err := os.ReadFile(targetFile) //nolint:gosec

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *ExplorerTestSuite) TestCopyFileErrFailedReading() {
	sourceDirectory := cdeps.Mother().DirectoryInvalid(s.targetDirectory)

	actual := s.SUT.CopyFile(s.file, sourceDirectory, s.targetDirectory)

	expected := fmt.Errorf("failed to read %q from %q [%s]", s.file, sourceDirectory, embed.Extract(actual.Error()))

	s.Equal(expected, actual)
}

func (s *ExplorerTestSuite) TestCopyFileErrFailedWriting() {
	sourceDirectory, file, _ := cdeps.Mother().FileValid(s.targetDirectory)

	targetDirectory := cdeps.Mother().DirectoryInvalid(s.targetDirectory)

	actual := s.SUT.CopyFile(file, sourceDirectory, targetDirectory)

	expected := fmt.Errorf("failed to write %q on %q [%s]", file, targetDirectory, embed.Extract(actual.Error()))

	s.Equal(expected, actual)
}

func (s *ExplorerTestSuite) TestCopyDependency() {
	sourceDirectory, files := cdeps.Mother().FilesValid(s.targetDirectory, s.extensions)

	var (
		err              error
		expected, actual []string
	)

	for _, file := range files {
		actual, err = s.SUT.CopyDependency(file, sourceDirectory, s.targetDirectory)

		s.NoError(err)

		expected = []string{filepath.Join(s.targetDirectory, file)}

		s.Equal(expected, actual)
	}

	for _, file := range files {
		s.FileExists(filepath.Join(s.targetDirectory, file))
	}
}

func (s *ExplorerTestSuite) TestCopyDependencyRegexp() {
	const (
		RExEveryMinFile   = `^.+\.min\.(js|css)$`
		RExEveryWoff2File = `^.+\.woff2$`
	)

	regExp := cdeps.Mother().RandomString([]string{RExEveryMinFile, RExEveryWoff2File})

	sourceDirectory, files := cdeps.Mother().FilesValid(s.targetDirectory, s.extensions)

	var (
		err              error
		expected, actual []string
	)

	actual, err = s.SUT.CopyDependency(regExp, sourceDirectory, s.targetDirectory)

	s.NoError(err)

	expected = cdeps.Mother().FilesFilter(regExp, files, s.targetDirectory)

	s.ElementsMatch(expected, actual)

	for _, file := range expected {
		s.FileExists(file)
	}
}

func (s *ExplorerTestSuite) TestCopyDependencyErrFailed() {
	sourceDirectory := cdeps.Mother().DirectoryInvalid(s.targetDirectory)

	_, actual := s.SUT.CopyDependency(s.file, sourceDirectory, s.targetDirectory)

	expected := fmt.Errorf("failed to copy %q from %q [%s]", s.file, sourceDirectory, embed.Extract(actual.Error()))

	s.Equal(expected, actual)
}

func (s *ExplorerTestSuite) TearDownTest() {
	_ = os.Remove(s.file)
	_ = os.RemoveAll(s.targetDirectory)
}

func TestIntegrationExplorerSuite(t *testing.T) {
	suite.Run(t, new(ExplorerTestSuite))
}
