// Package database contains a Postgres client and methods for communicating with the database.
package database

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"template-api-go/config"
)

// Client holds the database client and prepared statements.
type Client struct {
	DB                         *sqlx.DB
	GetExampleDataStatement    *sqlx.Stmt
	RecordExampleDataStatement *sqlx.Stmt
}

// Init sets up a new database client.
func (c *Client) Init(ctx context.Context, config *config.Config) error {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s%s",
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseURL,
		config.DatabasePort,
		config.DatabaseDB,
		config.DatabaseOptions,
	)

	db, err := sqlx.ConnectContext(ctx, "postgres", connString)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	db.SetMaxOpenConns(config.DatabaseMaxConnections)
	db.SetMaxIdleConns(config.DatabaseMaxIdleConnections)

	c.DB = db

	err = c.prepareGetExampleDataStmt()
	if err != nil {
		return fmt.Errorf("failed to prepare get example data statement: %w", err)
	}

	err = c.prepareRecordExampleDataStmt()
	if err != nil {
		return fmt.Errorf("failed to prepare record example data statement: %w", err)
	}

	return nil
}

// Close closes the database connection and statements.
func (c *Client) Close() error {
	err := c.GetExampleDataStatement.Close()
	if err != nil {
		return fmt.Errorf("error closing get example data statement: %w", err)
	}

	err = c.RecordExampleDataStatement.Close()
	if err != nil {
		return fmt.Errorf("error closing record example data statement: %w", err)
	}

	err = c.DB.Close()
	if err != nil {
		return fmt.Errorf("error closing database: %w", err)
	}

	return nil
}
