package config

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

