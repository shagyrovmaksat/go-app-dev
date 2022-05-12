package main

import (
	"bufio"
	"context"
	api "endtermProject/course"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

var mainMenu = "Enter number of action:\n" +
	"1) Create course\n" +
	"2) Delete course\n" +
	"3) Get course\n" +
	"4) Update course\n" +
	"5) Exit\n"

func main() {
	fmt.Println("Course client started")

	connection, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	c := api.NewCourseServiceClient(connection)

	run := true
	var courseId string
	var courseTitle string
	var courseDescription string
	var courseAuthor string
	var courseCategory string
	var courseCost int
	reader := bufio.NewReader(os.Stdin)
	for {
		if !run {
			break
		}

		fmt.Println(mainMenu)
		var action int
		fmt.Scanf("%d", &action)
		switch action {
		case 1:
			fmt.Println("Enter course id: ")
			fmt.Scanln(&courseId)
			fmt.Println("Enter course title: ")
			courseTitle, err = reader.ReadString('\n')
			fmt.Println("Enter course description: ")
			courseDescription, err = reader.ReadString('\n')
			fmt.Println("Enter course author: ")
			courseAuthor, err = reader.ReadString('\n')
			fmt.Println("Enter course category: ")
			courseCategory, err = reader.ReadString('\n')
			fmt.Println("Enter course cost: ")
			fmt.Scanf("%d", &courseCost)

			createCourseResponse, err := c.CreateCourse(context.Background(), &api.CreateCourseRequest{
				Id:          courseId,
				Title:       courseTitle,
				Description: courseDescription,
				Author:      courseAuthor,
				Category:    courseCategory,
				Cost:        int32(courseCost),
			})
			if err != nil {
				log.Fatalf("Unexpected error: %v\n", err)
			}
			fmt.Printf(":) Course has been created: %v\n", createCourseResponse)
		case 2:
			fmt.Println("Enter course id: ")
			fmt.Scanln(&courseId)
			//delete
			deleteCourseResponse, err := c.DeleteCourse(context.Background(), &api.DeleteCourseRequest{Id: courseId})
			if err != nil {
				fmt.Printf("Error happened while deleting: %v\n", err)
			}
			fmt.Printf(":( Course was deleted. %v\n", deleteCourseResponse)
		case 3:
			//get
			fmt.Println("Enter course id: ")
			fmt.Scanln(&courseId)
			getCourseRequest := &api.GetCourseRequest{Id: courseId}
			getCourseResponse, err := c.GetCourse(context.Background(), getCourseRequest)
			if err != nil {
				fmt.Printf("Error happened while geting course: %v \n", err)
			}
			fmt.Printf(";з Get course result: %v\n", getCourseResponse)
		case 4:
			//update
			fmt.Println("Enter course id: ")
			fmt.Scanln(&courseId)
			fmt.Println("Enter course title: ")
			courseTitle, err = reader.ReadString('\n')
			fmt.Println("Enter course description: ")
			courseDescription, err = reader.ReadString('\n')
			fmt.Println("Enter course author: ")
			courseAuthor, err = reader.ReadString('\n')
			fmt.Println("Enter course category: ")
			courseCategory, err = reader.ReadString('\n')
			fmt.Println("Enter course cost: ")
			fmt.Scanf("%d", &courseCost)

			updateCourseResponse, err := c.UpdateCourse(context.Background(), &api.UpdateCourseRequest{
				Id:          courseId,
				Title:       courseTitle,
				Description: courseDescription,
				Author:      courseAuthor,
				Category:    courseCategory,
				Cost:        int32(courseCost),
			})
			if err != nil {
				fmt.Printf("Error happened while updating: %v\n", err)
			}
			fmt.Printf("^0^ Course was updated: %v\n", updateCourseResponse)
		default:
			run = !run
		}
	}
	fmt.Println("Course client finished.")
}
