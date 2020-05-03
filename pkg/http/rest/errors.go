package rest

import (
	"log"
	"net/http"
	"runtime/debug"
)

// The ServerError helper writes an error message and stack trace to the log, then // sends a generic 500 Internal Server Error response to the user.
func (app *App) ServerError(w http.ResponseWriter, err error) {
	log.Printf("%s\n%s", err.Error(), debug.Stack())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

// The ClientError helper sends a specific status code and corresponding description // to the user. We'll use this later in the book to send responses like 400 "Bad
// Request" when there's a problem with the request that the user sent.
func (app *App) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implement a NotFound helper. This is simply a
// convenience wrapper around ClientError which sends a 404 Not Found response to // the user.
func (app *App) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}
