package syncenv_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"

	"github.com/bastean/x/tools/pkg/errors"
	"github.com/bastean/x/tools/pkg/syncenv"
)

type BackupTestSuite struct {
	suite.Default
	SUT       *syncenv.Backup
	directory string
}

func (s *BackupTestSuite) SetupSuite() {
	s.directory = "ignore"
	s.SUT = new(syncenv.Backup)
}

func (s *BackupTestSuite) SetupTest() {
	s.NoError(os.RemoveAll(s.directory))
	s.NoError(os.MkdirAll(s.directory, 0700))
}

func (s *BackupTestSuite) TestSentinel() {
	s.Equal(".syncenv.bak", syncenv.ExtBackup)
}

func (s *BackupTestSuite) TestCreate() {
	source, file, expected := syncenv.Mother().RandomFile(s.directory)

	s.NoError(s.SUT.Create(filepath.Join(source, file)))

	backup := filepath.Join(source, file+".syncenv.bak")

	s.FileExists(backup)

	actual, err := os.ReadFile(backup) //nolint:gosec

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *BackupTestSuite) TestCreateErrFailedReading() {
	file := filepath.Join(syncenv.Mother().RandomUndefinedDir(s.directory), syncenv.Mother().RandomUndefinedFile(s.directory))

	actual := s.SUT.Create(file)

	expected := fmt.Errorf("failed to read %q [%s]", file, errors.Extract(actual))

	s.Equal(expected, actual)
}

func (s *BackupTestSuite) TestCreateErrFailedWriting() {
	source, file, _ := syncenv.Mother().RandomFile(s.directory)

	file = filepath.Join(source, file)

	s.NoError(os.WriteFile(file+".syncenv.bak", []byte(""), 0400))

	actual := s.SUT.Create(file)

	expected := fmt.Errorf("failed to write %q [%s]", file+".syncenv.bak", errors.Extract(actual))

	s.Equal(expected, actual)
}

func (s *BackupTestSuite) TestRestore() {
	source, file, expected := syncenv.Mother().RandomFile(s.directory)

	file = filepath.Join(source, file)

	s.NoError(s.SUT.Create(file))

	s.NoError(os.Remove(file))

	s.NoFileExists(file)

	s.NoError(s.SUT.Restore(file))

	s.FileExists(file)

	actual, err := os.ReadFile(file) //nolint:gosec

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *BackupTestSuite) TestRestoreErrFailure() {
	file := filepath.Join(syncenv.Mother().RandomUndefinedDir(s.directory), syncenv.Mother().RandomUndefinedFile(s.directory))

	actual := s.SUT.Restore(file)

	expected := fmt.Errorf("failure to restore file %q [%s]", file, errors.Extract(actual))

	s.Equal(expected, actual)
}

func (s *BackupTestSuite) TestRemove() {
	source, file, _ := syncenv.Mother().RandomFile(s.directory)

	backup := filepath.Join(source, file)

	s.NoError(s.SUT.Create(backup))

	s.NoError(s.SUT.Remove(backup))

	backup = filepath.Join(source, file+".syncenv.bak")

	s.NoFileExists(backup)
}

func (s *BackupTestSuite) TestRemoveErrFailure() {
	backup := filepath.Join(syncenv.Mother().RandomUndefinedDir(s.directory), syncenv.Mother().RandomFilename())

	actual := s.SUT.Remove(backup)

	expected := fmt.Errorf("failure to remove backup %q [%s]", backup+".syncenv.bak", errors.Extract(actual))

	s.Equal(expected, actual)
}

func (s *BackupTestSuite) TearDownTest() {
	s.NoError(os.RemoveAll(s.directory))
}

func TestIntegrationBackupSuite(t *testing.T) {
	suite.Run(t, new(BackupTestSuite))
}
