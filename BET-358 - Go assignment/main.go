package main

import (
	"github.com/gin-gonic/gin"
	DB "rabietf.me/go-assignment/db"
	"rabietf.me/go-assignment/handlers"
	"rabietf.me/go-assignment/middlewares"
)

func main() {
	router := gin.Default()
	DB.ConnectToDB()

	router.POST("/users", handlers.SignUp)
	router.POST("/login", handlers.SignIn)

	router.POST("/shops", middlewares.VerifyAuth(), handlers.CreateShop)
	router.GET("/shops", handlers.GetShops)
	router.GET("/shops/:id", handlers.GetShopById)
	router.PUT("/shops/:id", middlewares.VerifyAuth(), handlers.EditShop)
	router.DELETE("/shops/:id", middlewares.VerifyAuth(), handlers.DeleteShop)

	router.POST("/products", middlewares.VerifyAuth(), handlers.CreateProduct)
	router.GET("/products", handlers.GetProducts)
	router.GET("/products/:id", handlers.GetProductById)
	router.PUT("/products/:id", middlewares.VerifyAuth(), handlers.EditProduct)
	router.DELETE("/products/:id", middlewares.VerifyAuth(), handlers.DeleteProduct)

	router.GET("/categories", handlers.GetCategories)

	router.Run("localhost:8080")
}
