package service

import (
	"fmt"
	"forum/internal/model"
	"net/http"
)

type RegisterUser interface {
	CreateUser(user *model.User) error
	GetUser(user *model.User) (*model.User, error)
}

type RegisterUserService struct {
	repo RegisterUser
}

func NewRegisterUserService(repo RegisterUser) *RegisterUserService {
	return &RegisterUserService{
		repo: repo,
	}
}

func (rus *RegisterUserService) RegisterUser(user *model.User) (int, error) {
	userFromDB, err := rus.repo.GetUser(user)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("internal Server Error")
	}
	if userFromDB.Email == user.Email || userFromDB.Username == user.Username {
		return http.StatusBadRequest, fmt.Errorf("user with this email or username already exists")
	}
	if len(user.Email) < 6 || len(user.Email) > 50 || len(user.Username) < 6 || len(user.Username) > 50 {
		return http.StatusBadRequest, fmt.Errorf("bad request")
	}
	return http.StatusOK, rus.repo.CreateUser(user)
}
