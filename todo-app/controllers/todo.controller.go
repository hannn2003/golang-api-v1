package controllers

import (
	"todo-app/services"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	services.GetTodos(c)
}

func CreateTodo(c *gin.Context){
	services.CreateTodo(c)
}

func UpdateTodo(c *gin.Context){
	services.UpdateTodo(c)
}

func DeleteTodo(c *gin.Context) {
	services.DeleteTodo(c)
}

func SoftDeleteTodo(c *gin.Context){
	services.SoftDeleteTodo(c)
}

func UploadImage(c *gin.Context) {
	services.UploadFile(c)
}