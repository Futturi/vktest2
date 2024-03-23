package service

import (
	"errors"
	"sort"
	"strconv"
	"time"
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
	if ann.Price < 0 || ann.Price > 1000000000 {
		return 0, errors.New("incorrect price")
	}
	id, err := strconv.Atoi(idUser)
	if err != nil {
		return 0, err
	}
	ann.Data = time.Now().Unix()
	return a.repo.CreateAnn(ann, id)
}

func (a *Announcement_Service) GetAnnsWithoutAuth(page, sor, sorTo, minP, maxP string) ([]models.AnnuncRes, error) {
	if minP == "" {
		minP = "0"
	}
	if maxP == "" {
		maxP = "999999999"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return []models.AnnuncRes{}, err
	}
	minPrice, err := strconv.Atoi(minP)
	if err != nil {
		return []models.AnnuncRes{}, err
	}
	maxPrice, err := strconv.Atoi(maxP)
	if err != nil {
		return []models.AnnuncRes{}, err
	}

	ann, err := a.repo.GetAnnsWithoutAuth(pageInt, minPrice, maxPrice)
	if err != nil {
		return []models.AnnuncRes{}, err
	}
	if sor == "date" || sor == "" {
		if sorTo == "down" || sorTo == "" {
			sort.Slice(ann, func(i, j int) bool {
				return ann[i].Data < ann[j].Data
			})
		} else if sorTo == "up" {
			sort.Slice(ann, func(i, j int) bool {
				return ann[i].Data > ann[j].Data
			})
		}
	}
	if sor == "price" {
		if sorTo == "down" || sorTo == "" {
			sort.Slice(ann, func(i, j int) bool {
				return ann[i].Price < ann[j].Price
			})
		} else if sorTo == "up" {
			sort.Slice(ann, func(i, j int) bool {
				return ann[i].Price > ann[j].Price
			})
		}
	}

	var result []models.AnnuncRes

	for _, an := range ann {
		result = append(result, models.AnnuncRes{
			Id:    an.Id,
			Name:  an.Name,
			Body:  an.Body,
			Image: an.Image,
			Price: an.Price,
			Data:  time.Unix(an.Data, 0).Format("2006/01/02"),
		})
	}

	return result, nil
}

func (a *Announcement_Service) GetAnns(page, sor, sorTo, minP, maxP, idUser string) ([]models.AnnuncARes, error) {
	if minP == "" {
		minP = "0"
	}
	if maxP == "" {
		maxP = "99999999"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return []models.AnnuncARes{}, err
	}
	minPrice, err := strconv.Atoi(minP)
	if err != nil {
		return []models.AnnuncARes{}, err
	}
	maxPrice, err := strconv.Atoi(maxP)
	if err != nil {
		return []models.AnnuncARes{}, err
	}
	id, err := strconv.Atoi(idUser)
	if err != nil {
		return []models.AnnuncARes{}, err
	}

	ann, err := a.repo.GetAnns(pageInt, minPrice, maxPrice, id)
	if err != nil {
		return []models.AnnuncARes{}, err
	}
	if sor == "date" || sor == "" {
		if sorTo == "down" || sorTo == "" {
			sort.Slice(ann, func(i, j int) bool {
				return ann[i].Data > ann[j].Data
			})
		} else if sorTo == "up" {
			sort.Slice(ann, func(i, j int) bool {
				return ann[i].Data < ann[j].Data
			})
		}
	}
	if sor == "price" {
		if sorTo == "down" || sorTo == "" {
			sort.Slice(ann, func(i, j int) bool {
				return ann[i].Price > ann[j].Price
			})
		} else if sorTo == "up" {
			sort.Slice(ann, func(i, j int) bool {
				return ann[i].Price < ann[j].Price
			})
		}
	}

	var result []models.AnnuncARes

	for _, an := range ann {
		result = append(result, models.AnnuncARes{
			Id:     an.Id,
			Name:   an.Name,
			Body:   an.Body,
			Image:  an.Image,
			Price:  an.Price,
			Data:   time.Unix(an.Data, 0).Format("2006/01/02"),
			IsYour: an.IsYour,
		})
	}

	return result, nil
}
