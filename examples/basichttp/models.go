package main

import (
	"fmt"
	"time"
)

// articleSummary is a model representing one article in a list of articles. 
type articleSummary struct {
	articleName string
	authorName string
	link string
	upvoteCout int
	commentCount int
	postedAt time.Time
}

// generateArticles creates example article content. If this was a real
// application, this would read the articles from a database.
func generateArticles(count int) []articleSummary {
	generateArticle := func(id int) articleSummary {
		return articleSummary {
			articleName: fmt.Sprintf("article title %d", id),
			authorName: fmt.Sprintf("author %d", id),
			link: fmt.Sprintf("/article/%d", id),
			upvoteCout: id * 13,
			commentCount: id * 3,
			postedAt: time.Now(),
		}
	}
	articles := make([]articleSummary, count)
	for i := range articles {
		articles[i] = generateArticle(i)
	}
	return articles
}
