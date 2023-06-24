package domain

import (
	"database/sql"
	"server/models"

	"github.com/labstack/echo/v4"
)

func RetAllPosts(db *sql.DB) ([]*models.POSTS, error) {
	rows, err := db.Query(`
	SELECT * FROM posts 
	ORDER BY updated_at ASC
	LIMIT 10000
	OFFSET 4
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := make([]*models.POSTS, 0)

	for rows.Next() {
		var post models.POSTS

		err := rows.Scan(&post.Id, &post.Title, &post.Desc, &post.Img, &post.Updated_At)

		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	return posts, nil
}

func RetSinglePost(ctx echo.Context, db *sql.DB, id int) (*models.POSTS, error) {
	row := db.QueryRow(`
	SELECT 
	id, 
	description, 
	title, 
	COALESCE(img, '') AS img, 
	updated_at 
	FROM posts WHERE id = ?`, id)

	var singlePost models.POSTS

	err := row.Scan(&singlePost.Id, &singlePost.Desc, &singlePost.Img, &singlePost.Title, &singlePost.Updated_At)

	if err != nil {
		return nil, err
	}

	return &singlePost, nil
}

func RetRecentPosts(ctx echo.Context, db *sql.DB) ([]*models.POSTS, error) {
	rows, err := db.Query(`
	SELECT id,
	description, title,
	COALESCE(img, '') AS img, 
	updated_at 
	FROM posts
	ORDER BY updated_at ASC
	LIMIT 4
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	recentPosts := make([]*models.POSTS, 0)

	for rows.Next() {
		var recentPost models.POSTS

		err := rows.Scan(&recentPost.Id, &recentPost.Desc, &recentPost.Img, &recentPost.Title, &recentPost.Updated_At)

		if err != nil {
			return nil, err
		}

		recentPosts = append(recentPosts, &recentPost)
	}

	return recentPosts, nil
}
