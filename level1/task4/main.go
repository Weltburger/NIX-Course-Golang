package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 1; i < 101; i++ {
		wg.Add(1)
		go func(post int) {
			MakeRequest("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(post))
			wg.Done()
		}(i)
	}
	wg.Wait()
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
	fmt.Printf("<%s> - Status Code: [%d] - Latency: %d ms\n\n",
		url, resp.StatusCode, time.Since(start).Milliseconds())
}
