package service

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
	"vktest2/internal/models"
	"vktest2/internal/repository"
	"vktest2/internal/utils"
)

const (
	salt = "ijniljew9753nvnalqo3jnbskqo2u57dnveks,cmvjrkwlamckgfk3la,mcmvkrlq;s,cmngjeuqjajdurjfc"
)

type Auth_Service struct {
	repo repository.AuthRepo
}

func NewAuth_Service(repo repository.AuthRepo) *Auth_Service {
	return &Auth_Service{repo: repo}
}

func (a *Auth_Service) SignUp(user models.User) (int, error) {
	if user.Password == "" || len(user.Password) > 70 {
		return 0, errors.New("incorrect password")
	}
	if user.Username == "" || len(user.Username) > 40 {
		return 0, errors.New("incorrect username")
	}
	user = models.User{Username: user.Username, Password: utils.HashPass(user.Password)}
	return a.repo.SignUp(user)
}

type Claims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

func (a *Auth_Service) SignIn(user models.User) (string, error) {
	user = models.User{Username: user.Username, Password: utils.HashPass(user.Password)}
	id, err := a.repo.SignIn(user)
	if err != nil {
		return "", err
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{id, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		IssuedAt:  time.Now().Unix(),
	}})
	return t.SignedString([]byte(salt))
}

func (r *Auth_Service) SetHeader(accesstoken string) (int, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, errors.New("invalid signing method")
		}
		return []byte(salt), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.Id, nil
}
