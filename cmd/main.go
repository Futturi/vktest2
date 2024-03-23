package main

import (
	"context"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"vktest2/internal/handler"
	"vktest2/internal/repository"
	"vktest2/internal/server"
	"vktest2/internal/service"
	"vktest2/pkg"
)

// @title Announcement app
// @version 1.0
// @description API Server 4 Announcement Application

// @host localhost:8082
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logg := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logg)
	err := InitViper()
	if err != nil {
		slog.Error("error while reading config", slog.Any("error", err))
	}
	pcfg := pkg.PConfig{
		Username: viper.GetString("db.username"),
		Port:     viper.GetString("db.port"),
		Host:     viper.GetString("db.host"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	db, err := pkg.InitPostgres(pcfg)
	if err != nil {
		slog.Error("error while creating db", slog.Any("error", err))
	}
	repo := repository.NewRepository(db)
	servi := service.NewService(repo)
	handl := handler.NewHandler(servi)
	serv := new(server.Server)
	go func() {
		if err := serv.InitServer(viper.GetString("port"), handl.InitHandlers()); err != nil {
			slog.Error("error while starting server", slog.Any("error", err))
		}
	}()
	logg.Info("statring app in port: ", slog.String("port", viper.GetString("port")))

	if err := pkg.Migrat(viper.GetString("db.host")); err != nil {
		slog.Error("error with migratedb", slog.Any("error", err))
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logg.Info("shutdown server", slog.String("port", viper.GetString("port")))
	if err := serv.StopServer(context.Background()); err != nil {
		slog.Error("error while stopping server", slog.Any("error", err))
		os.Exit(1)
	}
	if err := pkg.ShutDown(db); err != nil {
		slog.Error("error while closing db", slog.Any("error", err))
		os.Exit(1)
	}
}

func InitViper() error {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
