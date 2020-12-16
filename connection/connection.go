package connection

import (
	"database/sql"
	"fmt"

	// postgres driver
	_ "github.com/lib/pq"
	"github.com/udaya2899/go-gin-starter/configuration"
)

// NewConnection creates a new connection from the given configs
func NewConnection(dbConfig configuration.DatabaseConfiguration) (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("Failed to open connection to Database, err: %v", err)
	}

	return db, nil

}
