package tests_test

import (
	"fmt"
	"testing"

	"github.com/difmaj/cloudwalk-software-engineer-test/internal/models"
	"github.com/difmaj/cloudwalk-software-engineer-test/internal/report"
	"github.com/stretchr/testify/suite"
)

type SuiteReport struct {
	suite.Suite
}

func (s *SuiteReport) SetupSuite()    {}
func (s *SuiteReport) TearDownSuite() {}

func TestSuiteReport(t *testing.T) {
	suite.Run(t, new(SuiteReport))
}

func (s *SuiteReport) TestGenerateReport() {
	s.T().Run("success", func(t *testing.T) {
		response, err := report.Generate(&models.LogData{
			Games: []*models.LogGame{
				{
					InitGame: [][]byte{
						[]byte("\\g_time\\1553270153"),
					},
					Clients: []*models.LogClient{
						{
							ClientID: 1,
							UserName: []byte("Player 1"),
						},
						{
							ClientID: 2,
							UserName: []byte("Player 2"),
						},
					},
					Kills: []*models.LogKillEvent{
						{
							KillerID:     1,
							VictimID:     2,
							KillMethodID: 1,
						},
						{
							KillerID:     1,
							VictimID:     2,
							KillMethodID: 1,
						},
						{
							KillerID:     2,
							VictimID:     1,
							KillMethodID: 1,
						},
						{
							KillerID:     2,
							VictimID:     1,
							KillMethodID: 1,
						},
						{
							KillerID:     19,
							VictimID:     2,
							KillMethodID: 6,
						},
					},
				},
			},
		})

		s.Require().NoError(err)
		s.Require().NotNil(response)
		fmt.Println(response)
	})
}
