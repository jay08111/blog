package domain

import (
	"server/models"
)

func PostsSelectAll() ([]*models.POSTS, error) {
	db := MainDB
	rows, err := db.Query(`SELECT * FROM posts`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := make([]*models.POSTS, 0)

	for rows.Next() {
		var post models.POSTS

		err := rows.Scan(&post.Id, &post.Title, &post.Desc, &post.Img, &post.Updated_at)

		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	return posts, nil
}
