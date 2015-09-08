package main 

import (
	"golang.org/x/net/html"
	"encoding/xml"
	"fmt"
)

type WSJArticleInfo struct {
	XMLName xml.Name `xml:"item"`
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
}

func (article WSJArticleInfo) GetLink() string { return article.Link }
func (article WSJArticleInfo) GetDescription() string { return article.Description }
func (article WSJArticleInfo) GetTitle() string { return article.Title }

type WSJArticle struct {
	Info WSJArticleInfo
	Data string
}

func (article WSJArticle) GetInfo() ArticleInfo { 
	//tmp := ArticleInfo.(article.Info)
	return article.Info 
}

func (article *WSJArticle) GetData() string { return article.Data }
func (article *WSJArticle) SetData(data string) { article.Data = data }

func (article WSJArticle) DoParse(parser *html.Tokenizer) error {
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

func (channel WSJRSSChannel) GetArticles() []Article { 
	tmp := []Article(channel.Articles)
	return tmp //channel.Articles
}

type WSJRSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel WSJRSSChannel `xml:"channel"`
	RSSLink string
}

func (rss WSJRSS) GetLink() string { return rss.RSSLink }

func (rss WSJRSS) GetChannel() RSSChannel { 
	tmp := rss.Channel
	return tmp 
}
