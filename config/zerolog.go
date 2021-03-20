package config

type ZeroLog struct {
	Level                string `mapstructure,json,yaml:"level`
	CallerSkipFrameCount int    `mapstructure,json,yaml:"caller_skip_frame_count"`
	LogInConsole         bool   `mapstructure,json,yaml:"log_in_console"`
	LogFilePath          string `mapstructure,json,yaml:"log_file_path"`
}
