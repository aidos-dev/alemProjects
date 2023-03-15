package repository

import (
	"database/sql"

	"forum/internal/model"
)

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *commentRepository {
	return &commentRepository{
		db: db,
	}
}

func (cr *commentRepository) CreateComment(userId, postId uint, text string) (uint, error) {
	// empty comment text check
	// if len(text) < 1 || len(text) > 1500 || userId == 0 || postId == 0 || userId < 1 || postId < 1 {
	// 	return 0, fmt.Errorf("bad request")
	// }

	records := `INSERT INTO comment(text, post_id, user_id) 
				VALUES (?,?,?)`
	query, err := cr.db.Prepare(records) // 13 описывается и подготавивается запрос который инсертит значения в таблицу комментариев
	if err != nil {
		return 0, err
	}
	result, err := query.Exec(text, postId, userId)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return uint(id), err
}

func (cr *commentRepository) GetAllCommentsByPostId(postId uint) ([]model.CommentRepresentation, error) {
	records := `SELECT 			
						t4.comment_id,
						t4.text,
						t4.user_id,
						t4.username,
						t4.post_id,
						COALESCE(t5.likes, 0) as comment_likes,
						COALESCE(t5.dislikes, 0) as comment_dislikes
						FROM
						(SELECT t1.comment_id as comment_id,
						t1.text as text,
						t1.user_id as user_id,
						t1.post_id as post_id,
						t2.username as username
						FROM
								(SELECT comment_id, text, post_id, user_id
								FROM comment) as t1
						LEFT JOIN (SELECT user_id, username
								FROM user) AS t2 ON t1.user_id = t2.user_id) AS t4
						LEFT JOIN (SELECT comment_id,
						SUM(CASE WHEN positive = true THEN 1 ELSE 0 END) AS likes,
						SUM(CASE WHEN positive = false THEN 1 ELSE 0 END) AS dislikes
					FROM comment_like
					GROUP BY comment_id) AS t5 ON t4.comment_id = t5.comment_id
						WHERE t4.post_id = ?`
	rows, err := cr.db.Query(records, postId)
	if err != nil {
		return nil, err
	}
	var allComments []model.CommentRepresentation
	for rows.Next() {
		var tempComment model.CommentRepresentation
		err := rows.Scan(&tempComment.CommentId, &tempComment.Text, &tempComment.UserId, &tempComment.Username, &tempComment.PostId, &tempComment.AmountLikes, &tempComment.AmountDisLikes)
		if err != nil {
			return nil, err
		}
		allComments = append(allComments, tempComment)
	}
	return allComments, nil
}
