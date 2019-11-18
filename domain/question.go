package domain

type Question struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Category Category `json:"category"`
}

type Questions = []Question 
