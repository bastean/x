package syncenv_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/x/tools/pkg/errors"
	"github.com/bastean/x/tools/pkg/syncenv"
)

type EnvTestSuite struct {
	suite.Suite
	SUT       *syncenv.Env
	directory string
}

func (s *EnvTestSuite) SetupSuite() {
	s.directory = "ignore"
	s.SUT = new(syncenv.Env)
}

func (s *EnvTestSuite) SetupTest() {
	s.NoError(os.RemoveAll(s.directory))
	s.NoError(os.MkdirAll(s.directory, 0700))
}

func (s *EnvTestSuite) TestDump() {
	envs := syncenv.RandomEnvs()

	source, file := syncenv.RandomEnvFile(strings.Join(envs, ""), s.directory)

	actual, err := s.SUT.Dump(filepath.Join(source, file))

	expected := strings.Split(strings.Join(envs, ""), "\n")

	s.NoError(err)

	s.Equal(expected, actual)
}

func (s *EnvTestSuite) TestDumpErrFailedReading() {
	file := syncenv.RandomUndefinedEnvFile(s.directory)

	path := syncenv.RandomUndefinedEnvPath(s.directory)

	source := filepath.Join(path, file)

	_, actual := s.SUT.Dump(source)

	expected := fmt.Errorf("failed to read %q [%s]", source, errors.Extract(actual))

	s.Equal(expected, actual)
}

func (s *EnvTestSuite) TestSync() {
	templateEnvs := syncenv.RandomTemplateEnvs()

	templateSource, templateFile := syncenv.RandomEnvFile(strings.Join(templateEnvs, ""), s.directory)

	templateEnvs, err := s.SUT.Dump(filepath.Join(templateSource, templateFile))

	s.NoError(err)

	targetEnvs := syncenv.RandomFileEnvs()

	targetSource, targetFile := syncenv.RandomEnvFile(strings.Join(targetEnvs, ""), s.directory)

	target := filepath.Join(targetSource, targetFile)

	s.NoError(s.SUT.Sync(templateEnvs, target))

	envs, err := os.ReadFile(target) //nolint:gosec

	s.NoError(err)

	syncEnvs := strings.Split(string(envs), "\n")

	for i, templateEnv := range templateEnvs {
		s.Contains(syncEnvs[i], templateEnv)
	}
}

func (s *EnvTestSuite) TestSyncErrOverwriting() {
	templateEnvs := syncenv.EnvsWithEmptyValues()

	targetSource, targetFile := syncenv.RandomEnvFile("", s.directory)

	target := filepath.Join(targetSource, targetFile)

	s.NoError(os.Chmod(target, 0400))

	actual := s.SUT.Sync(templateEnvs, target)

	expected := fmt.Errorf("failure to overwrite %q [%s]", target, errors.Extract(actual))

	s.Equal(expected, actual)
}

func (s *EnvTestSuite) TearDownTest() {
	s.NoError(os.Chmod(s.directory, 0700)) //nolint:gosec
	s.NoError(os.RemoveAll(s.directory))
}

func TestIntegrationEnvSuite(t *testing.T) {
	suite.Run(t, new(EnvTestSuite))
}
