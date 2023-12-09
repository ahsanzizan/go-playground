package main

import (
	"html/template"
	"log"
	"net/http"
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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleView(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusNotFound)
		return
	}

	renderTemplate(w, "view", p)
}

func handleEdit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func handleSave(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
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
	const port string = ":8080"
	http.HandleFunc("/view/", handleView)
	http.HandleFunc("/edit/", handleEdit)
	http.HandleFunc("/save/", handleSave)
	log.Fatal(http.ListenAndServe(port, nil))
}
