package services

import (
	"mainapp/pkg/models"
	"errors"
	"github.com/google/uuid"
)

var allPolls []models.Poll


// Poll retrival functions 

func GetAllPolls() []models.Poll {
	return allPolls
}

func findPollIndexById(id uuid.UUID) int {
	for i, poll := range allPolls {
		if poll.Id == id {
			return i
		}
	}

	return -1
}

// Poll creation and modification

func AddPoll(poll models.Poll) error {
	index := findPollIndexById(poll.Id)
	
	if index != -1 {
		return errors.New("already exists")
	}

	allPolls = append(allPolls, poll)	
	return nil
}

func RemovePollById(id uuid.UUID) error {
	pollIndex := findPollIndexById(id)
	
	if pollIndex == -1 {
		return errors.New("poll not found")
	}

	allPolls = append(allPolls[:pollIndex], allPolls[pollIndex + 1:]...)
	return nil
}

func EditPollById(id uuid.UUID, title string) (models.Poll, error) {
	index := findPollIndexById(id)

	if index == -1 {
		return models.Poll{}, errors.New("poll not found")
	}

	currentPoll := &allPolls[index]
	currentPoll.Title = title

	return *currentPoll, nil 
}

