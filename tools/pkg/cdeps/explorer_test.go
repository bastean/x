package cdeps_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"

	"github.com/bastean/x/tools/pkg/cdeps"
	"github.com/bastean/x/tools/pkg/errors"
)

type ExplorerTestSuite struct {
	suite.Default
	SUT             *cdeps.Explorer
	directory, file string
	extensions      []string
}

func (s *ExplorerTestSuite) SetupSuite() {
	s.directory = "ignore"

	s.file = "ignore.test"

	s.extensions = []string{".min.js", ".min.css", ".woff2"}

	s.SUT = new(cdeps.Explorer)
}

func (s *ExplorerTestSuite) SetupTest() {
	_ = os.Remove(s.file)
	_ = os.RemoveAll(s.directory)
}

func (s *ExplorerTestSuite) TestCreateDirectory() {
	s.NoError(s.SUT.CreateDirectory(s.directory))
	s.DirExists(s.directory)
}

func (s *ExplorerTestSuite) TestCreateDirectoryErrFailedCreation() {
	directory := ""

	actual := s.SUT.CreateDirectory(directory)

	expected := fmt.Errorf("failed to create %q [%s]", directory, errors.Extract(actual))

	s.Equal(expected, actual)
}

func (s *ExplorerTestSuite) TestCopyFile() {
	path, file, expected := cdeps.Mother().RandomFile(s.directory)

	s.NoError(s.SUT.CopyFile(file, path, s.directory))

	path = filepath.Join(s.directory, file)

	s.FileExists(path)

	actual, err := os.ReadFile(path) //nolint:gosec

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *ExplorerTestSuite) TestCopyFileErrFailedReading() {
	actual := s.SUT.CopyFile(s.file, s.directory, "./")

	expected := fmt.Errorf("failed to read %q from %q [%s]", s.file, s.directory, errors.Extract(actual))

	s.Equal(expected, actual)
}

func (s *ExplorerTestSuite) TestCopyFileErrFailedWriting() {
	path, file, _ := cdeps.Mother().RandomFile(s.directory)

	directory := cdeps.Mother().RandomUndefinedDir(s.directory)

	actual := s.SUT.CopyFile(file, path, directory)

	expected := fmt.Errorf("failed to write %q on %q [%s]", file, directory, errors.Extract(actual))

	s.Equal(expected, actual)
}

func (s *ExplorerTestSuite) TestCopyDependency() {
	path, files := cdeps.Mother().RandomFiles(s.directory, s.extensions)

	for _, file := range files {
		s.NoError(s.SUT.CopyDependency(file, path, s.directory))
	}

	for _, file := range files {
		s.FileExists(filepath.Join(s.directory, file))
	}
}

func (s *ExplorerTestSuite) TestCopyDependencyRegexp() {
	const (
		everyMinFile   = `^.+\.min\.(js|css)$`
		everyWoff2File = `^.+\.woff2$`
	)

	path, files := cdeps.Mother().RandomFiles(s.directory, s.extensions)

	s.NoError(s.SUT.CopyDependency(everyMinFile, path, s.directory))

	s.NoError(s.SUT.CopyDependency(everyWoff2File, path, s.directory))

	for _, file := range files {
		s.FileExists(filepath.Join(s.directory, file))
	}
}

func (s *ExplorerTestSuite) TestCopyDependencyErrFailed() {
	file := s.file

	source := cdeps.Mother().RandomUndefinedFile(s.directory)

	actual := s.SUT.CopyDependency(file, source, s.directory)

	expected := fmt.Errorf("failed to copy %q from %q [%s]", file, source, errors.Extract(actual))

	s.Equal(expected, actual)
}

func (s *ExplorerTestSuite) TearDownTest() {
	_ = os.Remove(s.file)
	_ = os.RemoveAll(s.directory)
}

func TestIntegrationExplorerSuite(t *testing.T) {
	suite.Run(t, new(ExplorerTestSuite))
}
