package models

type GameReportOut struct {
	Games map[string]*GameReportGameOut `json:"games"`
}

type GameReportGameOut struct {
	TotalKills   int                        `json:"total_kills"`
	Players      []*GameReportGamePlayerOut `json:"players"`
	Kills        map[string]int             `json:"kills"`
	KillsByMeans map[string]int             `json:"kills_by_means"`
}

type GameReportGamePlayerOut struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
