package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"vktest2/internal/models"
)

type Announcement_Postgre struct {
	db *sqlx.DB
}

func NewAnnouncement_Postgre(db *sqlx.DB) *Announcement_Postgre {
	return &Announcement_Postgre{db: db}
}

func (r *Announcement_Postgre) CreateAnn(ann models.Annunc, idUser int) (int, error) {
	var id int
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query1 := fmt.Sprintf("INSERT INTO %s(name, body, image, price) VALUES($1,$2,$3,$4) RETURNING id", annTable)
	row := tx.QueryRow(query1, ann.Name, ann.Body, ann.Image, ann.Price)
	if err = row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	query2 := fmt.Sprintf("INSERT INTO %s(user_id, announcement_id) VALUES($1, $2)", UseranTable)
	_, err = tx.Exec(query2, idUser, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
