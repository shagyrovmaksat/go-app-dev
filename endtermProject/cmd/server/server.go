package main

import (
	"context"
	api "endtermProject/course"
	"endtermProject/models"
	_ "github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	_ "net/http"
)

type server struct {
	api.CourseServiceServer
}

func (*server) CreateCourse(ctx context.Context, req *api.CreateCourseRequest) (*api.CreateCourseResponse, error) {
	id := req.GetId()
	title := req.GetTitle()
	description := req.GetDescription()
	author := req.GetAuthor()
	category := req.GetCategory()
	cost := req.GetCost()

	course := models.Course{
		ID:          id,
		Title:       title,
		Description: description,
		Author:      author,
		Category:    category,
		Cost:        int(cost),
	}
	models.DB.Create(&course)

	return &api.CreateCourseResponse{
		Course: &api.Course{
			Id:          id,
			Title:       title,
			Description: description,
			Author:      author,
			Category:    category,
			Cost:        cost,
		},
	}, nil
}

func (*server) DeleteCourse(ctx context.Context, req *api.DeleteCourseRequest) (*api.DeleteCourseResponse, error) {
	id := req.GetId()
	var course models.Course
	if err := models.DB.Where("id=?", id).First(&course).Error; err != nil {
		return nil, err
	}
	models.DB.Delete(&course)

	return &api.DeleteCourseResponse{
		Id: id,
	}, nil
}

func (*server) UpdateCourse(ctx context.Context, req *api.UpdateCourseRequest) (*api.UpdateCourseResponse, error) {
	id := req.GetId()
	var course models.Course
	if err := models.DB.Where("id=?", id).First(&course).Error; err != nil {
		return nil, err
	}

	title := req.GetTitle()
	description := req.GetDescription()
	author := req.GetAuthor()
	category := req.GetCategory()
	cost := req.GetCost()

	newCourse := models.Course{
		ID:          id,
		Title:       title,
		Description: description,
		Author:      author,
		Category:    category,
		Cost:        int(cost),
	}
	models.DB.Model(&course).Update(newCourse)

	return &api.UpdateCourseResponse{
		Course: &api.Course{
			Id:          id,
			Title:       title,
			Description: description,
			Author:      author,
			Category:    category,
			Cost:        int32(cost),
		},
	}, nil
}

func (*server) GetCourse(ctx context.Context, req *api.GetCourseRequest) (*api.GetCourseResponse, error) {
	id := req.GetId()
	var course models.Course
	if err := models.DB.Where("id=?", id).First(&course).Error; err != nil {
		return nil, err
	}

	title := course.Title
	description := course.Description
	author := course.Author
	category := course.Category
	cost := course.Cost

	return &api.GetCourseResponse{
		Course: &api.Course{
			Id:          id,
			Title:       title,
			Description: description,
			Author:      author,
			Category:    category,
			Cost:        int32(cost),
		},
	}, nil
}

func main() {
	models.ConnectDB()
	s := grpc.NewServer()
	server := &server{}
	api.RegisterCourseServiceServer(s, server)

	t, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(t); err != nil {
		log.Fatal(err)
	}
}
