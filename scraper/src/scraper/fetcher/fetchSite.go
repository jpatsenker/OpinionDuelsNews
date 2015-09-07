package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Oh nose, error reading body in get article")
		return "", err
	}

	fmt.Println("article is:", string(body))
	return "", nil
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

	GetArticle(articles[0])
}
