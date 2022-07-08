package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/create/", createHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func (p *Page) save() error {
	filename := p.Title + ".txt"
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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)

	if err != nil {
		fmt.Fprintf(w, "An error occurred trying to get file %s, trace error: %s ", title, err.Error())
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/create/"):]
	page := &Page{Title: title, Body: []byte("this is a file from web app")}
	err := page.save()
	if err != nil {
		fmt.Fprintf(w, "An error occurred, please try again.")
		return
	}
	fmt.Fprintf(w, "An new File was created.")

}
