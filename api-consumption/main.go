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

func main() {
	res, err := http.Get("https://www.ahsanzizan.xyz/api/blog")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var blogResponse BlogResponse
	err = json.Unmarshal(body, &blogResponse)
	if err != nil {
		log.Fatal(err)
	}

	// Now you can access the blogs array and prettify it
	prettifiedData, err := json.MarshalIndent(blogResponse.Blogs, " ", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(prettifiedData))
}
