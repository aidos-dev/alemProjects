package repository

import (
	"database/sql"

	"forum/internal/model"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(user *model.User) error {
	records := `INSERT INTO user(email, username, password)
				VALUES (?,?,?)`
	query, err := ur.db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(user.Email, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUser(user *model.User) (*model.User, error) {
	records := `SELECT email, username
				FROM user
				WHERE email = ? OR username = ?
				LIMIT 1`
	query, err := ur.db.Prepare(records)
	if err != nil {
		return nil, err
	}
	rows, err := query.Query(user.Email, user.Username)
	if err != nil {
		return nil, err
	}
	var tempUser model.User
	for rows.Next() {
		rows.Scan(&tempUser.Email, &tempUser.Username)
	}
	return &tempUser, err
}

func (ur *userRepository) GetUserByUsernameAndPassword(user *model.User) (*model.User, error) {
	records := `SELECT username, password, user_id
				FROM user
				WHERE username = ? AND password = ?
				`
	query, err := ur.db.Prepare(records)
	if err != nil {
		return nil, err
	}
	rows, err := query.Query(user.Username, user.Password)
	if err != nil {
		return nil, err
	}
	var tempUser model.User
	for rows.Next() {
		rows.Scan(&tempUser.Username, &tempUser.Password, &tempUser.ID)
	}
	return &tempUser, err
}
