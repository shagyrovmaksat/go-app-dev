package main

import (
	"rest/handlers"
	"rest/models"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	models.ConnectDB()

	route.GET("/courses", handlers.GetAllCourses)
	route.GET("/course/:id", handlers.GetCourseById)
	route.GET("/course", handlers.GetCoureByAuthor)
	route.POST("/course", handlers.CreateCourse)
	route.PUT("/course/:id", handlers.UpdateCourse)
	route.DELETE("/course/:id", handlers.DeleteCourse)

	route.Run()
}
