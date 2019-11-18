package question

import "github.com/c8112002/bbs_api/domain"

type QuestionRepository struct {

}

func (r QuestionRepository) FetchList() domain.Questions {
	questions := domain.Questions{
		domain.Question{ID: 1, Title: "title1", Body: "body1", Category: domain.Category{ID: 1, Name: "カテゴリ1"}},
		domain.Question{ID: 2, Title: "title2", Body: "body2", Category: domain.Category{ID: 2, Name: "カテゴリ2"}},
	}
	return questions
}
