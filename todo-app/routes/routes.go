package routes

import (
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	todoRoutes := r.Group("/todos") 
	{
		todoRoutes.GET("/", controllers.CreateTodo)
		todoRoutes.POST("/create", controllers.CreateTodo)
		todoRoutes.PUT("/update/:id", controllers.UpdateTodo)
		todoRoutes.DELETE("/delete/:id", controllers.DeleteTodo)
		todoRoutes.PUT("/soft-delete/:id", controllers.SoftDeleteTodo)
		todoRoutes.PUT("/upload-file", controllers.UploadImage)
	}
}