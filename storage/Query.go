package storage

import (
	"database/sql"
	"log"

	// TODO: Move to main package? (PostgreSql driver)
	_ "github.com/lib/pq"
)

// QueryRunner ...
// Used to run queries against a database.
type QueryRunner struct {
	Database *sql.DB
}

// NewQueryRunner ...
// Instantiates a new QueryRunner instance.
func NewQueryRunner(database *sql.DB) *QueryRunner {
	return &QueryRunner{Database: database}
}

// ReadRows ...
// Executes the specified query and returns the results.
func (c *QueryRunner) ReadRows(query string) *sql.Rows {

	// TODO: Validate database connection is open

	// Execute the query
	rows, err := c.Database.Query(query)
	if err != nil {
		// TODO: Return status enum value
		log.Fatal(err)
	}

	return rows
}