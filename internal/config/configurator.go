package config

import "github.com/spf13/viper"

type Entity struct {
	App App
	DB  DB
}

type App struct {
	Addr                 string
	Port                 string
	MaxConcurrentStreams int
}

type DB struct {
	Hostname string
	Port     string
	User     string
	Pass     string
	Name     string
}

func NewConfig() (*Entity, error) {
	viper.AutomaticEnv()
	config := &Entity{
		App: App{
			Addr:                 viper.GetString("APP_ADDR"),
			Port:                 viper.GetString("APP_PORT"),
			MaxConcurrentStreams: viper.GetInt("APP_MAX_STREAMS"),
		},
		DB: DB{
			Hostname: viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Pass:     viper.GetString("DB_PASS"),
			Name:     viper.GetString("DB_NAME"),
		},
	}

	return config, nil
}
