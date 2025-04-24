package controllers

import (
	"net/http"
	"vse-bank/models"

	"github.com/gin-gonic/gin"
)

func GetBanks(c *gin.Context) {
	var banks []models.Bank
	if err := DB.Find(&banks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, banks)
}

func AddBank(c *gin.Context) {
	var bank models.Bank
	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&bank).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, bank)
}

func UpdateBank(c *gin.Context) {
	id := c.Param("id")
	var bank models.Bank

	if err := DB.First(&bank, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bank not found"})
		return
	}

	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Save(&bank).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bank)
}

func DeleteBank(c *gin.Context) {
	id := c.Param("id")
	var bank models.Bank

	if err := DB.First(&bank, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bank not found"})
		return
	}

	if err := DB.Delete(&bank).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bank deleted successfully"})
}
