package config

import (
	"encoding/json"
	"os"
	"time"
)

func NewConfig(path string) (Config, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

type Config struct {
	Server Server `json:"server"`
	Logger Logger `json:"logger"`
	Mongo  Mongo  `json:"mongo"`
}

// Server holds data for server configuration
type Server struct {
	Port                string        `json:"port"`
	ReadTimeoutSeconds  time.Duration `json:"read_timeout_seconds"`
	WriteTimeoutSeconds time.Duration `json:"write_timeout_seconds"`
	Debug               bool          `json:"debug"`
}

// Mongo holds data for connecting to mongo cluster
type Mongo struct {
	AppName                    string           `json:"app_name"`
	Credentials                MongoCredentials `json:"credentials"`
	ConnectionTimeoutInSeconds int64            `json:"connection_timeout_in_seconds"`
	Hosts                      []string         `json:"hosts"`
	MaxPoolSize                uint64           `json:"max_pool_size"`
	DatabaseName               string           `json:"database_name"`
	Collections                Collections      `json:"collections"`
}

type Collections struct {
	Beer string `json:"beer"`
}

// MongoCredentials holds data for auth in mongo cluster
type MongoCredentials struct {
	AuthMechanism           string            `json:"auth_mechanism"`
	AuthMechanismProperties map[string]string `json:"auth_mechanism_properties"`
	AuthSource              string            `json:"auth_source"`
	Username                string            `json:"username"`
	Password                string            `json:"password"`
	PasswordSet             bool              `json:"password_set"`
}

// Logger holds data for logrus configuration
type Logger struct {
	TimestampFormat string   `json:"timestamp_format"`
	FieldsOrder     []string `json:"fields_order"`
	HideKeys        bool     `json:"hide_keys"`
	NoColors        bool     `json:"no_colors"`
	NoFieldsColors  bool     `json:"no_fields_colors"`
	ShowFullLevel   bool     `json:"show_full_level"`
	LogLevel        string   `json:"log_level"`
}
