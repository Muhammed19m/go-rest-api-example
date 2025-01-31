package app

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"rest-api/internal/database"
	"rest-api/internal/http/server"
)

func Run( /* config Config */ ) error {
	err := database.InitDB()
	if errors.Is(err, sql.ErrNoRows) {
	}

	// defer dbhand.Close()
	if err:= server.Run(server.Config{Port: os.Getenv("PORT")}); err != nil {
		return fmt.Errorf("server run: %w", err)
	}
	
	return nil	
}