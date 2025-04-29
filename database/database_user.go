package database

func (d *Database) CreateUser(username string) (string, error) {
	stmt, err := d.db.PrepareContext(
		d.Context,
		`INSERT INTO my_user (username) VALUES ($1) RETURNING user_id`,
	)

	if err != nil {
		return "", err
	}

	defer stmt.Close()

	var id string

	err = stmt.QueryRowContext(d.Context, username).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil

}

func (d *Database) GetAllUsers() ([]User, error) {
	stmt, err := d.db.PrepareContext(
		d.Context,
		`SELCT user_id, username FROM my_user`,
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

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (d *Database) GetUserById(id string) (*User, error) {
	stmt, err := d.db.PrepareContext(
		d.Context,
		`SELECT user_id, username FROM my_user 
		WHERE user_id = $1`,
	)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result := stmt.QueryRowContext(d.Context, id)

	user := &User{}

	if err := result.Scan(&user.Id, &user.Name); err != nil {
		return nil, err
	}

	return user, nil
}
