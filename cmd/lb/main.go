package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var port string

func main() {
	flag.StringVar(&port, "port", ":80", "the port for the load balancer to run on")
	flag.Parse()

	http.HandleFunc("GET /", requestHandler)

	fmt.Fprintf(os.Stdout, "\nStarting server at %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	w.WriteHeader(http.StatusOK)
}

func logRequest(r *http.Request) {
	fmt.Fprintf(os.Stdout, "\nReceived request from %s\n", r.Host)
	fmt.Fprintf(os.Stdout, "%s / %s\n", r.Method, r.Proto)
	fmt.Fprintf(os.Stdout, "Host: %s\n", r.Host)
	fmt.Fprintf(os.Stdout, "User-Agent: %s\n", r.UserAgent())
	fmt.Fprintf(os.Stdout, "Accept: %s\n", r.Header.Get("Accept"))
}
