package repository

import (
	"github.com/jmoiron/sqlx"
	"vktest2/internal/models"
)

const (
	userTable   = "users"
	annTable    = "announcements"
	UseranTable = "users_announcements"
)

type Repository struct {
	AuthRepo
	AnnouncementRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{AuthRepo: NewAuth_Repo(db),
		AnnouncementRepo: NewAnnouncement_Postgre(db)}
}

type AuthRepo interface {
	SignUp(user models.User) (int, error)
	SignIn(user models.User) (int, error)
}

type AnnouncementRepo interface {
	CreateAnn(ann models.Annunc, idUser int) (int, error)
	GetAnnsWithoutAuth(pageInt, minPrice, maxPrice int) ([]models.Annunc, error)
	GetAnns(pageInt, minPrice, maxPrice, id int) ([]models.AnnuncA, error)
}
