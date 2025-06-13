package handlers

import (
	"mainapp/pkg/models"
	"mainapp/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListAllVotes(c *gin.Context) {
	votes := services.GetAllVotes()

	c.JSON(http.StatusOK, gin.H{"data": votes})
}

func ListUserVotes(c *gin.Context) {
	userID, err := GetUserIdFromCookie(c)
	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	votes, err := services.GetUserVotes(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": votes})
}

func SubmitVote(c *gin.Context) {
	// no cookie
	userID, err := GetUserIdFromCookie(c)
	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var req models.UserResponse
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vote := &req
	vote.UserId = userID

	savedVote, err := services.AddVote(*vote)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vote created", "data": savedVote})
}

func DeleteVote(c *gin.Context) {
	userID, err := GetUserIdFromCookie(c)
	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	parsedObjectId, _ := primitive.ObjectIDFromHex(id)

	err = services.RemoveVoteById(parsedObjectId, userID)
	
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"deleted_vote_id": id})
}
