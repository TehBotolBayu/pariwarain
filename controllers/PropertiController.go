package controllers

import (
	"net/http"

	"github.com/TehBotolBayu/pariwarainAPI/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PropertiController struct {
	DB *gorm.DB
}

func NewPropertiController(db *gorm.DB) *PropertiController {
	return &PropertiController{
		DB: db,
	}
}

func (pc *PropertiController) GetAllProperties(c *gin.Context) {
	var properties []models.Properti
	if err := pc.DB.Find(&properties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, properties)
}

func (pc *PropertiController) CreateProperty(c *gin.Context) {
	var property models.Properti
	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := pc.DB.Create(&property).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, property)
}

func (pc *PropertiController) UpdateProperty(c *gin.Context) {
	var property models.Properti
	propertyID := c.Param("id")

	if err := pc.DB.First(&property, "id = ?", propertyID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}

	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.DB.Save(&property).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, property)
}

func (pc *PropertiController) DeleteProperty(c *gin.Context) {
	var property models.Properti
	propertyID := c.Param("id")

	if err := pc.DB.First(&property, "id = ?", propertyID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}

	if err := pc.DB.Delete(&property).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Property deleted successfully"})
}
