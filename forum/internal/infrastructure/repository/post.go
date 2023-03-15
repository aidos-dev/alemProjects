package repository

import (
	"database/sql"
	"fmt"
	"forum/internal/model"
	"strings"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (pr *postRepository) GetAllPosts() ([]model.PostRepresentation, error) {
	records := `SELECT t4.post_id,
						t4.text,
						t4.username,
						t4.heading,
						t4.comments,
						coalesce(t5.likes,0),
						coalesce(t5.dislikes,0)
				FROM (SELECT t1.post_id AS post_id, 
									t1.text as text, 
									t1.username as username,
									t1.heading as heading, 
									coalesce(t2.amount_comments, 0) AS comments
							FROM 
							(SELECT post_id, text, username, heading 
							FROM post INNER JOIN user
							ON post.user_id = user.user_id) AS t1
								LEFT JOIN 
							(SELECT post_id, COUNT(comment_id) AS amount_comments
							FROM comment
							GROUP BY post_id) AS t2 
								ON t1.post_id = t2.post_id) AS t4
								LEFT JOIN (SELECT post_id,
				SUM(CASE WHEN positive = true THEN 1 ELSE 0 END) AS likes,
				SUM(CASE WHEN positive = false THEN 1 ELSE 0 END) AS dislikes
				FROM post_like
				GROUP BY post_id) AS t5 ON t4.post_id = t5.post_id
				ORDER BY t4.post_id DESC`
	rows, err := pr.db.Query(records)
	if err != nil {
		return nil, err
	}
	records1 := `SELECT 
					category_name,
					post_id
				FROM post_category 
				LEFT JOIN category ON category.category_id = post_category.category_id`
	rows1, err := pr.db.Query(records1)
	if err != nil {
		return nil, err
	}
	type tempStruct1 struct {
		category_name string
		post_id       uint
	}
	var tempStruct tempStruct1
	var tempStructs []tempStruct1
	for rows1.Next() {
		rows1.Scan(&tempStruct.category_name, &tempStruct.post_id)
		tempStructs = append(tempStructs, tempStruct)
	}

	var allPosts []model.PostRepresentation
	for rows.Next() {
		var tempPost model.PostRepresentation
		rows.Scan(&tempPost.PostId, &tempPost.Text, &tempPost.Username, &tempPost.Heading, &tempPost.AmountComments, &tempPost.AmountLikes, &tempPost.AmountDisLikes)
		for _, temp := range tempStructs {
			if temp.post_id == *&tempPost.PostId {
				tempPost.Categories = append(tempPost.Categories, temp.category_name)
			}
		}
		allPosts = append(allPosts, tempPost)
	}

	return allPosts, nil
}

var ErrInvalidPost = fmt.Errorf("Bad Request")

func (pr *postRepository) CreatePost(heading string, text string, userId uint) (uint, error) {
	// if len(heading) < 1 || len(heading) > 30 || len(text) < 6 || len(text) > 1500 {
	// 	return 0, ErrInvalidPost
	// }
	records := `INSERT INTO post(heading, text, user_id)
				VALUES (?,?,?)`
	query, err := pr.db.Prepare(records)
	if err != nil {
		return 0, err
	}
	result, err := query.Exec(heading, text, userId)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return uint(id), err
}

func (pr *postRepository) GetPostById(postId uint) (*model.PostRepresentation, error) {
	records := `SELECT t4.post_id,
						t4.text,
						t4.username,
						t4.heading,
						t4.comments,
						coalesce(t5.likes,0),
						coalesce(t5.dislikes,0)
					FROM (SELECT t1.post_id AS post_id, 
									t1.text as text, 
									t1.username as username,
									t1.heading as heading, 
									coalesce(t2.amount_comments, 0) AS comments
							FROM 
							(SELECT post_id, text, username, heading 
							FROM post INNER JOIN user
							ON post.user_id = user.user_id) AS t1
								LEFT JOIN 
							(SELECT post_id, COUNT(comment_id) AS amount_comments
							FROM comment
							GROUP BY post_id) AS t2 
								ON t1.post_id = t2.post_id) AS t4
								LEFT JOIN (SELECT post_id,
					SUM(CASE WHEN positive = true THEN 1 ELSE 0 END) AS likes,
					SUM(CASE WHEN positive = false THEN 1 ELSE 0 END) AS dislikes
					FROM post_like
					GROUP BY post_id) AS t5 ON t4.post_id = t5.post_id
					WHERE t4.post_id = ?`
	rows := pr.db.QueryRow(records, postId)
	records1 := `SELECT 
					category_name,
					post_id
				FROM post_category 
				LEFT JOIN category ON category.category_id = post_category.category_id`
	rows1, err := pr.db.Query(records1)
	if err != nil {
		return nil, err
	}
	type tempStruct1 struct {
		category_name string
		post_id       uint
	}
	var tempStruct tempStruct1
	var tempStructs []tempStruct1
	for rows1.Next() {
		rows1.Scan(&tempStruct.category_name, &tempStruct.post_id)
		tempStructs = append(tempStructs, tempStruct)
	}
	var tempPost model.PostRepresentation
	err = rows.Scan(&tempPost.PostId, &tempPost.Text, &tempPost.Username, &tempPost.Heading, &tempPost.AmountComments, &tempPost.AmountLikes, &tempPost.AmountDisLikes)
	if err != nil {
		return nil, err
	}
	for _, temp := range tempStructs {
		if temp.post_id == *&tempPost.PostId {
			tempPost.Categories = append(tempPost.Categories, temp.category_name)
		}
	}
	return &tempPost, nil
}

func (pr *postRepository) AddCategoryToPost(categoryId uint, postId uint) (uint, error) {
	// records1 := `REPLACE INTO post_category(post_category_id, post_id, category_id)
	// VALUES((SELECT post_category_id FROM post_category WHERE category_id = ? AND post_id = ?), ?, ?);`
	records := `INSERT INTO post_category(category_id, post_id)
				VALUES (?,?)`
	query, err := pr.db.Prepare(records)
	if err != nil {
		return 0, err
	}
	// result, err := query.Exec(categoryId, postId, postId, categoryId)
	result, err := query.Exec(categoryId, postId)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return uint(id), err
}

func (pr *postRepository) FilterAllPosts(filterBy string) ([]model.PostRepresentation, error) { // 16,1 провалились сюда. запрос который показывает посты по нужному фильтру
	records := `SELECT t4.post_id,
				t4.text,
				t4.username,
				t4.heading,
				t4.comments,
				coalesce(t5.likes,0) as likes,
				coalesce(t5.dislikes,0) as dislikes,
                coalesce(likes - dislikes, 0) as reaction
			FROM (SELECT t1.post_id AS post_id, 
							t1.text as text, 
							t1.username as username,
							t1.heading as heading, 
							coalesce(t2.amount_comments, 0) AS comments
					FROM 
					(SELECT post_id, text, username, heading 
					FROM post INNER JOIN user
					ON post.user_id = user.user_id) AS t1
						LEFT JOIN 
					(SELECT post_id, COUNT(comment_id) AS amount_comments
					FROM comment
					GROUP BY post_id) AS t2 
						ON t1.post_id = t2.post_id) AS t4
						LEFT JOIN (SELECT post_id,
			SUM(CASE WHEN positive = true THEN 1 ELSE 0 END) AS likes,
			SUM(CASE WHEN positive = false THEN 1 ELSE 0 END) AS dislikes
			FROM post_like
			GROUP BY post_id) AS t5 ON t4.post_id = t5.post_id`
	records1 := `SELECT 
					category_name,
					post_id
				FROM post_category 
				LEFT JOIN category ON category.category_id = post_category.category_id`
	flag := false
	switch {
	case filterBy == "recent":
		records += ` ORDER BY t4.post_id DESC`
	case filterBy == "oldest":
		records += ` ORDER BY t4.post_id ASC`
	case filterBy == "most_disliked":
		records += ` ORDER BY reaction ASC `
	case filterBy == "most_liked":
		records += ` ORDER BY reaction DESC`
	default:
		flag = true
		records += ` ORDER BY t4.post_id DESC`
		records1 += ` WHERE category_name = ?`
	}
	rows, err := pr.db.Query(records)
	if err != nil {
		return nil, err
	}
	rows1, err := pr.db.Query(records1, strings.Title(filterBy))
	if err != nil {
		return nil, err
	}
	type tempStruct1 struct {
		category_name string
		post_id       uint
	}
	var tempStruct tempStruct1
	var tempStructs []tempStruct1
	for rows1.Next() {
		rows1.Scan(&tempStruct.category_name, &tempStruct.post_id)
		tempStructs = append(tempStructs, tempStruct)
	}

	var allPosts []model.PostRepresentation // 16,1 запрос сверху выполнился, а здесь сборка идет
	for rows.Next() {
		var temp1 int
		var tempPost model.PostRepresentation
		rows.Scan(&tempPost.PostId, &tempPost.Text, &tempPost.Username, &tempPost.Heading, &tempPost.AmountComments, &tempPost.AmountLikes, &tempPost.AmountDisLikes, &temp1)
		for _, temp := range tempStructs {
			if temp.post_id == *&tempPost.PostId {
				tempPost.Categories = append(tempPost.Categories, temp.category_name)
				if flag {
					allPosts = append(allPosts, tempPost)
				}
			}
		}
		if !flag {
			allPosts = append(allPosts, tempPost)
		}
	}
	return allPosts, nil
}

func (pr *postRepository) PersonalFilter(filterBy string, userId uint) ([]model.PostRepresentation, error) {
	records := ""
	switch {
	case filterBy == "i_created":
		records = `SELECT t4.post_id,
							t4.text,
							t4.username,
							t4.heading,
							t4.comments,
							coalesce(t5.likes,0),
							coalesce(t5.dislikes,0),
							t5.user_id
					FROM (SELECT t1.post_id AS post_id, 
										t1.text as text, 
										t1.username as username,
										t1.heading as heading, 
										coalesce(t2.amount_comments, 0) AS comments
								FROM 
								(SELECT post_id, text, username, heading 
								FROM post INNER JOIN user
								ON post.user_id = user.user_id
								WHERE post.user_id = ?) AS t1
									LEFT JOIN 
								(SELECT post_id, COUNT(comment_id) AS amount_comments
								FROM comment
								GROUP BY post_id) AS t2 
									ON t1.post_id = t2.post_id) AS t4
									LEFT JOIN (SELECT post_id, user_id,
					SUM(CASE WHEN positive = true THEN 1 ELSE 0 END) AS likes,
					SUM(CASE WHEN positive = false THEN 1 ELSE 0 END) AS dislikes
					FROM post_like
					GROUP BY post_id) AS t5 ON t4.post_id = t5.post_id
					ORDER BY t4.post_id DESC`
	case filterBy == "i_liked":
		records = `SELECT t4.post_id,
							t4.text,
							t4.username,
							t4.heading,
							t4.comments,
							coalesce(t5.likes,0),
							coalesce(t5.dislikes,0),
							t5.user_id
					FROM (SELECT t1.post_id AS post_id, 
										t1.text as text, 
										t1.username as username,
										t1.heading as heading, 
										coalesce(t2.amount_comments, 0) AS comments
								FROM 
								(SELECT post_id, text, username, heading 
								FROM post INNER JOIN user
								ON post.user_id = user.user_id) AS t1
									LEFT JOIN 
								(SELECT post_id, COUNT(comment_id) AS amount_comments
								FROM comment
								GROUP BY post_id) AS t2 
									ON t1.post_id = t2.post_id) AS t4
									LEFT JOIN (SELECT post_id, user_id,
					SUM(CASE WHEN positive = true THEN 1 ELSE 0 END) AS likes,
					SUM(CASE WHEN positive = false THEN 1 ELSE 0 END) AS dislikes
					FROM post_like
					WHERE user_id = ? AND positive = true
					GROUP BY post_id) AS t5 ON t4.post_id = t5.post_id
					WHERE user_id NOT NULL
					ORDER BY t4.post_id DESC`
	}
	rows, err := pr.db.Query(records, userId)
	if err != nil {
		return nil, err
	}
	records1 := `SELECT 
					category_name,
					post_id
				FROM post_category 
				LEFT JOIN category ON category.category_id = post_category.category_id`
	rows1, err := pr.db.Query(records1)
	if err != nil {
		return nil, err
	}
	type tempStruct1 struct {
		category_name string
		post_id       uint
	}
	var tempStruct tempStruct1
	var tempStructs []tempStruct1
	for rows1.Next() {
		rows1.Scan(&tempStruct.category_name, &tempStruct.post_id)
		tempStructs = append(tempStructs, tempStruct)
	}

	var allPosts []model.PostRepresentation
	for rows.Next() {
		var tempPost model.PostRepresentation
		var temp uint
		rows.Scan(&tempPost.PostId, &tempPost.Text, &tempPost.Username, &tempPost.Heading, &tempPost.AmountComments, &tempPost.AmountLikes, &tempPost.AmountDisLikes, &temp)
		for _, temp := range tempStructs {
			if temp.post_id == *&tempPost.PostId {
				tempPost.Categories = append(tempPost.Categories, temp.category_name)
			}
		}
		allPosts = append(allPosts, tempPost)
	}
	return allPosts, nil
}
