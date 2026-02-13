package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type serverIndex struct {
	mu    sync.Mutex
	index int
}

func (s *serverIndex) GetIndex(n int) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	curr := s.index
	s.index = (s.index + 1) % n
	return curr
}

var port string
var index serverIndex

const CONFIG_PATH = "./servers.yaml"

func main() {
	flag.StringVar(&port, "port", ":80", "the port for the load balancer to run on")
	flag.Parse()

	servers, err := readYAMLConfig()
	if err != nil {
		fmt.Println("error parsing config: \n", err)
		os.Exit(1)
	}

	handler, err := requestHandlerFactory(servers)
	if err != nil {
		fmt.Println("error creating handler: \n", err)
		os.Exit(1)
	}
	http.HandleFunc("/", handler)

	fmt.Fprintf(os.Stdout, "\nStarting server at %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func readYAMLConfig() ([]string, error) {
	file, err := os.Open(CONFIG_PATH)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error opening file: %s\n", err)
		return nil, err
	}

	b, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error reading file: %s\n", err)
		return nil, err
	}
	var s = struct {
		Servers []string `yaml:"servers"`
	}{}

	err = yaml.Unmarshal(b, &s)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error unmarshalling config: %s\n", err)
		return nil, err
	}

	return s.Servers, nil
}

type RequestHandler func(http.ResponseWriter, *http.Request)

func requestHandlerFactory(servers []string) (RequestHandler, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		res, err := forwardTraffic(servers[index.GetIndex(len(servers))], r)
		if err != nil {
			fmt.Fprintf(os.Stdout, "error while forwarding traffic: %v\n", err)
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		defer res.Body.Close()

		w.WriteHeader(res.StatusCode)
		io.Copy(w, res.Body)
	}, nil
}

func forwardTraffic(path string, r *http.Request) (*http.Response, error) {
	fullPath := fmt.Sprintf("%s%s", path, r.URL.Path)
	fmt.Fprintf(os.Stdout, "path: %s, fullPath: %s, r.URL.Path: %s\n", path, fullPath, r.URL.Path)
	res, err := http.Get(fullPath)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error executing forward to: %s: %v\n", fullPath, err)
		return nil, err
	}

	return res, nil
}

func logRequest(r *http.Request) {
	fmt.Fprintf(os.Stdout, "\nReceived request from %s\n", r.Host)
	fmt.Fprintf(os.Stdout, "%s / %s\n", r.Method, r.Proto)
	fmt.Fprintf(os.Stdout, "Host: %s\n", r.Host)
	fmt.Fprintf(os.Stdout, "User-Agent: %s\n", r.UserAgent())
	fmt.Fprintf(os.Stdout, "Accept: %s\n", r.Header.Get("Accept"))
}
