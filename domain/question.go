package domain

type Question struct {
	ID int
	Title string
	Body string
}

type Questions = []Question 
