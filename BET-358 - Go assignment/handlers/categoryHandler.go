package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rabietf.me/go-assignment/models"
)

// GET request at /categories
// 200 and all the predefined categories
// 500 if something went wrong
func GetCategories(c *gin.Context) {
	var category models.Category

	categories, err := category.FindAll()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Please contact your administrator."})
		return
	}

	c.IndentedJSON(http.StatusOK, categories)
	return
}
