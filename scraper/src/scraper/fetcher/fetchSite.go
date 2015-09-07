package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"golang.org/x/net/html"
)


type Article struct {
		XMLName xml.Name `xml:"item"`
		Title string `xml:"title"`
		Link string `xml:"link"`
		Description string `xml:"description"`
	}

// Parse the wsj opinion rss feed into useable links.
// Returns an array of articles and an error
func GetStories(body []byte) ([]Article, error) {
	type WSJ struct {
		XMLName xml.Name `xml:"channel"`
		Language string `xml:"language"`
		Articles []Article `xml:"item"`
	}

	type RSS struct {
		XMLName xml.Name `xml:"rss"`
		Version float32 `xml:"version,attr"`
		WSJ WSJ `xml:"channel"` 
	}

	rss := RSS{Version:1.0, WSJ:WSJ{Language:"none"}}

	err := xml.Unmarshal(body, &rss)
	if err != nil {
		fmt.Printf("err:",err)
		return nil, err
	}

	for _, article := range rss.WSJ.Articles {
		fmt.Println("title:", article.Title,"\tdescr:", article.Description)
	}	

	return rss.WSJ.Articles, nil
}

// Get article body from an http page
func ExtractArticle(body *http.Response) (string, error){
	parser := html.NewTokenizer(body.Body)
	article := "done"
	i := 0
	for {
		i++
		token := parser.Next()
		
		switch {
			case token == html.ErrorToken :
				fmt.Println("found error token on first parser:", token, "\tloop is:", i)
				fmt.Println("actual token is:", parser.Token())
				return "", nil
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
								return "", nil
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
													return "", nil
												case token == html.StartTagToken:
													isStartArticleBody = tmp.Data == "p"
													if isStartArticleBody {
														parser.Next()
														tmp = parser.Token()	
														article = article + tmp.Data
														return article, nil
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
	return article, nil
}


// Request a page containing the article linked to
func GetArticle(article Article) (string, error){
	client := &http.Client{}

	req, err := http.NewRequest("GET", article.Link, nil)
	req.Header.Add("Referer", "https://www.google.com")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("oh nose, err with get article http request")
		return "", err
	}
	
	defer resp.Body.Close()	
	articleBody, err := ExtractArticle(resp)
	return articleBody, err
}

func main(){
	// do a simple http fetch:
	resp, err := http.Get("http://www.wsj.com/xml/rss/3_7041.xml")
	if err != nil {
		fmt.Println("OH NOSE: got an error when trying to fetch the datz:", err)
		return
	}

	// make sure the body gets closed laster
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Oh nose: error reading body:", err)
		return 
	}

	articles, err := GetStories(body)
	if err != nil {
		fmt.Println("oh nose, error working with body")
		return
	}

	articleBody, err :=	GetArticle(articles[0])
	fmt.Println("article body is:", articleBody)	
}
