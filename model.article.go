package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{
		ID:      1,
		Title:   "first article title",
		Content: "first article body",
	},
	{
		ID:      2,
		Title:   "second Article title",
		Content: "second article body",
	},
}

func getAllArticles() []article {
	return articleList
}

func getArticleById(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, errors.New("article not found")
}
