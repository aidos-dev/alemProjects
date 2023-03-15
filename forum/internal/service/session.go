package service

import (
	"fmt"

	"forum/internal/model"

	"github.com/google/uuid"
)

type SessionCreator interface {
	CreateSession(cookie string, userId uint) error
	DeleteSessionByUserId(userId uint) error
}

type SessionCreateService struct {
	repo SessionCreator
}

func NewSessionCreateService(repo SessionCreator) *SessionCreateService {
	return &SessionCreateService{
		repo: repo,
	}
}

func (scs *SessionCreateService) SessionCreate(user *model.User) (cookie string, err error) {
	cookieString := uuid.New().String()
	err = scs.repo.DeleteSessionByUserId(user.ID)
	if err != nil {
		return "", err
	}
	err = scs.repo.CreateSession(cookieString, user.ID)
	if err != nil {
		return "", err
	}
	return cookieString, nil
}

type SessionChecker interface {
	RetrieveSession(cookie string) (string, error)
	RetrieveUserBySession(cookie string) (*model.User, error)
	DeleteSessionByUserId(userId uint) error
}

type SessionCheckService struct {
	repo SessionChecker
}

func NewSessionCheckService(repo SessionChecker) *SessionCheckService {
	return &SessionCheckService{
		repo: repo,
	}
}

func (scs *SessionCheckService) SessionCheck(cookie string) (bool, error) {
	cookieFromDb, err := scs.repo.RetrieveSession(cookie)
	if err != nil {
		return false, err
	}
	if cookieFromDb != cookie {
		return false, fmt.Errorf("DB does not have this cookie")
	}
	return true, nil
}

func (scs *SessionCheckService) UserBySession(cookie string) (*model.User, error) {
	return scs.repo.RetrieveUserBySession(cookie)
}

func (scs *SessionCheckService) DeleteSession(userId uint) error {
	return scs.repo.DeleteSessionByUserId(userId)
}
