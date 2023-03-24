package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rabietf.me/go-assignment/models"
)

// POST request at /shops, creates a new shop linked to the authenticated user.
// USER MUST BE AUTHENTICATED TO PERFORM THIS REQUEST.
// 201 if successful.
// 500 if internal error.
// 400 if incorrect format.
func CreateShop(c *gin.Context) {
	var newShop models.Shop

	if err := c.BindJSON(&newShop); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Incorrect format, please send in JSON: name and address"})
		return
	}

	userID, ok := c.Get("userID")

	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	newShop.OwnerID = int64(userID.(float64)) // Had to convert it to float and then to int cause otherwise it'd crash

	id, err := newShop.Save()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"shopId": id, "message": "You created a shop!"})
}

// GET request at /shops,
// 200 and all the shops if successful.
// 404 if there are no shops in database.
// 500 if internal error.
func GetShops(c *gin.Context) {
	var shop models.Shop

	shops, err := shop.FindAll()

	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No shop found in database."})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	c.IndentedJSON(http.StatusOK, shops)
	return
}

// GET request at /shops/:id,
// 200 and the requested shop if successful.
// 404 if the requested shop doesn't exist in database.
// 500 if internal error.
func GetShopById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Incorrect format, please enter a correct ID."})
		return
	}

	var shop models.Shop

	ok, err := shop.FindById(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Shop doesn't not exist."})
		return
	}

	c.IndentedJSON(http.StatusOK, shop)
	return
}

// PUT request at /shops/:id
// 200 if successful.
// 400 if bad formatting.
// 403 if user isn't owner of this shop.
// 404 if shop doesn't exist.
// 500 if something went wrong.
func EditShop(c *gin.Context) {
	var shop models.Shop
	var newShop models.Shop

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Incorrect format, please enter a correct ID."})
		return
	}

	if err := c.BindJSON(&newShop); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Incorrect format, please send in JSON: name and address"})
		return
	}

	ctxID, ok := c.Get("userID")

	userID := int64(ctxID.(float64))

	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	ok, err = shop.FindById(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Shop doesn't not exist."})
		return
	}

	if userID != shop.OwnerID {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "User doesn't have permission to perform an update on this store."})
		return
	}

	err = shop.Update(newShop)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Shop updated successfuly."})
	return

}

// DELETE request at /shops/:id
// 200 if successful.
// 403 if user doesn't own this shop.
// 404 if shop doesn't exist.
// 500 if something went wrong
func DeleteShop(c *gin.Context) {
	var shop models.Shop

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Incorrect format, please enter a correct ID."})
		return
	}

	ctxID, ok := c.Get("userID")

	userID := int64(ctxID.(float64))

	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	ok, err = shop.FindById(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Shop doesn't not exist."})
		return
	}

	if userID != shop.OwnerID {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "User doesn't have permission to perform an delete this store."})
		return
	}

	err = shop.Delete()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Shop deleted successfuly."})
	return
}
