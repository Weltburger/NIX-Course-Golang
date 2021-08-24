package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}*/

func main() {
	MakeRequest("https://jsonplaceholder.typicode.com/comments?postId=7")
}

func MakeRequest(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	/*post := make([]Post, 0)
	if err := json.Unmarshal(body, &post); err != nil {
		log.Fatal(err)
	}
	fmt.Println(post)*/

	fmt.Printf("<%s> - Status Code: [%d] - Latency: %d ms",
		url, resp.StatusCode, time.Since(start).Milliseconds())
}
