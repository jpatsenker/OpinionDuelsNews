package main

import(
	//"golang.org/x/net/html"
)



// Generic article type
// Gets body from an article link
type Article interface {

	//DoParse(*html.Tokenizer) error

	SetData(string)
	GetData() string

	GetLink() string
	GetDescription() string
	GetTitle() string
}

// Generic RSS feed
// TODO: have this provide a "get new articles"
type RSS interface {
	GetLink() string
	GetChannel() RSSChannel
}

type RSSChannel interface {
	GetArticle(int) Article
	GetNumArticles() int
}
