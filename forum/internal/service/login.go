package service

import (
	"fmt"
	"forum/internal/model"
	"net/http"
)

type LoginUser interface {
	GetUserByUsernameAndPassword(user *model.User) (*model.User, error)
}

type LoginUserService struct {
	repo LoginUser
}

func NewLoginUserService(repo LoginUser) *LoginUserService {
	return &LoginUserService{
		repo: repo,
	}
}

func (lus *LoginUserService) LoginUser(user *model.User) (uint, bool, int, error) {
	if len(user.Username) < 6 || len(user.Username) > 50 || len(user.Password) < 4 || len(user.Password) > 120 {
		return 0, false, http.StatusBadRequest, fmt.Errorf("bad request")
	}
	userFromDB, err := lus.repo.GetUserByUsernameAndPassword(user)
	if err != nil {
		return 0, false, http.StatusInternalServerError, fmt.Errorf("internal server error")
	}
	if user.Username == userFromDB.Username && user.Password == userFromDB.Password {
		return userFromDB.ID, true, http.StatusOK, nil
	}
	return 0, false, http.StatusUnauthorized, fmt.Errorf("username or password is not correct")
}
