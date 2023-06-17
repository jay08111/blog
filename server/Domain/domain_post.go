package domain

import "server/models"

func PostsSelectAll() ([]*models.POSTS, error) {
	db := MainDB

	q := `SELECT * FROM posts`

	rows, err := db.Query(q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []*models.POSTS

	for rows.Next() {
		post := &models.POSTS{}
		err := rows.Scan()

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
