package gpt

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/KaueSabinoSRV17/lesson_automation/internal/infra/env"
	"github.com/KaueSabinoSRV17/lesson_automation/internal/questions"
)

type QuestionRequest struct {
	Model    string       `json:"model"`
	Messages []ChatMember `json:"messages"`
}

type QuestionResponse struct {
	Choices []struct {
		Message ChatMember `json:"message"`
	} `json:"choices"`
}

func AskQuestionToChatGpt(question questions.Question) int {
	envs := env.GetEnvs("../../.env", "OPENAI_API_KEY")
	gptClient := DefaultClient(envs["OPENAI_API_KEY"])
	gptClient.ChatMembers[1].Content = question.FormatQuestionAndAlternatives()

	payload := QuestionRequest{
		Model:    Model,
		Messages: gptClient.ChatMembers,
	}
	var result QuestionResponse

	response, err := gptClient.Post("/chat/completions", payload)
	if err != nil {
		log.Fatal("Could not get Completion from ChatGPT\n", err)
		return 0
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		log.Fatal("Could not Unmarshal ChatGPT Answer")
	}

	return CheckWhatIsTheCorrectAnswer(question.Alternatives, result.Choices[0].Message.Content)
}

func CheckWhatIsTheCorrectAnswer(alternatives []questions.Alternative, response string) int {
	for _, alternative := range alternatives {
		formatedId := strconv.Itoa(int(alternative.ID))
		if strings.Contains(response, formatedId) {
			return alternative.ID
		}
	}
	return 0
}
