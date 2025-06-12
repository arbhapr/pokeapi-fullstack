package model

// Config holds the configuration values for the application
type Config struct {
	APIVersion   string
	Environment  string
	BaseURL      string
	SourceURL    string
	Port         string
	LimitPerPage int
}
