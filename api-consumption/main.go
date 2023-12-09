package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Blog struct {
	ID        string   `json:"_id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	CreatedAt string   `json:"createdAt"`
	Link      string   `json:"link"`
	Author    string   `json:"author"`
	Tags      []string `json:"tags"`
}

type BlogResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Blogs   []Blog `json:"blogs"`
}

const API_URL string = "https://www.ahsanzizan.xyz/api/blog"

func print(message string) {
	fmt.Println(message)
}

func printErr(err error) {
	log.Fatal(err)
}

func main() {
	res, err := http.Get(API_URL)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		printErr(err)
	}

	var blogResponse BlogResponse
	err = json.Unmarshal(body, &blogResponse)
	if err != nil {
		printErr(err)
	}

	prettifiedData, err := json.MarshalIndent(blogResponse.Blogs, " ", "    ")
	if err != nil {
		printErr(err)
	}
	print(string(prettifiedData))
}
