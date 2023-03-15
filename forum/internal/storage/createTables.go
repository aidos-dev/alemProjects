package storage

import (
	"database/sql"
	"fmt"
)

func CreateTables(db *sql.DB) error {
	usersTable := `
	CREATE TABLE IF NOT EXISTS user (
	    user_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	    email TEXT,
		username TEXT,
	    password TEXT
	);`
	postTable := `
	CREATE TABLE IF NOT EXISTS post (
		post_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		heading TEXT,
		text TEXT,
		user_id INTEGER REFERENCES user(user_id)
	);`
	commentTable := `
	CREATE TABLE IF NOT EXISTS comment (
		comment_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		text TEXT,
		post_id INTEGER REFERENCES post(post_id),
		user_id INTEGER REFERENCES user(user_id)
	);`
	postLikeTable := `
	CREATE TABLE IF NOT EXISTS post_like (
		post_like_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		user_id INTEGER REFERENCES user(user_id),
		post_id INTEGER REFERENCES post(post_id),
		positive BOOLEAN
	);`
	commentLikeTable := `
	CREATE TABLE IF NOT EXISTS comment_like (
		comment_like_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		user_id INTEGER REFERENCES user(user_id),
		comment_id INTEGER REFERENCES comment(comment_id),
		positive BOOLEAN
	);`
	categoryTable := `
		CREATE TABLE IF NOT EXISTS category (
			category_id INTEGER PRIMARY KEY NOT NULL,
			category_name TEXT
		);`
	postCategoryTable := `
	CREATE TABLE IF NOT EXISTS post_category (
		post_category_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		post_id INTEGER,
		category_id INTEGER
	);`
	sessionTable := `
	CREATE TABLE IF NOT EXISTS session (
		session_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		user_id INTEGER REFERENCES user(user_id),
		cookie TEXT
	);`
	allTables := []string{usersTable, postTable, commentTable, postLikeTable, commentLikeTable, categoryTable, postCategoryTable, sessionTable}
	for _, table := range allTables {
		query, err := db.Prepare(table)
		if err != nil {
			return err
		}
		_, err = query.Exec()
		if err != nil {
			return err
		}
	}
	var name string
	err := db.QueryRow("SELECT category_name FROM category").Scan(&name)
	if err != nil {
		categories := []string{"Discussions", "Questions", "Ideas", "Articles", "Events", "Issues"}
		for i, category := range categories {
			query, err := db.Prepare(`INSERT INTO category(category_id, category_name) VALUES(?, ?);`)
			if err != nil {
				return err
			}
			_, err = query.Exec(i+1, category)
			if err != nil {
				return err
			}
		}
	}
	fmt.Println("Tables created successfully!")
	return nil
}
