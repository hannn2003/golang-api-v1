package main

import (
	"todo-app/db"
	"todo-app/routes"

	"github.com/gin-gonic/gin"
)
func main() {
	db.Connect()
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run()
}