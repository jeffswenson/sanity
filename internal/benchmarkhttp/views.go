package main

import (
	"fmt"

	"github.com/jeffswenson/sanity/pkg/attr"
	"github.com/jeffswenson/sanity/pkg/html"
	"github.com/jeffswenson/sanity/pkg/tag"
)

func indexDocument(articles []ArticleSummary) html.Node {
	return html.Document(
		attr.Lang("en"),
		tag.Head(
			tag.Title(html.InnerText("Sanity News")),
			tag.Link(attr.Rel("stylesheet"), attr.Href("/static/stylesheet.css")),
		),
		tag.Body(
			navigationHeader(),
			html.ForEach(articles, articleView),
		),
	)
}

func articleView(article ArticleSummary) html.Node {
	return tag.Div(
		attr.Class("article"),
		tag.Div(
			attr.Class("article-title-line"),
			tag.A(
				attr.Class("article-name"),
				attr.Href(article.Link),
				html.InnerText(article.ArticleName),
			),
		),
		tag.Div(
			attr.Class("article-subtitle-line"),
			tag.Span(
				attr.Class("article-author"),
				html.InnerText(article.AuthorName),
			),
			tag.Span(html.InnerText(fmt.Sprintf("upvotes: %d", article.UpvoteCount))),
			tag.Span(html.InnerText(article.PostedAt.Format("2006-01-02 15:04:05"))),
		),
	)
}

func navigationHeader() html.Node {
	return tag.Nav(
		attr.Class("navigation"),
		html.InnerText("Sanity News"),
	)
}
