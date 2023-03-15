package repository

import (
	"database/sql"
	"errors"
	"os"

	"forum/internal/storage"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	userRepository
	sessionRepository
	postRepository
	commentRepository
	reactionRepository
}

func NewRepository(db *sql.DB) *Repository { // 7,1 сборка всех реализаций методов баз данных
	return &Repository{
		*NewUserRepository(db),
		*NewSessionRepository(db),
		*NewPostRepository(db),
		*NewCommentRepository(db),
		*NewReactionRepository(db),
	}
}

func RunDb() (*sql.DB, error) {
	if _, err := os.Stat("database.db"); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create("database.db")
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}
	db, err := sql.Open("sqlite3", "file:database.db?_foreign_keys=on")
	if err != nil {
		return nil, err
	}
	err = storage.CreateTables(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// db, err := sql.Open("sqlite3", "file:forum.db?_foreign_keys=on") -> This will ensure that foreign key constraints are enabled for the SQLite database connection.

/* func RunDb() (*sql.DB, error) {
	if _, err := os.Stat("database.db"); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create("database.db")
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}

	db, err := sql.Open("sqlite3", "file:database.db?_foreign_keys=on")
	if err != nil {
		return nil, err
	}
	// defer db.Close() will ensure that the database connection is closed properly, even if an error occurs while creating tables or if the function is interrupted.
	defer db.Close()
	err = storage.CreateTables(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
*/
