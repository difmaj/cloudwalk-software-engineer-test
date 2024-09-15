package tests_test

import (
	"testing"

	"github.com/difmaj/cloudwalk-software-engineer-test/internal/parser"
	"github.com/stretchr/testify/suite"
)

type SuiteParser struct {
	suite.Suite
}

func (s *SuiteParser) SetupSuite()    {}
func (s *SuiteParser) TearDownSuite() {}

func TestSuiteParser(t *testing.T) {
	suite.Run(t, new(SuiteParser))
}

func (s *SuiteParser) TestParseLog() {
	s.T().Run("success", func(t *testing.T) {
		logData, err := parser.ParseLog("data/valid.log")

		s.Require().NoError(err)
		s.Require().NotNil(logData)
		s.Require().Greater(len(logData.Games), 0)
	})

	s.T().Run("file not found", func(t *testing.T) {
		logData, err := parser.ParseLog("data/non_existent_file.log")

		s.Require().Error(err)
		s.Require().Nil(logData)
	})
}
