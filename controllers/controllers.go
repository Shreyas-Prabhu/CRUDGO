package controllers

import (
	"log"

	"github.com/Shreyas-Prabhu/CRUDGO/initializers"
	"github.com/Shreyas-Prabhu/CRUDGO/models"
	"github.com/gin-gonic/gin"
)

func PostEmployee(c *gin.Context) {
	//Request Body

	var body struct {
		Name    string
		Company string
	}

	c.Bind(&body)
	emp := models.Employee{Name: body.Name, Company: body.Company}
	result := initializers.DB.Create(&emp)
	if result.Error != nil {
		log.Fatal("Cannot push the employee")
	}

	c.JSON(200, gin.H{
		"Inserted Employee": emp,
	})
}

func GetAllEmployee(c *gin.Context) {
	var employees []models.Employee
	result := initializers.DB.Find(&employees)
	if result.Error != nil {
		log.Fatal("Cannot push the employee")
	}
	c.JSON(200, gin.H{
		"All Employee": employees,
	})
}

func GetEmployee(c *gin.Context) {

	id := c.Param("id")
	var emp models.Employee
	initializers.DB.First(&emp, id)

	c.JSON(200, gin.H{
		"Employee": emp,
	})
}

func UpdateEmployee(c *gin.Context) {

	id := c.Param("id")
	var emp models.Employee
	initializers.DB.First(&emp, id)

	var body struct {
		Name    string
		Company string
	}

	c.Bind(&body)

	initializers.DB.Model(&emp).Updates(models.Employee{Name: body.Name, Company: body.Company})

	c.JSON(200, gin.H{
		"Updated Employee": emp,
	})
}

func DeleteEmployee(c *gin.Context) {

	id := c.Param("id")
	initializers.DB.Delete(&models.Employee{}, id)

	c.JSON(200, gin.H{
		"message": "deleted",
	})
}
