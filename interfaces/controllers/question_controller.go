package controllers

import (
	question2 "github.com/c8112002/bbs_api/interfaces/database/question"
	"github.com/c8112002/bbs_api/usecases/question"
	input2 "github.com/c8112002/bbs_api/usecases/question/input"
	"github.com/labstack/echo/v4"
	"net/http"
)

type QuestionController struct {
	QuestionUseCase question.UseCase
}

func NewQuestionController() *QuestionController {
	return &QuestionController{
		QuestionUseCase: question.Interactor{
			QuestionRepository: question2.QuestionRepository{},
		},
	}
}

func (controller *QuestionController) ShowList(c echo.Context) error {

	input := input2.QuestionsFetchInput{}
	output := controller.QuestionUseCase.FetchQuestions(input)

	return c.JSON(http.StatusOK, output.Questions)
}
