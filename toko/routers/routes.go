package routers

import (
	"toko/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)

	})
	r.GET("/item", controllers.FindItem)
	r.POST("/item", controllers.CreateItem)
	r.GET("/item/:id", controllers.DFindItemID)
	r.DELETE("/item/:id", controllers.DeleteItem)
	r.PATCH("/item/:id", controllers.UpdateItem)

	return r
}
