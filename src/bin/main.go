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

	// Posts endpoints
	router.GET("/polls", handlers.ListAllPolls)
	router.POST("/polls", handlers.CreateNewPoll)
	router.DELETE("/polls/:id", handlers.DeletePoll)
	router.PATCH("/polls/:id", handlers.UpdatePoll)
	
	// Vote endpoints
	router.GET("/votes/", handlers.ListAllVotes)
	router.POST("/votes/", handlers.SubmitVote)
	router.DELETE("/votes/:id", handlers.DeleteVote)
	router.GET("/votes/:pollId", handlers.ListPollVotes)
	
	router.Run()
}
