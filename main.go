package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// allow port select from the command-line
	listenPortPtr := flag.Int("port", 8888, "TCP port to listen on")
	flag.Parse()

	listenString := ":" + strconv.Itoa(*listenPortPtr)
	log.Printf("Preparing to listen on %s\n", listenString)

	// create an http.Handler that will log access then serve files from pwd
	loggedFileServer := loggingHandler(http.FileServer(http.Dir(".")))
	log.Println("Logging requests, timestamped, with format:")
	log.Println("clientIP:port method path")

	// serve!
	panic(http.ListenAndServe(listenString, loggedFileServer))
}

// loggingHandler prints an access log line then runs the ServeHTTP func of the
// 'next' http.Handler, given as an argument. This is a simple way to 'chain'
// http.Handlers.
func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
