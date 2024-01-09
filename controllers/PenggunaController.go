package controllers

import (
	"net/http"

	"github.com/TehBotolBayu/pariwarainAPI/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PenggunaController struct {
	DB *gorm.DB
}

func NewPenggunaController(db *gorm.DB) *PenggunaController {
	return &PenggunaController{
		DB: db,
	}
}

func (uc *PenggunaController) GetAllPengguna(c *gin.Context) {
	var Penggunas []models.Pengguna
	if err := uc.DB.Find(&Penggunas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Penggunas)
}

func (uc *PenggunaController) CreatePengguna(c *gin.Context) {
	var Pengguna models.Pengguna
	if err := c.ShouldBindJSON(&Pengguna); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.DB.Create(&Pengguna).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, Pengguna)
}

func (uc *PenggunaController) UpdatePengguna(c *gin.Context) {
	var Pengguna models.Pengguna
	PenggunaID := c.Param("id")

	if err := uc.DB.First(&Pengguna, "id = ?", PenggunaID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna not found"})
		return
	}

	if err := c.ShouldBindJSON(&Pengguna); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.DB.Save(&Pengguna).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Pengguna)
}

func (uc *PenggunaController) DeletePengguna(c *gin.Context) {
	var Pengguna models.Pengguna
	PenggunaID := c.Param("id")

	if err := uc.DB.First(&Pengguna, "id = ?", PenggunaID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna not found"})
		return
	}

	if err := uc.DB.Delete(&Pengguna).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pengguna deleted successfully"})
}
