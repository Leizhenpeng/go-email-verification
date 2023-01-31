package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbUrI     string `mapstructure:"MONGODB_URI"`
	RedisUrI  string `mapstructure:"REDIS_URI"`
	RedisPass string `mapstructure:"REDIS_PASSWORD"`
	RedisDb   int    `mapstructure:"REDIS_DB"`

	Port string `mapstructure:"PORT"`

	JwtKey        string `mapstructure:"JWT_KEY"`
	JwtAccessAge  int    `mapstructure:"JWT_ACCESS_AGE"`
	JwtRefreshAge int    `mapstructure:"JWT_FRESH_AGE"`

	EmailFrom string `mapstructure:"EMAIL_FROM"`
	SmtpHost  string `mapstructure:"SMTP_HOST"`
	SmtpUser  string `mapstructure:"SMTP_USER"`
	SmtpPass  string `mapstructure:"SMTP_PASS"`
	SmtpPort  string `mapstructure:"SMTP_PORT"`

	Origin string `mapstructure:"CLIENT_ORIGIN"`
}

var config *Config

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	return nil
}

func GetConfig() *Config {
	return config
}
