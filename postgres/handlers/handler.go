package handlers

import (
	"net/http"
	"rest/models"

	"github.com/gin-gonic/gin"
)

type CreateCourseInput struct {
	Author      string `json:"author" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Cost        int    `json:"cost" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateCourseInput struct {
	Author      string `json:"author" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Cost        int    `json:"cost" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func GetAllCourses(context *gin.Context) {
	var courses []models.Course
	models.DB.Find(&courses)

	context.JSON(http.StatusOK, gin.H{"courses": courses})
}

func GetCourseById(context *gin.Context) {
	var course models.Course
	if err := models.DB.Where("id=?", context.Param("id")).First(&course).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Course doesn't exists"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"course": course})
}

func GetCoureByAuthor(context *gin.Context) {
	var author string
	var found bool
	if author, found = context.GetQuery("author"); !found {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Sorry, there's no query param 'author'"})
		return
	}

	var course models.Course
	if err := models.DB.Where("author=?", author).First(&course).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Course doesn't exists"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"course": course})
}

func CreateCourse(context *gin.Context) {
	var input CreateCourseInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course := models.Course{Author: input.Author, Title: input.Title, Cost: input.Cost, Description: input.Description}
	models.DB.Create(&course)

	context.JSON(http.StatusOK, gin.H{"course": course})
}

func UpdateCourse(context *gin.Context) {
	var course models.Course
	if err := models.DB.Where("id=?", context.Param("id")).First(&course).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Course with such id doesn't exists"})
		return
	}

	var input UpdateCourseInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&course).Update(input)
	context.JSON(http.StatusOK, gin.H{"course": course})
}

func DeleteCourse(context *gin.Context) {
	var course models.Course

	if err := models.DB.Where("id=?", context.Param("id")).First(&course).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Course doesn't exists"})
		return
	}

	models.DB.Delete(&course)
	context.JSON(http.StatusOK, gin.H{"deleted": true})
}
