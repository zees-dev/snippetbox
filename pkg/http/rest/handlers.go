package rest

import (
	"fmt"
	"net/http"
	"strconv"
)

// Home - a function which writes a plain-text "Hello from Snippetbox" message as the HTTP response body
func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	app.RenderHTML(w, "home.page.html")
}

// ShowSnippet - TODO
func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet (ID %d)...", id)
}

// NewSnippet - TODO
func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the new snippet form..."))
}
