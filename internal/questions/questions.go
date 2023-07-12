package questions

import (
	"github.com/KaueSabinoSRV17/lesson_automation/internal/studeoapi"
)

type Alternative struct {
	ID          uint64 `json:"idAlternativa"`
	Description string `json:"descricao"`
}

type Question struct {
	ID           uint64        `json:"idQuestao"`
	Description  string        `json:"descricaoTexto"`
	Alternatives []Alternative `json:"alternativaList"`
}

func GetQuestionsFromQuestionary(api studeoapi.Client) []Question {
	path := "/objeto-ensino-api-controller/api/questao/shortname/2023_EGRAD_ADSIS6GA-52_EGRAD_NGER100_053/questionario/227751"
	var questions []Question
	api.Get(path, nil, &questions)
	return questions
}
