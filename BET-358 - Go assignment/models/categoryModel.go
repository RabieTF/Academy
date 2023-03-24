package models

import DB "rabietf.me/go-assignment/db"

type Category struct {
	ID   int64
	Name string
}

// Method for finding all categories in database.
// Returns (categories, nil) if successful.
// Returns (nil, err) if something went wrong.
func (category Category) FindAll() ([]Category, error) {
	var categories []Category

	rows, err := DB.Connection.Query("SELECT * FROM Categories")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var cat Category
		if err := rows.Scan(&cat.ID, &cat.Name); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}

	return categories, nil
}
