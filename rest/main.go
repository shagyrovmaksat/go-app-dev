package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Course struct {
	ID          string `json:"id"`
	Title       string `json:"artist"`
	Cost        int    `json:"cost"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

var courses = []Course{
	{ID: "1", Title: "Course 1", Author: "John", Cost: 12, Description: "jdafl;jdslfjds"},
	{ID: "2", Title: "Course 2", Author: "ABC", Cost: 22, Description: "jdafl;jdslfjds"},
	{ID: "3", Title: "Course 3", Author: "xD", Cost: 10, Description: "jdafl;jdslfjds"},
}

func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func postCourse(c *gin.Context) {
	var newCourse Course
	if err := c.BindJSON(&newCourse); err != nil {
		fmt.Println(err)
		return
	}
	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

func getCourseByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range courses {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Course not found"})
}

func deleteCourseByID(c *gin.Context) {
	id := c.Param("id")
	for i, a := range courses {
		if a.ID == id {
			courses = append(courses[:i], courses[i+1:]...)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Course not found"})
}

func main() {
	router := gin.Default()
	router.GET("/courses", getCourses)
	router.POST("/course", postCourse)
	router.GET("/course/:id", getCourseByID)
	router.DELETE("/course/:id", deleteCourseByID)

	router.Run("localhost:8080")
}
