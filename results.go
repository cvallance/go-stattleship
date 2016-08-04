package stattleship

type HeaderDetails struct {
	PerPage int
	Total   int
}

type GamesResult struct {
	HomeTeams    []*Team   `json:"home_teams"`
	AwayTeams    []*Team   `json:"away_teams"`
	Leagues      []*League `json:"leagues"`
	WinningTeams []*Team   `json:"winning_teams"`
	Seasons      []*Season `json:"seasons"`
	Venues       []*Venue  `json:"venues"`
	Games        []*Game   `json:"games"`
}
