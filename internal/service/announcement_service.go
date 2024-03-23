package service

import (
	"errors"
	"strconv"
	"vktest2/internal/models"
	"vktest2/internal/repository"
)

type Announcement_Service struct {
	repo repository.AnnouncementRepo
}

func NewAnnouncement_Service(repo repository.AnnouncementRepo) *Announcement_Service {
	return &Announcement_Service{repo: repo}
}

func (a *Announcement_Service) CreateAnn(ann models.Annunc, idUser string) (int, error) {
	if len(ann.Body) > 1000 {
		return 0, errors.New("too big body")
	}
	if ann.Name == "" || len(ann.Name) > 50 {
		return 0, errors.New("incorrect name")
	}
	if ann.Price < 0 {
		return 0, errors.New("incorrect price")
	}
	id, err := strconv.Atoi(idUser)
	if err != nil {
		return 0, err
	}
	return a.repo.CreateAnn(ann, id)
}
