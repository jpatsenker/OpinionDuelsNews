package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
)

func GetStories(body string) ([]string, error) {

}
func main(){
	fmt.Println("hello world 3")

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
	fmt.Println("body is:", string(body))

}
