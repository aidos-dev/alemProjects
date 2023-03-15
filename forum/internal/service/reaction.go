package service

import (
	"database/sql"
	"errors"
)

type React interface {
	React(postOrComment string, userId uint, postId uint, positive bool) (uint, error)
	CheckReact(postOrComment string, userId uint, postId uint, positive bool) error
	Delete(postOrComment string, userId uint, postId uint) error
}

type ReactService struct {
	repo React
}

func NewReacttService(repo React) *ReactService {
	return &ReactService{
		repo: repo,
	}
}

func (rs *ReactService) React(postOrComment string, userId uint, postId uint, positive bool) (uint, error) { // 19,1 проверка есть ли у поста или коммента реакция, чтобы можно было убрать если уже есть
	err := rs.repo.CheckReact(postOrComment, userId, postId, positive)
	if err == nil { // если ошибки нет, значит реакция есть и мы убираем реакцию
		err = rs.repo.Delete(postOrComment, userId, postId)
		if err != nil {
			return 0, err
		}
		return 0, nil
	} else if !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	return rs.repo.React(postOrComment, userId, postId, positive) // 19,1 а тут изменяем на другую или ставим реакцию
}
