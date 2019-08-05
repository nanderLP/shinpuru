package core

import (
	"io"
)

type ConfigDiscord struct {
	Token         string
	GeneralPrefix string
	OwnerID       string
	ClientID      string
	ClientSecret  string
}

type ConfigDatabaseCreds struct {
	Host     string
	User     string
	Password string
	Database string
}

type ConfigDatabaseFile struct {
	DBFile string
}

type ConfigDatabaseType struct {
	Type   string
	MySql  *ConfigDatabaseCreds
	Sqlite *ConfigDatabaseFile
}

type ConfigLogging struct {
	CommandLogging bool
	LogLevel       int
}

type ConfigEtc struct {
	TwitchAppID string
}

type ConfigWS struct {
	Addr       string       `json:"addr"`
	TLS        *ConfigWSTLS `json:"tls"`
	PublicAddr string       `json:"publicaddr"`
}

type ConfigWSTLS struct {
	Enabled bool   `json:"enabled"`
	Cert    string `json:"certfile"`
	Key     string `json:"keyfile"`
}

type Config struct {
	Version   int `yaml:"configVersionPleaseDoNotChange"`
	Discord   *ConfigDiscord
	Database  *ConfigDatabaseType
	Logging   *ConfigLogging
	Etc       *ConfigEtc
	WebServer *ConfigWS
}

type ConfigParser interface {
	Decode(r io.Reader) (*Config, error)
	Encode(w io.Writer, c *Config) error
}

func NewDefaultConfig() *Config {
	return &Config{
		Version: 5,
		Discord: &ConfigDiscord{
			GeneralPrefix: "sp!",
		},
		Database: &ConfigDatabaseType{
			Type:  "sqlite",
			MySql: new(ConfigDatabaseCreds),
			Sqlite: &ConfigDatabaseFile{
				DBFile: "shinpuru.sqlite3.db",
			},
		},
		Logging: &ConfigLogging{
			CommandLogging: true,
			LogLevel:       4,
		},
		Etc: new(ConfigEtc),
		WebServer: &ConfigWS{
			Addr:       ":8080",
			PublicAddr: "https://example.com:8080",
			TLS: &ConfigWSTLS{
				Enabled: false,
			},
		},
	}
}
