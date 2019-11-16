package question

import (
	"github.com/c8112002/bbs_api/usecases/question/input"
	"github.com/c8112002/bbs_api/usecases/question/output"
)

type UseCase interface {

	FetchQuestions(input input.QuestionsFetchInput) output.QuestionsFetchOutput
}
