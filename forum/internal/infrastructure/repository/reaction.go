package repository

import (
	"database/sql"
	"fmt"
)

type reactionRepository struct {
	db *sql.DB
}

func NewReactionRepository(db *sql.DB) *reactionRepository {
	return &reactionRepository{
		db: db,
	}
}

func (rr *reactionRepository) CheckReact(postOrComment string, userId uint, postId uint, positive bool) error {
	records := fmt.Sprintf("SELECT %s_like_id FROM %s_like WHERE user_id = ? AND %s_id = ? AND positive = ?", postOrComment, postOrComment, postOrComment)
	query, err := rr.db.Prepare(records)
	if err != nil {
		return err
	}
	defer query.Close()
	var id uint
	err = query.QueryRow(userId, postId, positive).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (rr *reactionRepository) Delete(postOrComment string, userId uint, postId uint) error {
	records := fmt.Sprintf("DELETE FROM %s_like WHERE user_id = ? AND %s_id = ?", postOrComment, postOrComment)
	query, err := rr.db.Prepare(records)
	if err != nil {
		return err
	}
	defer query.Close()
	_, err = query.Exec(userId, postId)
	if err != nil {
		return err
	}
	return nil
}

func (rr *reactionRepository) React(postOrComment string, userId uint, postId uint, positive bool) (uint, error) {
	records := fmt.Sprintf(`REPLACE INTO %s_like(%s_like_id, user_id, %s_id, positive)
	VALUES((SELECT %s_like_id FROM %s_like WHERE user_id = ? AND %s_id = ?), ?, ?, ?);`, postOrComment, postOrComment, postOrComment, postOrComment, postOrComment, postOrComment)
	query, err := rr.db.Prepare(records)
	if err != nil {
		return 0, err
	}
	defer query.Close()
	result, err := query.Exec(userId, postId, userId, postId, positive)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return uint(id), err
}
