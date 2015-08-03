package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {

	listenPortPtr := flag.Int("port", 8888, "TCP port to listen on")
	flag.Parse()

	listenString := ":" + strconv.Itoa(*listenPortPtr)
	log.Printf("Preparing to listen on %s\n", listenString)
	panic(http.ListenAndServe(listenString, http.FileServer(http.Dir("."))))
}
