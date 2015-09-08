package main

import(
	"golang.org/x/net/html"
)

type _ArticleInfo interface {
	GetLink() string
	GetDescription() string
	GetTitle() string
}

// Generic article type
// Gets body from an article link
type _Article interface {
	GetInfo() _ArticleInfo
	SetInfo(_ArticleInfo) 

	DoParse(html.Tokenizer) error

	SetData(string)
	GetData(string)
}

// Generic RSS feed
// TODO: have this provide a "get new articles"
type _RSS interface {
	GetLink() string
	GetArticles() []_Article
	GetNewArticles([]_Article) []_Article 
}

