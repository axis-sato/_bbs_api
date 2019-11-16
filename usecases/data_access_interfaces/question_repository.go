package data_access_interfaces

import "github.com/c8112002/bbs_api/domain"

type QuestionRepository interface {
	FetchList() domain.Questions
}