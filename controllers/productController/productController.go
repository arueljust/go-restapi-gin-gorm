package productController

import (
	"encoding/json"
	"net/http"

	"github.com/arueljust/databases"
	"github.com/arueljust/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(c *gin.Context) {
	var products []models.Product
	databases.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"message": "data berhasil didapat",
		"data":    products,
	})

}
func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	err := databases.DB.First(&product, id).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data tidak ditemukan!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}
func Create(c *gin.Context) {
	var product models.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	databases.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil ditambah", "data": product})

}
func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	update := databases.DB.Model(&product).Where("id = ?", id).Updates(&product)
	if update.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "data tidak dapat diupdate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diupdate", "data": &product})

}
func Delete(c *gin.Context) {
	var product models.Product
	var input struct {
		Id json.Number
	}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id, _ := input.Id.Int64()
	delete := databases.DB.Delete(&product, id)
	if delete.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "gagal hapus / data tidak ada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})
}
