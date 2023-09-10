package main

import (
	"fmt"

	"github.com/jeffswenson/sanity/pkg/attr"
	"github.com/jeffswenson/sanity/pkg/html"
	"github.com/jeffswenson/sanity/pkg/tag"
)

// indexDocument is the view for the '/' url.
func indexDocument(articles []articleSummary) html.Node {
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

// articleView is a helper for indexDocument. It accepts a model and produces
// html that displays the content of the model. 
func articleView(article articleSummary) html.Node {
	return tag.Div(
		attr.Class("article"),
		tag.Div(
			attr.Class("article-title-line"),
			tag.A(
				attr.Class("article-name"),
				attr.Href(article.link),
				html.InnerText(article.articleName),
			),
		),
		tag.Div(
			attr.Class("article-subtitle-line"),
			tag.Span(
				attr.Class("article-author"),
				html.InnerText(article.authorName),
			),
			tag.Span(html.InnerText(fmt.Sprintf("upvotes: %d", article.upvoteCout))),
			tag.Span(html.InnerText(article.postedAt.Format("2006-01-02 15:04:05"))),
		),
	)
}

// navigationHeader renders a static html node that contains the site's title
// and navigation menu.
func navigationHeader() html.Node {
	// TODO put navigation elements in this.
	return tag.Nav(
		attr.Class("navigation"),
		html.InnerText("Sanity News"),
	)
}
