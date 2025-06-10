package handlers

import (
	"fmt"
	"mainapp/pkg/models"
	"mainapp/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListAllVotes(c *gin.Context) {
	votes := services.GetAllVotes()

	c.JSON(http.StatusOK, gin.H{"data": votes})
}

func ListPollVotes(c *gin.Context) {
	id := c.Param("pollId")

	parsedUUID, _ := uuid.Parse(id)

	votes := services.GetPollVotes(parsedUUID)

	c.JSON(http.StatusOK, gin.H{"message": "Vote created", "data": votes})
}

func SubmitVote(c *gin.Context) {
	var req models.UserResponse

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vote := &req

	err := services.AddVote(*vote)
	if err != nil {
		fmt.Println()
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vote created", "data": vote})
}

func DeleteVote(c *gin.Context) {
	id := c.Param("id")

	parsedUUID, _ := uuid.Parse(id)

	services.RemoveVoteById(parsedUUID)

	c.JSON(http.StatusOK, gin.H{"deleted_vote_id": id})
}
