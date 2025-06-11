package services

import (
	"mainapp/pkg/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var votes []models.UserResponse

func GetAllVotes() []models.UserResponse  {
	return votes
}

func GetPollVotes(pollId primitive.ObjectID) []models.UserResponse {
	var pollVotes []models.UserResponse

	for i, v := range votes {
		if v.PollId == pollId {
			pollVotes = append(pollVotes, votes[i])
		}
	}

	return pollVotes
}

func AddVote(vote models.UserResponse) error {
	votes = append(votes, vote)
	return nil
}

func getVoteIndex(id primitive.ObjectID) int {
	for i, vote := range votes {
		if vote.Id == id {
			return i
		}
	}
	return -1
}

func RemoveVoteById(id primitive.ObjectID) {
	voteIndex := getVoteIndex(id)

	if voteIndex == -1 {
		return 
	}

	votes = append(votes[:voteIndex], votes[voteIndex + 1:]...)
}
