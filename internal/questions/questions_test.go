package questions

import (
	"strings"
	"testing"

	"github.com/KaueSabinoSRV17/lesson_automation/internal/studeoapi"
)

// Questionary ID: 227751

func TestGetQuestionsFromQuestionary(t *testing.T) {
	api, err := studeoapi.AuthenticatedInstance()
	if err != nil {
		t.Error("Could not authenticate")
	}
	questions := GetQuestionsFromQuestionary(api)
	if len(questions) == 0 {
		t.Error("Should not be an empty list")
	}
}

func TestFormatQuestionAndAlternatives(t *testing.T) {
	question := Question{
		Description: "How much is 2 plus 2?",
		Alternatives: []Alternative{
			{ID: 1, Description: "4"},
		},
	}
	fullString := question.FormatQuestionAndAlternatives()
	if !strings.Contains(fullString, "4") {
		t.Error("Should Contain 4, an Alternative")
	}
}
