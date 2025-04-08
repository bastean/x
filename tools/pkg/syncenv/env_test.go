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
	_ = os.RemoveAll(s.directory)
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

	targetEnvs := syncenv.RandomTargetEnvs()

	targetSource, targetFile := syncenv.RandomEnvFile(strings.Join(targetEnvs, ""), s.directory)

	targetEnvs, err = s.SUT.Dump(filepath.Join(targetSource, targetFile))

	s.NoError(err)

	synchronized := filepath.Join(targetSource, targetFile)

	s.NoError(s.SUT.Sync(templateEnvs, targetEnvs, synchronized))

	envs, err := os.ReadFile(synchronized) //nolint:gosec

	s.NoError(err)

	syncEnvs := strings.Split(string(envs), "\n")

	for i, templateEnv := range templateEnvs {
		s.Contains(syncEnvs[i], templateEnv)
	}
}

func (s *EnvTestSuite) TestSyncErrOverwriting() {
	templateEnvs := syncenv.EnvsWithEmptyValues()

	targetEnvs := syncenv.EnvsWithEmptyValues()

	synchronized := syncenv.RandomUndefinedEnvPath(s.directory)

	actual := s.SUT.Sync(templateEnvs, targetEnvs, synchronized)

	expected := fmt.Errorf("failure to overwrite %q [%s]", synchronized, errors.Extract(actual))

	s.Equal(expected, actual)
}

func (s *EnvTestSuite) TearDownTest() {
	_ = os.RemoveAll(s.directory)
}

func TestIntegrationEnvSuite(t *testing.T) {
	suite.Run(t, new(EnvTestSuite))
}
