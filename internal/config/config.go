package config

import (
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Config struct {
	Server   HTTPServerConfiguration `mapstructure:"server"`
	Database DatabaseConfiguration   `mapstructure:"db"`
	Logging  LoggingConfiguration    `mapstructure:"logging"`
	Football FootballConfiguration   `mapstructure:"football"`
}

type HTTPServerConfiguration struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type DatabaseConfiguration struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type LoggingConfiguration struct {
	Level string `mapstructure:"level"`
	File  string `mapstructure:"file"`
}

type FootballConfiguration struct {
	APIKey  string `mapstructure:"api_key" envconfig:"FOOTBALL_API_KEY"`
	BaseURL string `mapstructure:"base_url"`
}

func Load() *Config {
	viper.SetConfigName("maka")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.maka")
	viper.SetEnvPrefix("MAKA")
	viper.AutomaticEnv()

	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to read config:", err)
		os.Exit(1)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println("Failed to unmarshal config:", err)
		os.Exit(1)
	}
	return &cfg

}

func (db DatabaseConfiguration) OpenDB() *sqlx.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		db.Username, db.Password, db.Host, db.Port, db.Database)
	conn := sqlx.MustConnect("pgx", dsn)
	conn.SetConnMaxLifetime(time.Hour)
	return conn
}
