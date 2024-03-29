package questions

import (
	"fmt"
	"strings"
	"time"

	"github.com/KaueSabinoSRV17/lesson_automation/internal/studeoapi"
)

type Alternative struct {
	ID          int    `json:"idAlternativa"`
	Description string `json:"descricao"`
}

type Question struct {
	ID           int           `json:"idQuestao"`
	Description  string        `json:"descricaoTexto"`
	Alternatives []Alternative `json:"alternativaList"`
}

func (q *Question) FormatQuestionAndAlternatives() string {
	question := fmt.Sprintf("Tenho uma pergunta e algumas alternativas. Com base nelas, me diga qual o ID da alternativa correta %s", q.Description)
	var alternatives []string
	alternatives = append(alternatives, question)
	for _, alternative := range q.Alternatives {
		alternatives = append(alternatives, fmt.Sprintf("ID: %v, Descrição: %v", alternative.ID, alternative.Description))
	}
	fullString := strings.Join(alternatives, "\n")
	return fullString
}

func GetQuestionsFromQuestionary(api studeoapi.Client) []Question {
	path := "/objeto-ensino-api-controller/api/questao/shortname/2023_EGRAD_ADSIS6GA-52_EGRAD_NGER100_053/questionario/227751"
	var questions []Question
	api.Get(path, nil, &questions)
	return questions
}

func GetAnswersFromQuestionsList(questions []Question) []int {
	answers := make([]int, len(questions))
	channel := make(chan int)
	for index, question := range questions {
		go func(q Question, i int) {
			answer := MockGPTQuestion(q, i)
			channel <- answer
		}(question, index)
	}

	for index := range questions {
		answer := <-channel
		answers[index] = answer
	}

	return answers
}

func MockGPTQuestion(question Question, index int) int {
	time.Sleep(time.Second * 2)
	fmt.Printf("\nAnswer to %v is DeezNuttzz. Number %v\n", question.Description, index)
	return 0
}
