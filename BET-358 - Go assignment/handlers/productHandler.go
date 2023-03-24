package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"rabietf.me/go-assignment/models"
)

// Helper function that checks that all of the given categories are valid and exist in the database.
func checkAllElements(categories []string, dbCategories []models.Category) bool {
	set := make(map[string]bool)

	for _, v := range dbCategories {
		set[v.Name] = true
	}

	for _, v := range categories {
		if _, ok := set[strings.TrimSpace(v)]; !ok {
			return false
		}
	}

	return true
}

// POST request at /products, creates a new product within the defined shop (shopID)
// Verifies that user is creating product at a shop he owns.
// User must be authenticated.
// 201 if successful.
// 500 if something went wrong.
// 400 if incorrect JSON format.
// 403 if user is attempting to create a new product in a shop he doesn't own.
func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	ctxId, ok := c.Get("userID")

	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	userID := int64(ctxId.(float64))

	if err := c.BindJSON(&newProduct); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Incorrect format, please send in JSON: ShopID, Name, Description and Categories all in one string separated by a comma."})
		return
	}

	categories := strings.Split(newProduct.Categories, ",")

	var category models.Category

	dbCategories, err := category.FindAll()

	if ok = checkAllElements(categories, dbCategories); !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "One of the categories you mentionned is not a correct category, please check GET /categories to know the correct categories."})
		return
	}

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	var shop models.Shop

	ok, err = shop.FindById(newProduct.ShopID)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Shop doesn't not exist."})
		return
	}

	if shop.OwnerID != userID {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "User doesn't have permission to add product to this store."})
		return
	}

	id, err := newProduct.Save()

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"productId": id, "message": "You created a new product!"})
	return

}

// GET request at /products,
// 200 and all the products if successful.
// 404 if there are no products in database.
// 500 if internal error.
func GetProducts(c *gin.Context) {
	var product models.Product

	products, err := product.FindAll()

	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No product found in database."})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	c.IndentedJSON(http.StatusOK, products)
	return
}

// GET request at /products/:id,
// 200 and the requested product if successful.
// 404 if the requested product doesn't exist in database.
// 500 if internal error.
func GetProductById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Incorrect format, please enter a correct ID."})
		return
	}

	var product models.Product

	ok, err := product.FindById(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product doesn't not exist."})
		return
	}

	c.IndentedJSON(http.StatusOK, product)
	return
}

// PUT request at /products/:id
// 200 if successful.
// 400 for bad formatting.
// 403 if user isn't owner of the shop where the product belongs.
// 404 if product doesn't exist.
// 500 if something went wrong.
func EditProduct(c *gin.Context) {
	var product models.Product
	var newProduct models.Product

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Please enter a correct ID."})
		return
	}

	if err := c.BindJSON(&newProduct); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Incorrect format, please send in JSON: name, description and categories: in one string separated by a comma."})
		return
	}

	ctxID, ok := c.Get("userID")

	userID := int64(ctxID.(float64))

	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	ok, err = product.FindById(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product doesn't not exist."})
		return
	}

	var shop models.Shop

	shop.FindById(product.ShopID)

	if userID != shop.OwnerID {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "User doesn't have permission to perform an update on this product."})
		return
	}

	categories := strings.Split(newProduct.Categories, ",")

	var category models.Category

	dbCategories, err := category.FindAll()

	if ok = checkAllElements(categories, dbCategories); !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "One of the categories you mentionned is not a correct category, please check GET /categories to know the correct categories."})
		return
	}

	err = product.Update(newProduct)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Product updated successfuly."})
	return

}

// DELETE request at /products/:id
// 200 if successful.
// 400 for bad formatting.
// 403 if user isn't owner of the shop where the product belongs.
// 404 if product doesn't exist.
// 500 if something went wrong.
func DeleteProduct(c *gin.Context) {
	var product models.Product

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Please enter a correct ID."})
		return
	}

	ctxID, ok := c.Get("userID")

	userID := int64(ctxID.(float64))

	ok, err = product.FindById(id)

	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product doesn't not exist."})
		return
	}

	var shop models.Shop

	shop.FindById(product.ShopID)

	if userID != shop.OwnerID {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "User doesn't have permission to delete this product."})
		return
	}

	err = product.Delete()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Product deleted successfuly."})
	return

}
