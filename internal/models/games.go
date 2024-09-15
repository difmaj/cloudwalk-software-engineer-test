package models

type GamesReportOut struct {
	Game map[string]*GamesReportGameOut `json:"games"`
}

type GamesReportGameOut struct {
	TotalKills   int            `json:"total_kills"`
	Players      []string       `json:"players"`
	Kills        map[string]int `json:"kills"`
	KillsByMeans map[string]int `json:"kills_by_means"`
}
