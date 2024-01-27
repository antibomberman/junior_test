package main

import (
	"database/sql"
	"github.com/antibomberman/junior_test/internal/config"
	"github.com/antibomberman/junior_test/pkg/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log/slog"
)

func main() {
	cfg := config.Load()
	log := log.SetupLogger(cfg.Env)
	log.Info("starting service", slog.String("env", cfg.Env))

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: "migrations",
	})
	if err != nil {
		log.Error(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Error(err.Error())
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error(err.Error())

	}

	log.Info("Migrations applied successfully!")

}
