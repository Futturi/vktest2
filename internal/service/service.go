package service

import (
	"vktest2/internal/models"
	"vktest2/internal/repository"
)

type Service struct {
	AuthService
	AnnouncementService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{AuthService: NewAuth_Service(repo.AuthRepo), AnnouncementService: NewAnnouncement_Service(repo.AnnouncementRepo)}
}

type AuthService interface {
	SignUp(user models.User) (int, error)
	SignIn(user models.User) (string, error)
	SetHeader(header string) (int, error)
}

type AnnouncementService interface {
	CreateAnn(ann models.Annunc, idUser string) (int, error)
}
