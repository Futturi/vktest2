package pkg

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func InitPostgres(cfg PConfig) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("postgres", fmt.Sprintf("host =%s port =%s user =%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func ShutDown(db *sqlx.DB) error {
	return db.Close()
}
