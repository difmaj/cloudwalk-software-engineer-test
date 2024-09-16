package report

import (
	"encoding/json"
	"fmt"

	"github.com/difmaj/cloudwalk-software-engineer-test/internal/models"
)

func Generate(logData *models.LogData) (string, error) {
	response := &models.GameReportOut{
		Games: make(map[string]*models.GameReportGameOut, 0),
	}

	for index, game := range logData.Games {
		gameReport := &models.GameReportGameOut{
			TotalKills:   0,
			Players:      make([]*models.GameReportGamePlayerOut, 0),
			Kills:        make(map[string]int),
			KillsByMeans: make(map[string]int),
		}

		for _, client := range game.Clients {
			gameReport.Players = append(gameReport.Players, &models.GameReportGamePlayerOut{
				ID:   client.ClientID,
				Name: string(client.UserName),
			})
		}

		for _, kill := range game.Kills {
			gameReport.TotalKills++

			var killer *models.LogClient
			for _, client := range game.Clients {
				if client.ClientID == kill.KillerID {
					killer = client
					break
				}
			}

			if killer != nil {
				gameReport.Kills[string(killer.UserName)]++
			} else {

				var victim *models.LogClient
				for _, client := range game.Clients {
					if client.ClientID == kill.VictimID {
						victim = client
						break
					}
				}

				if victim != nil {
					gameReport.Kills[string(victim.UserName)]--
				}
			}
			gameReport.KillsByMeans[kill.KillMethodID.String()]++
			response.Games[fmt.Sprintf("game-%d", index)] = gameReport
		}
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(jsonResponse), nil
}
