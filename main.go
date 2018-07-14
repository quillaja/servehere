// Package main implements a simple http file server mainly for use in
// serving files locally when testing web apps.
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	bindAddress := flag.String("a", "localhost", "The bind address. Use '0.0.0.0' for all interfaces.")
	port := flag.String("p", "8000", "The port to serve on.")
	verbose := flag.Bool("v", true, "Output messages to stdout.")
	flag.Parse()

	log := log.New(os.Stdout, "", log.LstdFlags)

	handler := http.FileServer(http.Dir("."))
	if *verbose {
		handler = loggingHandler(handler, log)
		log.Printf("Starting HTTP fileserver on %s:%s.\n", *bindAddress, *port)
		log.Printf("Press CTRL-C to stop serving.\n\n")
	}

	log.Fatalln(http.ListenAndServe(*bindAddress+":"+*port, handler))
}

// loggingHander wraps an http.Handler and prints request/response info.
func loggingHandler(h http.Handler, log *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logw := newResponseLogger(w)
		h.ServeHTTP(logw, r)

		log.Printf("%s %s %d %s", r.Method, r.RequestURI, logw.statusCode, http.StatusText(logw.statusCode))
	})
}

// responseLogger wrapps a http.ResponseWriter and allows the user to get
// the http status code from the ResponseWriter.
type responseLogger struct {
	http.ResponseWriter
	statusCode int
}

// newResponseLogger creates a responseLogger.
func newResponseLogger(w http.ResponseWriter) *responseLogger {
	return &responseLogger{w, http.StatusOK}
}

// WriteHeader "overrides" ResponseWriter.WriteHeader() method and stores
// the status code.
func (rl *responseLogger) WriteHeader(code int) {
	rl.statusCode = code
	rl.ResponseWriter.WriteHeader(code)
}
