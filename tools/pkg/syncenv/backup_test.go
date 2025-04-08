package syncenv_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/x/tools/pkg/errors"
	"github.com/bastean/x/tools/pkg/syncenv"
)

type BackupTestSuite struct {
	suite.Suite
	SUT       *syncenv.Backup
	directory string
}

func (s *BackupTestSuite) SetupSuite() {
	s.Equal(syncenv.ExtBackup, `.bak`)

	s.Equal(syncenv.RExBackup, `^.+\.bak$`)

	s.directory = "ignore"

	s.SUT = new(syncenv.Backup)
}

func (s *BackupTestSuite) SetupTest() {
	_ = os.RemoveAll(s.directory)
}

func (s *BackupTestSuite) TestFile() {
	source, file, expected := syncenv.RandomFile(s.directory)

	s.NoError(s.SUT.File(file, source, source))

	backup := filepath.Join(source, file+".bak")

	s.FileExists(backup)

	actual, err := os.ReadFile(backup) //nolint:gosec

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *BackupTestSuite) TestFileErrFailedReading() {
	file := syncenv.RandomUndefinedFileWithExtension(s.directory)

	source := syncenv.RandomUndefinedPath(s.directory)

	actual := s.SUT.File(file, source, source)

	expected := fmt.Errorf("failed to read %q from %q [%s]", file, source, errors.Extract(actual))

	s.Equal(expected, actual)
}

func (s *BackupTestSuite) TestFileErrFailedWriting() {
	source, file, _ := syncenv.RandomFile(s.directory)

	target := syncenv.RandomUndefinedPath(s.directory)

	actual := s.SUT.File(file, source, target)

	expected := fmt.Errorf("failed to write %q on %q [%s]", file+".bak", target, errors.Extract(actual))

	s.Equal(expected, actual)
}

func (s *BackupTestSuite) TestRestore() {
	source, original, expected := syncenv.RandomFile(s.directory)

	s.NoError(s.SUT.File(original, source, source))

	backup := original + ".bak"

	original = filepath.Join(source, original)

	s.NoError(os.Remove(original))

	s.NoFileExists(original)

	s.NoError(s.SUT.Restore(backup, source))

	s.FileExists(original)

	actual, err := os.ReadFile(original) //nolint:gosec

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *BackupTestSuite) TestRestoreErrFailure() {
	original := syncenv.RandomUndefinedFileWithExtension(s.directory)

	backup := original + ".bak"

	source := syncenv.RandomUndefinedPath(s.directory)

	actual := s.SUT.Restore(backup, source)

	expected := fmt.Errorf("failure to restore file %q from %q [%s]", original, backup, errors.Extract(actual))

	s.Equal(expected, actual)
}

func (s *BackupTestSuite) TestRemove() {
	source, file, _ := syncenv.RandomFile(s.directory)

	s.NoError(s.SUT.File(file, source, source))

	backup := file + ".bak"

	s.NoError(s.SUT.Remove(backup, source))

	s.NoFileExists(filepath.Join(source, backup))
}

func (s *BackupTestSuite) TestRemoveErrFailure() {
	backup := syncenv.RandomFilename() + ".bak"

	source := syncenv.RandomUndefinedPath(s.directory)

	actual := s.SUT.Remove(backup, source)

	expected := fmt.Errorf("failure to remove backup %q from %q [%s]", backup, source, errors.Extract(actual))

	s.Equal(expected, actual)
}

func (s *BackupTestSuite) TearDownTest() {
	_ = os.RemoveAll(s.directory)
}

func TestIntegrationBackupSuite(t *testing.T) {
	suite.Run(t, new(BackupTestSuite))
}
