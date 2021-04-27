package models

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Desc    string `json:"description"`
}
