package gpt

import (
	"testing"

	"github.com/KaueSabinoSRV17/lesson_automation/internal/questions"
)

func TestAskQuestionToChatGpt(t *testing.T) {
	question := questions.Question{
		ID:          123,
		Description: "How Much is 2 plus 2?",
		Alternatives: []questions.Alternative{
			{ID: 1, Description: "1"},
			{ID: 2, Description: "2"},
			{ID: 3, Description: "3"},
			{ID: 4, Description: "4"},
		},
	}
	correctAlternativeId := AskQuestionToChatGpt(question)
	if correctAlternativeId != 4 {
		t.Error("Wrong answer, 2 plus 2 is 4")
	}
}

func TestCheckWhatIsTheCorrectAnswer(t *testing.T) {
	expected := 1
	response := "The correct ID is 1. Description: 4"
	alternatives := []questions.Alternative{
		{ID: 1, Description: "4"},
		{ID: 2, Description: "j"},
		{ID: 3, Description: "6"},
	}
	correctAlternativeId := CheckWhatIsTheCorrectAnswer(alternatives, response)
	if correctAlternativeId != expected {
		t.Error("Should be Correct")
	}
}
