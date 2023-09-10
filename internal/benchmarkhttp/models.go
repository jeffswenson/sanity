package main

import (
	"fmt"
	"time"
)

type ArticleSummary struct {
	ArticleName string
	AuthorName string
	Link string
	UpvoteCount int
	CommentCount int
	PostedAt time.Time
}

func generateArticles(count int) []ArticleSummary {
	generateArticle := func(id int) ArticleSummary {
		return ArticleSummary {
			ArticleName: fmt.Sprintf("article title %d", id),
			AuthorName: fmt.Sprintf("author %d", id),
			Link: fmt.Sprintf("/article/%d", id),
			UpvoteCount: id * 13,
			CommentCount: id * 3,
			PostedAt: time.Now(),
		}
	}
	articles := make([]ArticleSummary, count)
	for i := range articles {
		articles[i] = generateArticle(i)
	}
	return articles
}


