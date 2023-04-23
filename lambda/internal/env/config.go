package env

import "github.com/caarlos0/env/v8"

// Config contains all information needed to configure an application.
type Config struct {
	NotionAuthToken string `env:"NOTION_AUTH_TOKEN"`
	PostsDatabaseID string `env:"POSTS_DB"`
	Port            int    `env:"PORT" envDefault:"8008"`
	IsDev           bool   `env:"DEV"`
}

// NewConfig creates a new config by reading environmental variables.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// GetNotionAuthToken returns notion auth token.
func (cfg Config) GetNotionAuthToken() string {
	return cfg.NotionAuthToken
}

// GetPostsDatabaseID returns id of posts database in notion.
func (cfg Config) GetPostsDatabaseID() string {
	return cfg.PostsDatabaseID
}
