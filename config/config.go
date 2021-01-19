package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	Token string
	ListenAddr string
	DBHost string
	DBPort string
	DBUsername string
	DBDatabase string
	DBPassword string
}

var cfg = new(Config)

func init() {
	_, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func LoadConfig() (*Config, error) {
	newCfg := new(Config)

	newCfg.Token = os.Getenv("SENSITIVE_TOKEN")
	if newCfg.Token == "" {
		return nil, fmt.Errorf("SENSITIVE_TOKEN cannot be null")
	}

	newCfg.ListenAddr = os.Getenv("SENSITIVE_LISTEN_ADDR")
	if newCfg.ListenAddr == "" {
		return nil, fmt.Errorf("SENSITIVE_LISTEN_ADDR cannot be null")
	}

	newCfg.DBHost = os.Getenv("DB_SECURITY_READ_HOST")
	if newCfg.DBHost == "" {
		return nil, fmt.Errorf("DB_SECURITY_READ_HOST cannot be null")
	}

	newCfg.DBPort = os.Getenv("DB_SECURITY_READ_PORT")
	if newCfg.DBPort == "" {
		return nil, fmt.Errorf("DB_SECURITY_READ_PORT cannot be null")
	}

	newCfg.DBDatabase = os.Getenv("DB_SECURITY_READ_DATABASE")
	if newCfg.DBDatabase == "" {
		return nil, fmt.Errorf("DB_SECURITY_READ_DATABASE cannot be null")
	}

	newCfg.DBUsername = os.Getenv("DB_SECURITY_READ_USERNAME")
	if newCfg.DBUsername == "" {
		return nil, fmt.Errorf("DB_SECURITY_READ_USERNAME cannot be null")
	}

	newCfg.DBPassword = os.Getenv("DB_SECURITY_READ_PASSWORD")
	if newCfg.DBPassword == "" {
		return nil, fmt.Errorf("DB_SECURITY_READ_PASSWORD cannot be null")
	}

	cfg = newCfg

	return cfg, nil
}

func GetConfig() *Config {
	return cfg
}