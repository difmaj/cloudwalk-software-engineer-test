package tests_test

import (
	"testing"

	"github.com/difmaj/cloudwalk-software-engineer-test/internal/models"
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

func (s *SuiteParser) TestParseInitGameEventHandler() {
	tests := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{
			name:    "Valid InitGame Event",
			input:   []byte("\\g_time\\1553270153"),
			wantErr: false,
		},
		{
			name:    "Empty InitGame Event",
			input:   []byte(""),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			logData := &models.LogData{}
			err := parser.ParseInitGameEventHandler(tt.input, logData)
			s.Require().Equal(tt.wantErr, err != nil, "ParseInitGameEventHandler() error = %v, wantErr %v", err, tt.wantErr)
			s.Require().Equal(len(logData.Games), 1, "Expected only one game to be added")
		})
	}
}

func (s *SuiteParser) TestParseClientConectedEventHandler() {
	tests := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{
			name:    "Valid ClientConnect Event",
			input:   []byte("2"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			logData := &models.LogData{
				Current: &models.LogGame{
					Clients: []*models.LogClient{},
				},
			}

			err := parser.ParseClientConectedEventHandler(tt.input, logData)
			s.Require().Equal(tt.wantErr, err != nil, "ParseClientConectedEventHandler() error = %v, wantErr %v", err, tt.wantErr)
			s.Require().Equal(len(logData.Current.Clients), 1, "Expected only one client to be added")
		})
	}
}

func (s *SuiteParser) TestParseClientUserinfoChangedEventHandler() {
	tests := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{
			name:    "Valid ClientUserinfoChanged Event",
			input:   []byte("2 \\Player1"),
			wantErr: false,
		},
		{
			name:    "Invalid ClientUserinfoChanged Event with no game",
			input:   []byte("2 \\Player1"),
			wantErr: true,
		},
		{
			name:    "Invalid ClientUserinfoChanged Event with no data",
			input:   []byte("2"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			logData := &models.LogData{
				Current: &models.LogGame{
					Clients: []*models.LogClient{
						{ClientID: 2},
					},
				},
			}
			if tt.wantErr {
				logData.Current = nil
			}

			err := parser.ParseClientUserinfoChangedEventHandler(tt.input, logData)
			s.Require().Equal(tt.wantErr, err != nil, "ParseClientUserinfoChangedEventHandler() error = %v, wantErr %v", err, tt.wantErr)

			if !tt.wantErr {
				client := logData.Current.Clients[0]
				s.Require().NotNil(client.UserName, "Expected client username to be set")
				s.Require().Equal(string(client.UserName), "Player1", "Expected client username to be 'Player1', got %s", string(client.UserName))
			}
		})
	}
}

func (s *SuiteParser) TestParseKillEventHandler() {
	tests := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{
			name:    "Valid Kill Event",
			input:   []byte("2 3 6: killed 3 using MOD_ROCKET"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			logData := &models.LogData{
				Current: &models.LogGame{
					Kills: []*models.LogKillEvent{},
				},
			}

			err := parser.ParseKillEventHandler(tt.input, logData)
			s.Require().Equal(tt.wantErr, err != nil, "ParseKillEventHandler() error = %v, wantErr %v", err, tt.wantErr)
			s.Require().Greater(len(logData.Current.Kills), 0, "Expected kill event to be added")
		})
	}
}
