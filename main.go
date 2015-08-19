package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"strconv"
)

func main() {
	// allow port select from the command-line (0 default = OS-assigned "random" port)
	listenPortPtr := flag.Int("port", 0, "TCP port to listen on")
	flag.Parse()

	listenString := ":" + strconv.Itoa(*listenPortPtr)
	listener, err := net.Listen("tcp", listenString)
	if err != nil {
		log.Fatal("Listen:", err)
	}
	log.Printf("Listening on %s\n", listener.Addr())

	// create an http.Handler that will log access then serve files from pwd
	fileServerHandler := http.FileServer(http.Dir("."))
	loggedFileServer := loggingHandler(fileServerHandler)

	// register HTTP handlers (in the DefaultServeMux)
	http.Handle("/", loggedFileServer)
	log.Println("Logging requests, timestamped, with format:")
	log.Println("clientIP:port method path")

	// serve HTTP: using our listener (and DefaultServeMux)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Serve:", err)
	}
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
