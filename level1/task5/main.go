package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	err := os.MkdirAll("./storage/posts", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}

	for i := 1; i < 101; i++ {
		wg.Add(1)
		go func(post int, url string) {
			GetFiles("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(post),
				url + strconv.Itoa(post))
			wg.Done()
		}(i, "./storage/posts/post")
	}
	wg.Wait()
}

func GetFiles(url string, filePath string) {
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

	err = ioutil.WriteFile(filePath, body, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("<%s> - Status Code: [%d] - Latency: %d ms\n\n",
		url, resp.StatusCode, time.Since(start).Milliseconds())
}
