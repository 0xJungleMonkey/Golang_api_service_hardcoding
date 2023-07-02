package controllers

import (
	"net/http"
	"time"

	"github.com/0xJungleMonkey/timemanagementapp/database"
	"github.com/0xJungleMonkey/timemanagementapp/models"
	"github.com/gin-gonic/gin"
)

type CreateRecordInput struct {
	Start    time.Time
	Stop     time.Time
	Duration int
	Project  string
}

// Get /Records
// Find all records
func FindRecords(c *gin.Context) {
	var records []models.Record
	database.DB.Find(&records)
	c.JSON(http.StatusOK, gin.H{"data": records})
	//In the context of gorm.DB, the First() method does not directly return the retrieved record. Instead, it modifies the provided record2 variable and returns the gorm.DB object itself.
	//So when you write result2 := db.First(&record2), the result2 variable is of type *gorm.DB, not the record2 variable. That's why result2.ID is undefined because the *gorm.DB type does not have an ID field.
	//To access the retrieved record's fields, you need to use the record2 variable, which was passed as a reference to the First() method. This method modifies the record2 variable with the fetched record.
}

// Get /records/:id
// Find record with certain ID
func FindRecord(c *gin.Context) {
	var record models.Record
	if err := database.DB.Where("id=?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": record})
}

// Post /records
// Create new record
func CreateRecord(c *gin.Context) {
	//validate input
	var input CreateRecordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//create record
	record := models.Record{Start: input.Start, Stop: input.Stop, Duration: input.Duration, Project: input.Project}
	database.DB.Create(&record)
	c.JSON(http.StatusOK, gin.H{"data": record})
}

// PATCH /records/:id
// Update a record
func UpdateRecord(c *gin.Context) {
	var record models.Record
	if err := database.DB.Where("id=?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//Validate input
	var input CreateRecordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&record).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": record})

}

// Delete /books/:id
// Delete a book
func DeleteRecord(c *gin.Context) {
	var record models.Record
	if err := database.DB.Where("id = ? ", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	database.DB.Delete(&record)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
