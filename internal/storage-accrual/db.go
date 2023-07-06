package storageaccrual

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	db *sql.DB
}

//go:embed migrations/*.sql
var migrationsDir embed.FS

func NewDBStorage(dsn string) (*Storage, error) {
	err := runMigrations(dsn)
	if err != nil {
		return nil, fmt.Errorf("during initializing of new db session, error occurred: %w", err)
	}

	db, err := newDBSession(dsn)
	if err != nil {
		return nil, fmt.Errorf("during initializing of new db session, error occurred: %w", err)
	}

	return &Storage{db}, nil
}

func runMigrations(dsn string) error {
	d, err := iofs.New(migrationsDir, "migrations")
	if err != nil {
		return fmt.Errorf("failed to return an iofs driver: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, dsn)
	if err != nil {
		return fmt.Errorf("failed to get a new migrate instance: %w", err)
	}
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("failed to apply migrations to the DB: %w", err)
		}
	}
	return nil
}

func newDBSession(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("attempt to establish connection failed - %w", err)
	}
	return db, nil
}
