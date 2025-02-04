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
	
	db, err := database.InitDB(config)
	defer db.Close()

	
	if errors.Is(err, sql.ErrNoRows) {
	}
	
	if err:= server.Run(db, server.Config {Port: config.Port()}); err != nil {
		return fmt.Errorf("server run: %w", err)
	}
	
	return nil	
}