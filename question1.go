

package main
import (
	"fmt"
	"net/http"
    "io/ioutil"
       "time"
       "strconv"
	   "regexp"
    "strings"
)

func fetch (url string) string {
    // fmt.Println("Fetch Url", url)
    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Http get err:", err)
        return ""
    }
    if resp.StatusCode != 200 {
        fmt.Println("Http status code:", resp.StatusCode)
        return ""
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Read error", err)
        return ""
    }
    return string(body)
}

func parseUrls(url string) {
    body := fetch(url)
    body = strings.Replace(body, "n", "", -1)
    rp := regexp.MustCompile(`"email":(.*?),`)
    items := rp.FindAllStringSubmatch(body, -1)
    for _, item := range items {
        fmt.Println((item[1]))
    }
}

func query (start int, end int) {
	// func query (i int) {
	for i := start; i <= end; i++ {
		parseUrls("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i) + "/comments")
	}
}

func main() {
        start := time.Now()

		for i := 1; i <= 9; i++  {
			go query((i - 1) * 10 + 1, i+9)
		}
		query(91, 100)
        elapsed := time.Since(start)
        fmt.Printf("Took %s", elapsed)
}
