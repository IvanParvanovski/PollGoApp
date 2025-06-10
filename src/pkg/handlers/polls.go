package handlers

import (
	"fmt"
	"log"
	"mainapp/pkg/models"
	"mainapp/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListAllPolls(c *gin.Context) {
	polls := services.GetAllPolls()
	
	c.JSON(http.StatusOK, gin.H{"data": polls})
}

func CreateNewPoll(c *gin.Context) {
	var req models.Poll

	if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	poll := &req

	err := services.AddPoll(*poll)
	if err != nil {
		fmt.Println()
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Poll created", "data": req})
}


func DeletePoll(c *gin.Context) {
	id := c.Param("id")

	parsedUUID, _ := uuid.Parse(id)

	services.RemovePollById(parsedUUID)

	c.JSON(http.StatusOK, gin.H{"deleted_poll_id": id})
}

func UpdatePoll(c *gin.Context) {
	id := c.Param("id")
	
	parsedUUID, _ := uuid.Parse(id)
	
	var res models.PollUpdate
	if err := c.ShouldBindJSON(&res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updatedPoll, err := services.EditPollById(parsedUUID, res.Title)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Poll coulndn't be updated"})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Poll updated successfully",
		"poll":    updatedPoll, 
	})
}
