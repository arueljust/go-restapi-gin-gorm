package main

import (
	"github.com/arueljust/controllers/productController"
	"github.com/arueljust/databases"
	"github.com/gin-gonic/gin"
)

func init() {
	databases.Conn()
	databases.Load()
	databases.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/v1/products", productController.GetAll)
	r.GET("/v1/product/:id", productController.Show)
	r.POST("/v1/product", productController.Create)
	r.PUT("/v1/product/:id", productController.Update)
	r.DELETE("/v1/product", productController.Delete)

	r.Run()
}
