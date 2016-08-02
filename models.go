package stattleship

import "time"

type Game struct {
	Id                string     `json:"id"`
	CreatedAt         *time.Time `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
	AtNeutralSite     bool       `json:"at_neutral_site"`
	Attendance        int        `json:"attendance"`
	AwayTeamOutcome   string     `json:"away_team_outcome"`
	AwayTeamScore     int        `json:"away_team_score"`
	Daytime           bool       `json:"daytime"`
	Duration          int        `json:"duration"`
	EndedAt           *time.Time `json:"ended_at"`
	HomeTeamOutcome   string     `json:"home_team_outcome"`
	HomeTeamScore     int        `json:"home_team_score"`
	Interval          string     `json:"interval"`
	IntervalNumber    int        `json:"interval_number"`
	IntervalType      string     `json:"interval_type"`
	Label             string     `json:"label"`
	Name              string     `json:"name"`
	On                string     `json:"on"`
	Score             string     `json:"score"`
	ScoreDifferential int        `json:"score_differential"`
	ScoreLine         string     `json:"score_line"`
	Slug              string     `json:"slug"`
	StartedAt         *time.Time `json:"started_at"`
	Temperature       int        `json:"temperature"`
	TemperatureUnit   string     `json:"temperature_unit"`
	Timestamp         int        `json:"timestamp"`
	Title             string     `json:"title"`
	WeatherConditions string     `json:"weather_conditions"`
	WindDirection     string     `json:"wind_direction"`
	WindSpeed         int        `json:"wind_speed"`
	WindSpeedUnit     string     `json:"wind_speed_unit"`
	HomeTeamId        string     `json:"home_team_id"`
	AwayTeamId        string     `json:"away_team_id"`
	WinningTeamId     string     `json:"winning_team_id"`
	SeasonId          string     `json:"season_id"`
	VenueId           string     `json:"venue_id"`
}

type League struct {
	Id               string      `json:"id"`
	CreatedAt        *time.Time  `json:"created_at"`
	UpdatedAt        *time.Time  `json:"updated_at"`
	Abbreviation     string      `json:"abbreviation"`
	Color            interface{} `json:"color"`
	MinutesPerPeriod interface{} `json:"minutes_per_period"`
	Name             string      `json:"name"`
	Periods          interface{} `json:"periods"`
	Slug             string      `json:"slug"`
	Sport            string      `json:"sport"`
}

type Season struct {
	Id        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Name      string     `json:"name"`
	StartsOn  string     `json:"starts_on"`
	EndsOn    string     `json:"ends_on"`
	Slug      string     `json:"slug"`
	LeagueId  string     `json:"league_id"`
}

type Team struct {
	Id         string     `json:"id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	Location   string     `json:"location"`
	Name       string     `json:"name"`
	Nickname   string     `json:"nickname"`
	Slug       string     `json:"slug"`
	DivisionId string     `json:"division_id"`
	LeagueId   string     `json:"league_id"`
}

type Venue struct {
	Id           string     `json:"id"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	Abbreviation string     `json:"abbreviation"`
	City         string     `json:"city"`
	Country      string     `json:"country"`
	FieldType    string     `json:"field_type"`
	Name         string     `json:"name"`
	Slug         string     `json:"slug"`
	State        string     `json:"state"`
	TimeZone     string     `json:"time_zone"`
}
