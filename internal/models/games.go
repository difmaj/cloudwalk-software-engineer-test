package models

type MatchReportOut struct {
	Game map[string]*MatchReportGameOut `json:"games"`
}

type MatchReportGameOut struct {
	TotalKills   int            `json:"total_kills"`
	Players      []string       `json:"players"`
	Kills        map[string]int `json:"kills"`
	KillsByMeans map[string]int `json:"kills_by_means"`
}
