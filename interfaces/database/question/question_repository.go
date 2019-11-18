package question

import "github.com/c8112002/bbs_api/domain"

type QuestionRepository struct {

}

func (r QuestionRepository) FetchList() domain.Questions {
	questions := domain.Questions{
		domain.Question{ID: 1, Title: "title1", Body: "body1"},
		domain.Question{ID: 2, Title: "title2", Body: "body2"},
	}
	return questions
}
