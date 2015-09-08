package main 

import (
	"golang.org/x/net/html"
	"encoding/xml"
	"fmt"
)

type WSJArticle struct {
	
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`

	Data string
}


func (article WSJArticle) GetLink() string { return article.Link }
func (article WSJArticle) GetDescription() string { return article.Description }
func (article WSJArticle) GetTitle() string { return article.Title }

func (article WSJArticle) GetData() string { return article.Data }
func (article *WSJArticle) SetData(data string) { article.Data = data }

func (article *WSJArticle) DoParse(parser *html.Tokenizer) error {
	//parser := html.NewTokenizer(body.Body)
	i := 0
	for {
		i++
		token := parser.Next()
		
		switch {
			case token == html.ErrorToken :
				fmt.Println("found error token on first parser:", token, "\tloop is:", i)
				fmt.Println("actual token is:", parser.Token())
				return nil
			case token == html.StartTagToken:
				tmp := parser.Token()

				isStartArticle := tmp.Data == "article"
				if isStartArticle {
					fmt.Println("found start of article")
					// loop until we start to hit article tokens
					for {
						token = parser.Next()

						switch {
							case token == html.ErrorToken:
								return nil
							case token == html.StartTagToken:
								tmp = parser.Token()
								
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
										fmt.Println("hit inside the clearfix")
										// now loop until the body is at the first paragraph tag
										for {
											token = parser.Next()
											tmp = parser.Token()
											
											switch {
												case token == html.ErrorToken:
													return nil
												case token == html.StartTagToken:
													isStartArticleBody = tmp.Data == "p"
													if isStartArticleBody {
														parser.Next()
														tmp = parser.Token()	
														article.SetData(tmp.Data)
														return nil
													}
												case token == html.EndTagToken:
													break
											}
										}
										break // from the top for loop

									}
								}
							}
						}
				}
		}
	}
	return nil
}

type WSJRSSChannel struct {
	XMLName xml.Name `xml:"channel"`
	Articles []WSJArticle `xml:"item"`
}

func (channel *WSJRSSChannel) GetArticle(slot int) Article {
	if slot >= channel.GetNumArticles() {
		return nil
	}
	return &channel.Articles[slot]
}

func (channel *WSJRSSChannel) GetNumArticles() int {
	return len(channel.Articles)
}

type WSJRSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel WSJRSSChannel `xml:"channel"`
	RSSLink string
}

func (rss *WSJRSS) GetLink() string { return rss.RSSLink }

func (rss *WSJRSS) GetChannel() RSSChannel { 
	tmp := &rss.Channel
	return tmp 
}

// make sure all the structs implement the interfaces
var _ RSS = (*WSJRSS)(nil)
var _ RSSChannel = (*WSJRSSChannel)(nil)
var _ Article = (*WSJArticle)(nil)
