package database

func (d *Database) CreateComment(userId string, title string, content string) (string, error) {
	stmt, err := d.db.PrepareContext(
		d.Context,
		`INSERT INTO comment (title, content, user_id) 
		VALUES ($1, $2, $3) 
		RETURNING comment_id`,
	)

	if err != nil {
		return "", err
	}

	defer stmt.Close()

	var id string

	err = stmt.QueryRowContext(d.Context, title, content, userId).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (d *Database) GetAllCommentsWithUser() ([]*User, error) {
	stmt, err := d.db.PrepareContext(
		d.Context,
		`SELECT
			u.user_id, 
			u.username, 
			COALESCE(CAST(c.comment_id AS VARCHAR), '') AS comment_id,
			COALESCE(c.title, '') AS title, 
			COALESCE(c.content, '') AS content
		FROM my_user u
		LEFT JOIN comment c ON u.user_id = c.user_id
		ORDER BY u.user_id ASC;`,
	)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(d.Context)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*User

	// Check if user is already in the map
	// If not, create a new user and add it to the map

	userMap := make(map[string]*User)

	for rows.Next() {
		user := &User{}
		comment := &Comment{}

		if err := rows.Scan(
			&user.UserId,
			&user.Name,
			&comment.CommentId,
			&comment.Title,
			&comment.Content,
		); err != nil {
			return nil, err
		}

		_, exist := userMap[user.UserId]

		if !exist {
			userMap[user.UserId] = user
			users = append(users, user)
		}

		if comment.CommentId != "" {
			userMap[user.UserId].Comments = append(userMap[user.UserId].Comments, *comment)
		}
	}

	return users, nil
}

func (d *Database) GetCommentByCommentId(commentId string) (*Comment, error) {
	stmt, err := d.db.PrepareContext(
		d.Context,
		`SELECT comment_id, title, content, user_id 
		FROM comment WHERE comment_id = $1`,
	)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result := stmt.QueryRowContext(d.Context, commentId)

	comment := &Comment{}

	err = result.Scan(&comment.CommentId, &comment.Title, &comment.Content, &comment.UserId)

	if err != nil {
		return nil, err
	}

	return comment, nil
}
