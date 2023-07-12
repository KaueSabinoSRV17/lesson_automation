package questions

import (
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
