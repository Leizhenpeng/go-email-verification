package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DbUrI     string `mapstructure:"MONGODB_URI"`
	RedisUrI  string `mapstructure:"REDIS_URI"`
	RedisPass string `mapstructure:"REDIS_PASSWORD"`
	RedisDb   int    `mapstructure:"REDIS_DB"`

	Port string `mapstructure:"PORT"`

	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`

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
