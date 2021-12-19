package app

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB  *DBConfig
	DSN string
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Charset  string
	Database string
}

func initConfig() *Config {
	loadEnv()
	conf := &Config{DB: &DBConfig{
		Dialect:  os.Getenv("DB_CONNECTION"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Charset:  os.Getenv("DB_CHARSET"),
		Database: os.Getenv("DB_DATABASE"),
	},
	}
	return conf
}

func (c *Config) getDSN() string {
	return c.DB.Username + ":" + c.DB.Password + "@tcp(" + c.DB.Host + ")/" + c.DB.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
