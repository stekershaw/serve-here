# serve-here

Simple *go* app to serve files from the current directory.

Default port is 8888.

As a simple app, it implements only very simple logging of requests.

## Build/run

* Build: go build
* Run: ./serve-here -port=1234

Cross-compiling can be done as usual (after appropriate setup), e.g.:

    GOOS=linux GOARCH=amd64 go build
