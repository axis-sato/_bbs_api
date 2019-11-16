package question

import (
	"github.com/c8112002/bbs_api/usecases/data_access_interfaces"
	"github.com/c8112002/bbs_api/usecases/question/input"
	"github.com/c8112002/bbs_api/usecases/question/output"
)

type Interactor struct {
	QuestionRepository data_access_interfaces.QuestionRepository
}

func (interactor Interactor) FetchQuestions(input input.QuestionsFetchInput) output.QuestionsFetchOutput {
	questions := interactor.QuestionRepository.FetchList()
	return output.QuestionsFetchOutput{Questions: questions}
}
