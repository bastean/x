package cdeps_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/x/tools/pkg/cdeps"
)

type ExplorerTestSuite struct {
	suite.Suite
	SUT             *cdeps.Explorer
	directory, file string
	extensions      []string
}

func (s *ExplorerTestSuite) SetupSuite() {
	s.Equal(cdeps.EveryMinFile, `^.+\.min\.(js|css)$`)
	s.Equal(cdeps.EveryWoff2File, `^.+\.woff2$`)

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

	expected := fmt.Errorf("failed to create \"%s\"", directory)

	s.Equal(expected, actual)
}

func (s *ExplorerTestSuite) TestCopyFile() {
	path, file, expected := cdeps.RandomFile(s.directory)

	s.NoError(s.SUT.CopyFile(file, path, s.directory))

	path = filepath.Join(s.directory, file)

	s.FileExists(path)

	actual, err := os.ReadFile(path) //nolint:gosec

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *ExplorerTestSuite) TestCopyFileErrFailedReading() {
	actual := s.SUT.CopyFile(s.file, s.directory, "./")

	expected := fmt.Errorf("failed to read \"%s\" from \"%s\"", s.file, s.directory)

	s.Equal(expected, actual)
}

func (s *ExplorerTestSuite) TestCopyFileErrFailedWriting() {
	path, file, _ := cdeps.RandomFile(s.directory)

	directory := "undefined"

	actual := s.SUT.CopyFile(file, path, directory)

	expected := fmt.Errorf("failed to write \"%s\" on \"%s\"", file, directory)

	s.Equal(expected, actual)
}

func (s *ExplorerTestSuite) TestCopyDeps() {
	path, files := cdeps.RandomFiles(s.directory, s.extensions)

	s.NoError(s.SUT.CopyDeps(files, path, s.directory))

	for _, file := range files {
		s.FileExists(filepath.Join(s.directory, file))
	}
}

func (s *ExplorerTestSuite) TestCopyDepsMin() {
	path, files := cdeps.RandomFiles(s.directory, s.extensions)

	s.NoError(s.SUT.CopyDeps([]string{cdeps.EveryMinFile}, path, s.directory))

	s.NoError(s.SUT.CopyDeps([]string{cdeps.EveryWoff2File}, path, s.directory))

	for _, file := range files {
		s.FileExists(filepath.Join(s.directory, file))
	}
}

func (s *ExplorerTestSuite) TestCopyDepsErrFailed() {
	files := []string{s.file}

	source := "undefined"

	actual := s.SUT.CopyDeps(files, source, s.directory)

	expected := fmt.Errorf("failed to copy \"%s\" from \"%s\"", files, source)

	s.Equal(expected, actual)
}

func (s *ExplorerTestSuite) TearDownTest() {
	_ = os.Remove(s.file)
	_ = os.RemoveAll(s.directory)
}

func TestIntegrationExplorerSuite(t *testing.T) {
	suite.Run(t, new(ExplorerTestSuite))
}
