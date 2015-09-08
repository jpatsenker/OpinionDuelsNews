package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"golang.org/x/net/html"
)


// Parse the wsj opinion rss feed into useable links.
// Returns an array of articles and an error
func GetStories(rss RSS, body []byte) (error) {
	err := xml.Unmarshal(body, &rss)
	if err != nil {
		fmt.Printf("err:",err)
		return err
	}

	tmpChan := rss.GetChannel()
	for _, article := range tmpChan.GetArticles() {
		fmt.Println("title:", article.GetInfo().GetTitle(),"\tdescr:", article.GetInfo().GetDescription())
	}	

	return nil
}


// Request a page containing the article linked to
func GetArticle(article Article) (Article, error){
	client := &http.Client{}

	req, err := http.NewRequest("GET", article.GetInfo().GetLink(), nil)
	req.Header.Add("Referer", "https://www.google.com")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("oh nose, err with get article http request")
		return article, err
	}
	
	defer resp.Body.Close()	
	parser := html.NewTokenizer(resp.Body)
	article.DoParse(parser)
	return article, err
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
	rss := WSJRSS{}
	err = GetStories(rss, body)
	if err != nil {
		fmt.Println("oh nose, error working with body")
		return
	}
	articles := rss.GetChannel().GetArticles()
	article, err :=	GetArticle(articles[0])
	fmt.Println("article body is:", article.GetData())	
}
