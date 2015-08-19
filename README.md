# serve-here

Simple *go* app to serve files from the current directory.

Default port is chosen by the OS. Run the app with e.g. *-port=1234* to choose a specific port.

As a simple app, it implements only very simple logging of requests.

## Build/run

* Build: go build
* Run: ./serve-here -port=1234

Cross-compiling can be done as usual (after appropriate setup), e.g.:

    GOOS=linux GOARCH=amd64 go build
