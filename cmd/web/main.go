package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/zees-dev/snippetbox/pkg/config"
	"github.com/zees-dev/snippetbox/pkg/http/rest"
)

func main() {
	cfg := new(config.Server)

	flag.StringVar(&cfg.Port, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.HTMLDir, "html-dir", "./ui/html", "Path to HTML templates")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the serve mux we just created. If ListenAndServe() returns an error we
	// use the log.Fatal() function to record the error message and exit the program.
	mux := rest.CreateApp(cfg)
	log.Printf("Starting server on %s", cfg.Port)
	err := http.ListenAndServe(cfg.Port, mux.Routes())
	log.Fatal(err)
}
