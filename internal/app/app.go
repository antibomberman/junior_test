package app

import (
	"github.com/antibomberman/junior_test/internal/config"
	"github.com/antibomberman/junior_test/internal/controllers"
	"github.com/antibomberman/junior_test/internal/services"
	"github.com/gin-gonic/gin"
	"log/slog"
	"strconv"
)

type App struct {
	userSrv services.UserService
	cfg     *config.Config
	log     *slog.Logger
}

func NewApp(userSrv services.UserService, cfg *config.Config, log *slog.Logger) *App {
	return &App{
		userSrv: userSrv,
		cfg:     cfg,
		log:     log,
	}
}

func (a *App) RunServer() {
	router := gin.Default()
	userController := controllers.NewUserController(a.userSrv)

	router.GET("/user", userController.Index)
	router.GET("/user/:id", userController.Show)
	router.POST("/user", userController.Create)
	router.PUT("/user/:id", userController.Update)
	router.DELETE("/user/:id", userController.Delete)

	a.log.Info("starting server", slog.String("port", strconv.Itoa(a.cfg.Port)))
	router.Run(":" + strconv.Itoa(a.cfg.Port))
}
