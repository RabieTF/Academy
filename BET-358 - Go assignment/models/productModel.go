package models

import (
	"database/sql"

	DB "rabietf.me/go-assignment/db"
)

type Product struct {
	ID          int64
	ShopID      int64
	Name        string
	Description string
	Categories  string
}

// Method for inserting new product in database.
// Returns (productId, nil) if successful.
// Returns (0, err) if failed.
func (product Product) Save() (int64, error) {
	result, err := DB.Connection.Exec("INSERT INTO Products (shop_id, name, description, categories) VALUES (?, ?, ?, ?)", product.ShopID, product.Name, product.Description, product.Categories)

	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Method for finding product in database using id.
// Returns (true, nil) and puts product in object if product exists.
// Returns (false, nil) if product doesn't exist.
// Returns (false, err) if something went wrong.
func (product *Product) FindById(ID int64) (bool, error) {

	row := DB.Connection.QueryRow("SELECT * FROM Products WHERE ID = ?", ID)

	if err := row.Scan(&product.ID, &product.ShopID, &product.Name, &product.Description, &product.Categories); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// Method for finding all products in database.
// Returns (products, nil) if successful.
// Returns (nil, err) if something went wrong.
func (product Product) FindAll() ([]Product, error) {
	var products []Product

	rows, err := DB.Connection.Query("SELECT * FROM Products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var prd Product
		if err := rows.Scan(&prd.ID, &prd.ShopID, &prd.Name, &prd.Description, &prd.Categories); err != nil {
			return nil, err
		}
		products = append(products, prd)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// Method for updating an element in database
// Takes new data as paramater, updates the data of the ID in the connected object.
// Returns nil if success.
// Returns error otherwise
func (product Product) Update(newProduct Product) error {
	_, err := DB.Connection.Exec("UPDATE Products SET name=?, description=?, categories=? WHERE id=?", newProduct.Name, newProduct.Description, newProduct.Categories, product.ID)

	if err != nil {
		return err
	}

	return nil
}

// Method for deleting an element in database
// Uses ID of object to delete said data.
// Returns nil if success.
// Returns error otherwise
func (product Product) Delete() error {
	_, err := DB.Connection.Exec("DELETE FROM Products WHERE id=?", product.ID)

	if err != nil {
		return err
	}

	return nil
}
