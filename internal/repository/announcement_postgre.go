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
	query1 := fmt.Sprintf("INSERT INTO %s(name, body, image, price,data) VALUES($1,$2,$3,$4, $5) RETURNING id", annTable)
	row := tx.QueryRow(query1, ann.Name, ann.Body, ann.Image, ann.Price, ann.Data)
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

func (r *Announcement_Postgre) GetAnnsWithoutAuth(pageInt, minPrice, maxPrice int) ([]models.Annunc, error) {
	var ann []models.Annunc
	if pageInt == 1 || pageInt == 0 {
		query := fmt.Sprintf("SELECT * FROM %s WHERE price > $1 AND price < $2 LIMIT 5", annTable)
		if err := r.db.Select(&ann, query, minPrice, maxPrice); err != nil {
			return []models.Annunc{}, err
		}
	} else {
		pageInt = 5 * (pageInt - 1)
		query := fmt.Sprintf("SELECT * FROM %s WHERE price > $1 AND price < $2 LIMIT 5 offset $3", annTable)
		if err := r.db.Select(&ann, query, minPrice, maxPrice, pageInt); err != nil {
			return []models.Annunc{}, err
		}
	}
	return ann, nil
}

func (r *Announcement_Postgre) GetAnns(pageInt, minPrice, maxPrice, id int) ([]models.AnnuncA, error) {
	var ann []models.AnnuncA
	if pageInt == 1 || pageInt == 0 {
		query := fmt.Sprintf("SELECT * FROM %s WHERE price > $1 AND price < $2 LIMIT 5", annTable)
		if err := r.db.Select(&ann, query, minPrice, maxPrice); err != nil {
			return []models.AnnuncA{}, err
		}
	} else {
		pageInt = 5 * (pageInt - 1)
		query := fmt.Sprintf("SELECT * FROM %s WHERE price > $1 AND price < $2 LIMIT 5 offset $3", annTable)
		if err := r.db.Select(&ann, query, minPrice, maxPrice, pageInt); err != nil {
			return []models.AnnuncA{}, err
		}
	}

	var result []models.AnnuncA

	for _, an := range ann {
		var id1 int
		query1 := fmt.Sprintf("SELECT user_id FROM %s WHERE announcement_id = $1", UseranTable)
		row := r.db.QueryRow(query1, an.Id)
		if err := row.Scan(&id1); err != nil {
			return []models.AnnuncA{}, err
		}
		if id == id1 {
			result = append(result, models.AnnuncA{
				Id:     an.Id,
				Name:   an.Name,
				Body:   an.Body,
				Data:   an.Data,
				Image:  an.Image,
				Price:  an.Price,
				IsYour: true,
			})
		} else {
			result = append(result, models.AnnuncA{
				Id:     an.Id,
				Name:   an.Name,
				Body:   an.Body,
				Data:   an.Data,
				Image:  an.Image,
				Price:  an.Price,
				IsYour: false,
			})
		}
	}
	return result, nil
}
