package main

import(
	"golang.org/x/net/html"
)

type ArticleInfo interface {
	GetLink() string
	GetDescription() string
	GetTitle() string
}

// Generic article type
// Gets body from an article link
type Article interface {
	GetInfo() ArticleInfo

	DoParse(*html.Tokenizer) error

	SetData(string)
	GetData() string
}

// Generic RSS feed
// TODO: have this provide a "get new articles"
type RSS interface {
	GetLink() string
	GetChannel() RSSChannel
}

type RSSChannel interface {
	GetArticles() []Article
}
