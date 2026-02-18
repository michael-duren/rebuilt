package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/michael-duren/rebuilt/internal/words"
)

const workers = 50
const requests = 400_000

type requestResult struct {
	statusCode int
	err        error
}

var client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     30 * time.Second,
	},
}

func main() {
	port := flag.String("port", ":80", "the port number of the localhost addr to request")
	flag.Parse()
	url := fmt.Sprintf("http://127.0.0.1%s", *port)

	_, err := client.Get(url)
	if err != nil {
		fmt.Printf("initial request failed with error: %v, stopping program\n", err)
		os.Exit(1)
	}

	fmt.Println("starting worker pool requests")
	statusCodesCh := make(chan requestResult)
	dict := words.GetWords(requests)

	start := time.Now()
	go func() {
		for i := 0; i < requests; i += workers {
			var wg sync.WaitGroup
			for j := range workers {
				wg.Add(1)
				p := fmt.Sprintf("%s/%s", url, dict[i+j])
				go executeRequest(p, statusCodesCh, &wg)
			}
			wg.Wait()
		}
		close(statusCodesCh)
	}()

	counts := make(map[int]int)
	errs := make(map[string]int)

	for rr := range statusCodesCh {
		if rr.err != nil {
			errs[rr.err.Error()]++
			continue
		}
		counts[rr.statusCode]++
	}

	elapsed := time.Since(start)
	rps := float64(requests) / elapsed.Seconds()
	fmt.Printf("completed %d requests in %s (%.0f req/s)\n", requests, elapsed, rps)
	printResults(counts, errs)
}

func printResults(counts map[int]int, errs map[string]int) {
	fmt.Printf("===PRINTING HTTP STATUSCODE results from %d requests, NOTE: errored requests will not print===\n", requests)
	for code, count := range counts {
		fmt.Printf("STATUS CODE %d: %d RESPONSES\n", code, count)
	}

	fmt.Printf("===PRINTING ERROR results from %d requests===\n", requests)
	for err, errCount := range errs {
		fmt.Printf("ERROR: %q: %d RESPONSES\n", err, errCount)
	}
}

func executeRequest(path string, statusCodes chan<- requestResult, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := client.Get(path)
	if err != nil {
		statusCodes <- requestResult{err: err}
		return
	}

	// drains body so go uses same tcp connection
	defer res.Body.Close()
	io.Copy(io.Discard, res.Body)
	statusCodes <- requestResult{statusCode: res.StatusCode}
}
