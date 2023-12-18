package main

import (
	"github.com/Shreyas-Prabhu/CRUDGO/initializers"
	"github.com/Shreyas-Prabhu/CRUDGO/models"
)

func init() {
	initializers.InitEnv()
	initializers.LoadDBConfig()
}

func main() {
	initializers.DB.AutoMigrate(&models.Employee{})
}
