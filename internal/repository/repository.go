package repository

import "time"

// Repository represents a repository as it's stored in the database
type Repository struct {
	Owner               Owner            `json:"owner"`
	Name                string           `json:"name"`
	Description         string           `json:"description"`
	GitHubURL           string           `json:"githubURL"`
	StargazersCount     int              `json:"stargazersCount"`
	WeekStargazersCount int              `json:"weekStargazersCount"`
	GitHubCreatedAt     time.Time        `json:"githubCreatedAt"`
	LastCommitAt        time.Time        `json:"lastCommitAt"`
	VimColorSchemes     []VimColorScheme `json:"vimColorSchemes"`
}

// Owner represents the owner of a repository
type Owner struct {
	Name string `json:"name"`
}

// VimColorScheme represents a vim color scheme's meta data
type VimColorScheme struct {
	Name        string             `json:"name"`
	Data        VimColorSchemeData `json:"data"`
	Valid       bool               `json:"valid"`
	Backgrounds []string           `json:"backgrounds"`
}

// VimColorSchemeData represents the color values for light and dark backgrounds
type VimColorSchemeData struct {
	Light []VimColorSchemeGroup `json:"light,omitempty"`
	Dark  []VimColorSchemeGroup `json:"dark,omitempty"`
}

// VimColorSchemeGroup represents a vim color scheme group's data
type VimColorSchemeGroup struct {
	Name    string `json:"name"`
	HexCode string `json:"hexCode"`
}
