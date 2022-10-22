package controller

import (
	"net/http"

	"github.com/Tepph4/SA-Project/entity"
	"github.com/gin-gonic/gin"
)


// POST /buildings
func CreateBuilding(c *gin.Context) {
	var building entity.Building
	if err := c.ShouldBindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&building).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": building})
}

// GET /buildings
// List all buildings
func ListBuildings(c *gin.Context) {
	var buildings []entity.User
	if err := entity.DB().Raw("SELECT * FROM Buildings").Scan(&buildings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": buildings})
}

// GET /buildibg/:id
// Get buildibg by id
func GetBuilding(c *gin.Context) {
	var building entity.Building
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&building); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "building not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": building})
}

// DELETE /buildings/:id
func DeleteBuilding(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM building WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "building not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /buildings
func UpdateBuilding(c *gin.Context) {
	var building entity.Building
	if err := c.ShouldBindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", building.ID).First(&building); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "building not found"})
		return
	}

	if err := entity.DB().Save(&building).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": building})
}

