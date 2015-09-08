package main
import (
	"fmt"
    "os"
    "os/exec"
	// "net/http"
	// "net/url"
	// "bytes"
	// "io/ioutil"
	// "strings"

	)
func main() {
	// apiUrl := "http://access.alchemyapi.com/calls/text/"
 //    resource := "/TextGetRankedKeywords/"
 //    data := url.Values{}
 //    data.Set("apikey", "39995101e65858870797a627e548b1522f5c74a8")
 //    data.Add("text", "this is an example")
 //    data.Add("sentiment", "1")

 //    u, _ := url.ParseRequestURI(apiUrl)
 //    u.Path = resource
 //    urlStr := fmt.Sprintf("%v", u)
    
 //    client := &http.Client{}
 //    r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload
 //    // r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
 //    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
 //    // r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

 //    resp, _ := client.Do(r)
 //    fmt.Println(resp.Status)
    cmd := "curl"
    args := []string{"--data", "apikey=39995101e65858870797a627e548b1522f5c74a8&text=hello%20my%20name%20is%20test","http://access.alchemyapi.com/calls/text/TextGetRankedKeywords"}
    out, err := exec.Command(cmd, args...).Output();
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    fmt.Println(string(out))

}

// key 39995101e65858870797a627e548b1522f5c74a8
// curl --data "apikey=39995101e65858870797a627e548b1522f5c74a8&text=hello%20my%20name%20is%20test" http://access.alchemyapi.com/calls/text/TextGetRankedKeywords