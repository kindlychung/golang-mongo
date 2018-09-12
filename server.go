package main

//created by H.G Nuwan Indika

import (
	// Standard library packages

	// Third party packages
	"fmt"
	"os"
	"strings"

	"github.com/Sri600/golang-mongo/controllers"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func main() {

	// Get a UserController instance
	uc := controllers.NewUserController(getSession())

	// Get a user resource
	router := gin.Default()
	router.GET("/users", uc.UsersList)
	router.GET("/users/:id", uc.GetUser)
	router.DELETE("/users/:id", uc.RemoveUser)
	router.POST("/users", uc.CreateUser)
	router.PUT("/users/:id", uc.UpdateUser)
	println("Listening on port 8000...")
	router.Run(":8000")
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	mongoURL := strings.Replace(os.Getenv("MONGO_PORT"), "tcp", "mongodb", 1)
	if mongoURL == "" {
		panic("MONGO_PORT env var not set.")
	} else {
		fmt.Printf("MONGO_PORT: %s", mongoURL)
	}
	s, err := mgo.Dial(mongoURL)

	// Check if connection error, is mongo running?
	if err != nil {
		println("Is mongo running?")
		panic(err)
	}

	// Deliver session
	return s
}
