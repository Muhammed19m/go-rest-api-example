package app

import (
	"database/sql"
	"errors"
	"fmt"

	"rest-api/internal/config"
	"rest-api/internal/database"
	"rest-api/internal/http/server"
)

func Run(config *config.Config) error {
	err := database.InitDB(config)
	if errors.Is(err, sql.ErrNoRows) {
	}

	// defer dbhand.Close()
	if err:= server.Run(server.Config{Port: fmt.Sprint(config.Port())}); err != nil {
		return fmt.Errorf("server run: %w", err)
	}
	
	return nil	
}