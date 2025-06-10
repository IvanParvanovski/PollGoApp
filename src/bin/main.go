package main

import (
	"github.com/gin-gonic/gin"
	"mainapp/pkg/handlers"
)


func main() {
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
