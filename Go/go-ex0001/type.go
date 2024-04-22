package main

type Article struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Body   string  `json:"body"`
	Author string  `json:"author"`
	// Image  float64 `json:"image"`
}

// articles slice to seed record article data.
var articles = []Article{
	{ID: "1", Title: "Blue Train", Body: "Blue Train body string text", Author: "1"},
	{ID: "2", Title: "Green Train", Body: "Green Train body string text", Author: "2"},
	{ID: "3", Title: "Red Train", Body: "Red Train body string text", Author: "3"},
}
