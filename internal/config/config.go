package config

import (
	"fmt"
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
	
	var (psqlPassword string
		psqlUser string
		psqlDB string
		psqlHost string
		psqlPort string)
	
	psqlPassword = os.Getenv("POSTGRES_PASSWORD")
	psqlUser = os.Getenv("POSTGRES_USER")
	psqlDB = os.Getenv("POSTGRES_DB")
	psqlHost = os.Getenv("POSTGRES_HOST")
	psqlPort = os.Getenv("POSTGRES_PORT")
	
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	psqlHost, psqlUser, psqlPassword, psqlDB, psqlPort)
	
	return &Config{
		dbhost: connStr,
		port: int(port),
	}, nil
}