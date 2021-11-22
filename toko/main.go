package main

import (
	"toko/models"
	"toko/routers"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Pajake{})

	r := routers.SetupRoutes(db)
	r.Run()

}
