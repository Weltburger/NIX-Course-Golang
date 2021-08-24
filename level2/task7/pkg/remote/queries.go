package remote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"task7/models"
	"time"
)

func GetPosts(url string) []*models.Post {
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

	posts := make([]*models.Post, 0)
	if err := json.Unmarshal(body, &posts); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("<%s> - Status Code: [%d] - Latency: %d ms\n",
		url, resp.StatusCode, time.Since(start).Milliseconds())

	return posts
}

func GetComments(url string) []*models.Comment {
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

	comments := make([]*models.Comment, 0)
	if err := json.Unmarshal(body, &comments); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("<%s> - Status Code: [%d] - Latency: %d ms\n",
		url, resp.StatusCode, time.Since(start).Milliseconds())

	return comments
}
