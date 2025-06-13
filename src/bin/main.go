package main

import (
	"mainapp/pkg/global"
	"mainapp/pkg/handlers"

	"github.com/gin-gonic/gin"
)


func main() {
	globals.Init("mongodb://localhost:27017/")
	defer globals.DeInit()

	router := gin.Default()

	// Auth
	router.POST("/login", handlers.LoginHandler)
	router.POST("/register", handlers.RegisterHandler)
	router.POST("/logout", handlers.LogoutHandler)

	api := router.Group("/api", handlers.AuthMiddleware) 
	{
		// Posts endpoints
		api.GET("/polls", handlers.ListAllPolls)
		api.POST("/polls", handlers.CreateNewPoll)
		api.PATCH("/polls/:id", handlers.DeletePoll)
		api.DELETE("/polls/:id", handlers.UpdatePoll)

		// Vote endpoints
		api.GET("/votes/", handlers.ListAllVotes)
		api.POST("/votes/", handlers.SubmitVote)
		api.GET("/votes/:id", handlers.DeleteVote)
		api.DELETE("/votes/:pollId", handlers.ListPollVotes)
	}

	router.Run()
}
