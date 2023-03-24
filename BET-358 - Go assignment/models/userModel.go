package models

import (
	"database/sql"

	DB "rabietf.me/go-assignment/db"
)

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

// Method for inserting new user in database.
// Returns (userId, nil) if successful.
// Returns (0, err) if failed.
func (user User) Save() (int64, error) {
	result, err := DB.Connection.Exec("INSERT INTO Users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)

	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Method for finding is user exists in database by checking email, useful for new account creation and login.
// Returns (true, nil) and puts user data in object if user exists.
// Returns (false, nil) if user doesn't exist.
// Returns (false, err) if something went wrong.
func (user *User) Find(email string) (bool, error) {

	row := DB.Connection.QueryRow("SELECT * FROM Users WHERE email = ?", email)

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
