package controllers

import (
	"net/http"
	"vse-bank/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLoans(c *gin.Context) {
	limit := 10
	page := 1

	if limitParam := c.DefaultQuery("limit", "10"); limitParam != "" {
		var err error
		limit, err = strconv.Atoi(limitParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
			return
		}
	}

	if pageParam := c.DefaultQuery("page", "1"); pageParam != "" {
		var err error
		page, err = strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
			return
		}
	}

	offset := (page - 1) * limit

	var loans []models.Loan

	if err := DB.Preload("Bank").Preload("User").
		Limit(limit).
		Offset(offset).
		Find(&loans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loans)
}

func AddLoan(c *gin.Context) {
	var loan models.Loan
	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if loan.Amount <= 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Loan amount should be greater than 1"})
		return
	}

	if err := DB.Create(&loan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, loan)
}

func UpdateLoan(c *gin.Context) {
	id := c.Param("id")
	var loan models.Loan

	// Check if the loan exists
	if err := DB.First(&loan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	// Bind incoming JSON
	var updatedLoan models.Loan
	if err := c.ShouldBindJSON(&updatedLoan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loan.Amount = updatedLoan.Amount
	loan.BankID = updatedLoan.BankID
	loan.UserID = updatedLoan.UserID
	loan.LoanTypeID = updatedLoan.LoanTypeID

	if err := DB.Save(&loan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loan)
}

func DeleteLoan(c *gin.Context) {
	id := c.Param("id")
	var loan models.Loan

	if err := DB.First(&loan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	if err := DB.Delete(&loan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan deleted successfully"})
}
