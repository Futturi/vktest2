package main

import (
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"log/slog"
	"os"
	"vktest2/internal/handler"
	"vktest2/internal/repository"
	"vktest2/internal/server"
	"vktest2/internal/service"
	"vktest2/pkg"
)

func main() {
	logg := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logg)
	if err := gotenv.Load(); err != nil {
		slog.Error("error with godotenv", slog.Any("error", err))
	}
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
	if err := serv.InitServer(viper.GetString("port"), handl.InitHandlers()); err != nil {
		slog.Error("error while starting server", slog.Any("error", err))
	}
}

func InitViper() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
