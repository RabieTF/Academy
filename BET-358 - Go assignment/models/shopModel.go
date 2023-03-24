package models

import (
	"database/sql"

	DB "rabietf.me/go-assignment/db"
)

type Shop struct {
	ID      int64
	Name    string
	Address string
	OwnerID int64
}

// Method for inserting new shop in database.
// Returns (shopId, nil) if successful.
// Returns (0, err) if failed.
func (shop Shop) Save() (int64, error) {
	result, err := DB.Connection.Exec("INSERT INTO Shops (name, address, owned_by) VALUES (?, ?, ?)", shop.Name, shop.Address, shop.OwnerID)

	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Method for finding shop in database using id.
// Returns (true, nil) and puts shop in object if shop exists.
// Returns (false, nil) if shop doesn't exist.
// Returns (false, err) if something went wrong.
func (shop *Shop) FindById(ID int64) (bool, error) {

	row := DB.Connection.QueryRow("SELECT * FROM Shops WHERE ID = ?", ID)

	if err := row.Scan(&shop.ID, &shop.Name, &shop.Address, &shop.OwnerID); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// Method for finding all shops in database.
// Returns (shops, nil) if successful.
// Returns (nil, err) if something went wrong.
func (shop Shop) FindAll() ([]Shop, error) {
	var shops []Shop

	rows, err := DB.Connection.Query("SELECT * FROM Shops")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var shp Shop
		if err := rows.Scan(&shp.ID, &shp.Name, &shp.Address, &shp.OwnerID); err != nil {
			return nil, err
		}
		shops = append(shops, shp)
	}

	return shops, nil
}

// Method for updating an element in database
// Takes new data as paramater, updates the data of the ID in the connected object.
// Returns nil if success.
// Returns error otherwise
func (shop Shop) Update(newShop Shop) error {
	_, err := DB.Connection.Exec("UPDATE Shops SET name=?, address=? WHERE id=?", newShop.Name, newShop.Address, shop.ID)

	if err != nil {
		return err
	}

	return nil
}

// Method for deleting an element in database
// Uses ID of object to delete said data.
// Returns nil if success.
// Returns error otherwise
func (shop Shop) Delete() error {
	_, err := DB.Connection.Exec("DELETE FROM Shops WHERE id=?", shop.ID)

	if err != nil {
		return err
	}

	return nil
}
