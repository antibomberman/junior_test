package main

import (
	"database/sql"
	"github.com/antibomberman/junior_test/internal/app"
	"github.com/antibomberman/junior_test/internal/config"
	"github.com/antibomberman/junior_test/internal/repository"
	"github.com/antibomberman/junior_test/internal/services"
	"github.com/antibomberman/junior_test/pkg/log"
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

	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	myApp := app.NewApp(userService, cfg, log)
	myApp.RunServer()

}
