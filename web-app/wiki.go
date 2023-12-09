package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title       string
	Description string
	Body        []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	// Octal integer 0600 : file should be created with read-write permissions
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	page := &Page{Title: "test", Body: []byte("This is a sample page.")}
	page.save()
	load, err := loadPage("test")
	if err != nil {
		fmt.Println("An error occurred while loading")
	}
	fmt.Println(string(load.Body))
}
