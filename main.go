package main

import (
	"fmt"

	"github.com/Shreyas-Prabhu/CRUDGO/controllers"
	"github.com/Shreyas-Prabhu/CRUDGO/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.InitEnv()
	initializers.LoadDBConfig()
}

func main() {
	fmt.Println("CRUD with GORM , GiN and MySQL")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello USer",
		})
	})

	r.GET("/employees", controllers.GetAllEmployee)
	r.GET("/employee/:id", controllers.GetEmployee)
	r.POST("/employee", controllers.PostEmployee)
	r.PUT("/employee/:id", controllers.UpdateEmployee)
	r.DELETE("/employee/:id", controllers.DeleteEmployee)
	r.Run()
}
