package main

import (
	"github.com/sid/go-crud/initializers"
	"github.com/sid/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
    initializers.DB.AutoMigrate(&models.Post{})
}