package config

import "github.com/spf13/viper"

type Config struct {
	Logs Logs
}

type Logs struct {
	Path       string `yaml:"Path" json:"path"`
	Level      string `yaml:"Level" json:"level"`
	MaxAge     int    `yaml:"MaxAge" json:"max_age"`
	MaxBackups int    `yaml:"MaxBackups" json:"max_backups"`
}

func ReadConfig(cfgName, cfgType, cfgPath string) (*Config, error) {
	var cfg Config

	viper.SetConfigName(cfgName)
	viper.SetConfigType(cfgType)
	viper.AddConfigPath(cfgPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
