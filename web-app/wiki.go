package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title       string
	Description string
	Body        []byte
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-z0-9]+)$")

const PORT string = "8080"

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Page title not found")
	}
	return m[2], nil
}

func createHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, m[2])
	}
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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Handlers
func handleView(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Println("Handling View localhost" + PORT)
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
		return
	}

	renderTemplate(w, "view", p)
}

func handleEdit(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Println("Handling Edit localhost" + PORT)
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func handleSave(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Println("Handling Save localhost" + PORT)
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusNotFound)
}

func main() {
	http.HandleFunc("/view/", createHandler(handleView))
	http.HandleFunc("/edit/", createHandler(handleEdit))
	http.HandleFunc("/save/", createHandler(handleSave))
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
