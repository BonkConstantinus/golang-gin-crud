package controllers

import (
	"fmt"
	"net/http"
	"toko/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateItemInput struct {
	Nama  string             `json:"nama"`
	Pajak []CreatePajakInput `json:"pajak"`
}
type CreatePajakInput struct {
	NamaPajak  string  `json:"namapajak"`
	RatesPajak float32 `json:"ratespajak"`
}

func CreateItem(c *gin.Context) {
	var input CreateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(input.Pajak)
	var listPajak []models.Pajake
	for _, pajak := range input.Pajak {
		pajake := models.Pajake{Nama: pajak.NamaPajak, Rate: pajak.RatesPajak}
		listPajak = append(listPajak, pajake)
	}
	task := models.Item{Nama: input.Nama, Pajake: listPajak}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})

}

func FindItem(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var item []models.Item
	db.Preload("Pajake").Find(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}
func DFindItemID(c *gin.Context) {
	var item models.Item
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Preload("Pajake").Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func DeleteItem(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var item models.Item
	if err := db.Preload("Pajake").Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	db.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"data": true})

}

func UpdateItem(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var item models.Item
	if err := db.Preload("Pajake").Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input CreateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.Nama = input.Nama

	listPajak := input.Pajak
	tampung := item.Pajake
	for i := 0; i < len(listPajak); i++ {
		tampung[i].Nama = listPajak[i].NamaPajak
		tampung[i].Rate = listPajak[i].RatesPajak
	}
	db.Save(&item)

	c.JSON(http.StatusOK, gin.H{"data": item})
}
