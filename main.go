package main

import (
	"github.com/0xJungleMonkey/timemanagementapp/controllers"
	"github.com/0xJungleMonkey/timemanagementapp/database"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	database.ConnectDatabase()
	r.GET("/records", controllers.FindRecords)
	r.GET("/record/:id", controllers.FindRecord)
	r.POST("/records", controllers.CreateRecord)
	r.PATCH("/records/:id", controllers.UpdateRecord)
	r.DELETE("/records/:id", controllers.DeleteRecord)
	r.Run("localhost:8080")
}
