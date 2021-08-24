package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func main() {
	db, err := sql.Open("mysql",
		"root:1715rjvxbr7410@tcp(127.0.0.1:6603)/task6")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	posts := GetPosts("https://jsonplaceholder.typicode.com/posts?userId=7")

	wgReq := sync.WaitGroup{}
	for _, post := range posts {
		if _, err := db.Query("INSERT INTO task6.posts VALUES(?, ?, ?, ?)",
			post.UserID,
			post.ID,
			post.Title,
			post.Body); err != nil {
			log.Fatal(err)
		}

		wgReq.Add(1)
		go func(post *Post) {
			comments := GetComments("https://jsonplaceholder.typicode.com/comments?postId=" + strconv.Itoa(post.ID))
			wgComm := sync.WaitGroup{}
			for _, comment := range comments {
				wgComm.Add(1)
				go func(comment *Comment) {
					if _, err := db.Query("INSERT INTO task6.comments VALUES(?, ?, ?, ?, ?)",
						comment.PostID,
						comment.ID,
						comment.Name,
						comment.Email,
						comment.Body); err != nil {
						log.Fatal(err)
					}
					wgComm.Done()
				}(comment)
			}
			wgComm.Wait()
			wgReq.Done()
		}(post)
	}
	wgReq.Wait()
}

func GetPosts(url string) []*Post {
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

	posts := make([]*Post, 0)
	if err := json.Unmarshal(body, &posts); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("<%s> - Status Code: [%d] - Latency: %d ms\n",
		url, resp.StatusCode, time.Since(start).Milliseconds())

	return posts
}

func GetComments(url string) []*Comment {
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

	comments := make([]*Comment, 0)
	if err := json.Unmarshal(body, &comments); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("<%s> - Status Code: [%d] - Latency: %d ms\n",
		url, resp.StatusCode, time.Since(start).Milliseconds())

	return comments
}
