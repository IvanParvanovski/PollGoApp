package handlers

import (
	"mainapp/pkg/models"
	"mainapp/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListAllPolls(c *gin.Context) {
	polls := services.GetAllPolls()
	
	c.JSON(http.StatusOK, gin.H{"data": polls})
}

func CreateNewPoll(c *gin.Context) {
	userID, err := GetUserIdFromCookie(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var req models.Poll

	if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	poll := &req
	poll.UserId = userID

	savedPoll, err := services.AddPoll(*poll)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "Poll created", "data": savedPoll})
}


func DeletePoll(c *gin.Context) {
	userID, err := GetUserIdFromCookie(c)
	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	parsedObjectId, _ := primitive.ObjectIDFromHex(id)

	err = services.RemovePollById(parsedObjectId, userID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted_poll_id": id})
}

func UpdatePoll(c *gin.Context) {
	userID, err := GetUserIdFromCookie(c)
	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	parsedObjectId, _ := primitive.ObjectIDFromHex(id)
	
	var res models.PollUpdate
	if err := c.ShouldBindJSON(&res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updatedPoll, err := services.EditPollById(parsedObjectId, userID, res.Title)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Poll updated successfully",
		"poll":    updatedPoll, 
	})
}
