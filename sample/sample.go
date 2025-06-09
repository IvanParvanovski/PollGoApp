package main 
// import (
// 	"fmt"
// )

// type Poll struct {
// 	Id int
// 	Title string
// 	Questions []QuestionInterface
// }

// type QuestionInterface interface {
// 	GetQuestion() Question
// }
// func (m MultipleChoice) GetQuestion() Question {
// 	return m.Question
// }
// func (s SingleChoice) GetQuestion() Question {
// 	return s.Question
// }

// type Question struct {
// 	Id int
// 	Description string 
// 	Required bool 
// 	Answers []Answer
// }
// type MultipleChoice struct {
// 	Question

// 	MinSelection int
// 	MaxSelection int
// 	UserAnswers []int
// }
// type SingleChoice struct {
// 	Question

// 	UserAnswer int
// }
// type Answer struct {
// 	Id int 
// 	Description string 
// 	Count int
// }
// type Result struct {
// 	Id int
// 	SubmittedAnswers []Question // copy of the poll questions when submitted
// }


// func main() {
// 	vegetablesPoll := setupVegetables()
// 	answerQuestions(vegetablesPoll)
// 	displayPoll(vegetablesPoll)
// }

// func answerQuestions(vegetablesPoll Poll) {

// 	var currentQuestion QuestionInterface
	
// 	currentQuestion = vegetablesPoll.Questions[0]
// 	mc, ok := currentQuestion.(MultipleChoice)
// 	if !ok {
// 		fmt.Println("First question is not a MultipleChoice")
// 		return
// 	}
// 	mc.UserAnswers= []int{1, 3}
// 	vegetablesPoll.Questions[0] = mc

// 	currentQuestion = vegetablesPoll.Questions[1]
// 	sc, ok := currentQuestion.(SingleChoice)
// 	if !ok {
// 		fmt.Println("The second question is not a SingleChoice")
// 	}
// 	sc.UserAnswer = 1

// }

// func setupVegetables() Poll {
// 	vegetablesPoll := Poll{
// 		Id: 1,
// 		Title: "My Vegetables poll",
// 	}

// 	a1 := Answer{Id: 1, Description: "Cucumber", Count: 0}
// 	a2 := Answer{Id: 2, Description: "Carrot", Count: 0}
// 	a3 := Answer{Id: 3, Description: "Potato", Count: 0}
// 	a4 := Answer{Id: 4, Description: "Corn", Count: 0}
// 	baseQuestion1 := Question{
// 		Id: 1,
// 		Description: "Select your two favorite vegetables from the list below.",
// 		Required: false,
// 		Answers: []Answer{a1, a2, a3, a4},
// 	}
// 	q1 := MultipleChoice{
// 		Question: baseQuestion1,
// 		MinSelection: 1,
// 		MaxSelection: 2,
// 	}

// 	a5 := Answer{Id: 5, Description: "Ginger", Count: 0}
// 	a6 := Answer{Id: 6, Description: "Spinach", Count: 0}
// 	a7 := Answer{Id: 7, Description: "Cabbage", Count: 0}
// 	a8 := Answer{Id: 8, Description: "Garlic", Count: 0}
// 	baseQuestion2 := Question{
// 		Id: 2,
// 		Description: "Which vegetable do you like the most?",
// 		Required: true,
// 		Answers: []Answer{a5, a6, a7, a8},
// 	}
// 	q2 := SingleChoice{
// 		Question: baseQuestion2,
// 	}

// 	vegetablesPoll.Questions = append(vegetablesPoll.Questions, q1)
// 	vegetablesPoll.Questions = append(vegetablesPoll.Questions, q2)

// 	return vegetablesPoll
// }

// func displayPoll(poll Poll) {
// 	fmt.Printf("Poll ID: %d\nTitle: %s\n", poll.Id, poll.Title)
// 	for i, q := range poll.Questions {
// 		fmt.Printf("\nQuestion %d:\n", i+1)

// 		switch v := q.(type) {
// 		case MultipleChoice:
// 			fmt.Printf("Type: Multiple Choice\n")
// 			fmt.Printf("ID: %d\nDescription: %s\nRequired: %t\n", v.Id, v.Description, v.Required)
// 			fmt.Printf("Min: %d, Max: %d\n", v.MinSelection, v.MaxSelection)

// 			for _, a := range v.Answers {
// 				selected := ""
// 				for _, userID := range v.UserAnswers {
// 					if userID == a.Id {
// 						selected = " ← selected"
// 						break
// 					}
// 				}
// 				fmt.Printf("  - [%d] %s (Count: %d)%s\n", a.Id, a.Description, a.Count, selected)
// 			}

// 		case SingleChoice:
// 			fmt.Printf("Type: Single Choice\n")
// 			fmt.Printf("ID: %d\nDescription: %s\nRequired: %t\n", v.Id, v.Description, v.Required)

// 			for idx, a := range v.Answers {
// 				selected := ""
// 				if v.UserAnswer == idx {
// 					selected = " ← selected"
// 				}
// 				fmt.Printf("  - [%d] %s (Count: %d)%s\n", a.Id, a.Description, a.Count, selected)
// 			}

// 		default:
// 			fmt.Println("Unknown question type")
// 		}
// 	}
// }