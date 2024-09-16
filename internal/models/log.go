package models

import "github.com/difmaj/cloudwalk-software-engineer-test/internal/models/enums"

type LogData struct {
	Games   []*LogGame
	Current *LogGame
}

type LogGame struct {
	InitGame [][]byte
	Clients  []*LogClient
	Kills    []*LogKillEvent
}

type LogClient struct {
	ClientID int
	UserName []byte
}

type LogKillEvent struct {
	KillerID     int
	VictimID     int
	KillMethodID enums.DeathMeansID
}
