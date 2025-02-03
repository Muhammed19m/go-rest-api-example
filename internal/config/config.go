package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)
const ConfigFileName = "config.env"



type Config struct {
	dbhost string 
	port int
}

func (c *Config) DBHost() string {
 	return c.dbhost
}

func (c *Config) Port() int {
	return c.port	
}





func LoadConfig() (*Config, error) {
	err := godotenv.Load(ConfigFileName)
	if err != nil {
		return nil, err
	}

	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
	if err != nil {
		return nil, err
	}
	

	return &Config{
		dbhost: os.Getenv("DATABASE_URL"),
		port: int(port),
	}, nil
}