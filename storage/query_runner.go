package storage

import (
	"database/sql"
	"log"
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

// ExecuteReadQuery ...
// Executes the specified query and returns the results.
func (c *QueryRunner) ExecuteReadQuery(query string, params ...interface{}) *sql.Rows {
	// TODO: Validate database connection is open

	// Execute the query
	rows, err := c.Database.Query(query, params...)
	if err != nil {
		// TODO: Return status enum value
		log.Fatal(err)
	}

	return rows
}
