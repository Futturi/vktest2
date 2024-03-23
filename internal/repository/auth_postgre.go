package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"vktest2/internal/models"
)

type Auth_Repo struct {
	db *sqlx.DB
}

func NewAuth_Repo(db *sqlx.DB) *Auth_Repo {
	return &Auth_Repo{db: db}
}

func (r *Auth_Repo) SignUp(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(username, password_hash) VALUES($1,$2) RETURNING ID", userTable)
	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Auth_Repo) SignIn(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password_hash = $2", userTable)
	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
