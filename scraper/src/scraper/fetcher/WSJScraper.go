package fetcher

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html"
)

// WSJ new source types

type WSJArticle struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`

	// article body
	Data string
}

// TODO: add errors
func (article WSJArticle) GetLink() string        { return article.Link }
func (article WSJArticle) GetDescription() string { return article.Description }
func (article WSJArticle) GetTitle() string       { return article.Title }
func (article WSJArticle) GetData() string        { return article.Data }

// use ptrs for the next two because we want the article changed
func (article *WSJArticle) SetData(data string) { article.Data = data }
func (article *WSJArticle) DoParse(parser *html.Tokenizer) error {

articleTagLoop:
	for {
		token := parser.Next()

		switch {
		case token == html.ErrorToken:
			fmt.Println("OH NOSE!!!! ERROR before we hit the end")
			return nil
		case token == html.StartTagToken:
			tmp := parser.Token()

			isStartArticle := tmp.Data == "article"
			if isStartArticle {
				fmt.Println("found start of article")
				break articleTagLoop
			}
		}
	}

	// loop until we start to hit article tokens
articleStartLoop:
	for {
		token := parser.Next()

		switch {
		case token == html.ErrorToken:
			return nil
		case token == html.StartTagToken:
			tmp := parser.Token()

			isStartArticleBody := tmp.Data == "div"
			// loop until we are at the first paragraph of the article body
			if isStartArticleBody {
				isStartArticleBody = false
				for _, attr := range tmp.Attr {
					if attr.Key == "class" && attr.Val == "clearfix byline-wrap" {
						isStartArticleBody = true
						break
					}
				}
				if isStartArticleBody {
					break articleStartLoop
				}
			}
		}
	}

	// now loop until the body is at the first paragraph tag
articleBodyStartLoop:
	for {
		token := parser.Next()
		switch {
		case token == html.ErrorToken:
			return nil
		case token == html.StartTagToken:
			tmp := parser.Token()
			isStartArticleBody := tmp.Data == "p"
			if isStartArticleBody {
				break articleBodyStartLoop
			}
		}
	}

	// loop until article is all the way grabbed
	addParagraph := true
	for {
		token := parser.Next()
		switch {
		case token == html.ErrorToken:
			return nil
		case token == html.StartTagToken:
			tmp := parser.Token()
			isParagraph := tmp.Data == "p"
			if isParagraph {
				parser.Next()
				tmp = parser.Token()
				newBody := article.GetData()
				if addParagraph {
					newBody = newBody + "\n" + tmp.Data
				} else {
					addParagraph = true
					newBody = newBody + tmp.Data
				}

				article.SetData(newBody)
			}

			isLink := tmp.Data == "a"
			if isLink {
				parser.Next()
				tmp = parser.Token()
				newBody := article.GetData() + tmp.Data
				article.SetData(newBody)
				// TODO: check if this is ever at the end of a paragraph....
				addParagraph = false
			}
		case token == html.EndTagToken:
			isEnd := parser.Token().Data == "div"
			if isEnd {
				return nil
			}
		}
	}
	return nil
}

type WSJRSSChannel struct {
	XMLName  xml.Name     `xml:"channel"`
	Articles []WSJArticle `xml:"item"`
}

func (channel *WSJRSSChannel) GetArticle(slot int) Article {
	if slot >= channel.GetNumArticles() {
		// Check that the request doesn't go out of bounds
		// TODO: errors
		return nil
	}
	return &channel.Articles[slot]
}

func (channel *WSJRSSChannel) GetNumArticles() int {
	return len(channel.Articles)
}

type WSJRSS struct {
	XMLName xml.Name      `xml:"rss"`
	Channel WSJRSSChannel `xml:"channel"`
	RSSLink string
	// TODO: actually set string to the value of the link
}

func (rss *WSJRSS) GetLink() string { return rss.RSSLink }

func (rss *WSJRSS) GetChannel() RSSChannel {
	// return a pointer to the channel, interfaces implicitly have ptrs if they are there
	tmp := &rss.Channel
	return tmp
}

// make sure all the structs implement the interfaces
var _ RSS = (*WSJRSS)(nil)
var _ RSSChannel = (*WSJRSSChannel)(nil)
var _ Article = (*WSJArticle)(nil)
