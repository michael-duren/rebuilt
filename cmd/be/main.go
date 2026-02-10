package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	port     string
	serverID string
)

var colors = []string{
	"#e74c3c", "#3498db", "#2ecc71", "#f39c12",
	"#9b59b6", "#1abc9c", "#e67e22", "#2980b9",
}

func main() {
	flag.StringVar(&port, "port", ":80", "the port for the load balancer to run on")
	flag.Parse()

	serverID = fmt.Sprintf("%04x", rand.Intn(0xFFFF))

	http.HandleFunc("GET /", requestHandler)

	fmt.Fprintf(os.Stdout, "\nStarting server %s at %s\n", serverID, port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(os.Stdout, "\nReceived request from %s\n", r.Host)
	fmt.Fprintf(os.Stdout, "%s / %s\n", r.Method, r.Proto)
	fmt.Fprintf(os.Stdout, "Host: %s\n", r.Host)
	fmt.Fprintf(os.Stdout, "User-Agent: %s\n", r.UserAgent())
	fmt.Fprintf(os.Stdout, "Accept: %s\n", r.Header.Get("Accept"))

	hostname, _ := os.Hostname()
	color := colors[rand.Intn(len(colors))]
	requestID := fmt.Sprintf("%08x", rand.Intn(0xFFFFFFFF))
	timestamp := time.Now().Format(time.RFC3339)

	html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>Backend Server %s</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: monospace;
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
            background: #1a1a2e;
            color: #eee;
        }
        .card {
            border: 2px solid %s;
            border-radius: 12px;
            padding: 2rem 3rem;
            text-align: center;
            box-shadow: 0 0 30px %s44;
        }
        .server-id {
            font-size: 3rem;
            font-weight: bold;
            color: %s;
        }
        .label { color: #888; margin-top: 1rem; font-size: 0.85rem; }
        .value { font-size: 1.1rem; margin-bottom: 0.5rem; }
        .divider { border-top: 1px solid #333; margin: 1rem 0; }
    </style>
</head>
<body>
    <div class="card">
        <div class="label">SERVER ID</div>
        <div class="server-id">%s</div>
        <div class="divider"></div>
        <div class="label">HOSTNAME</div>
        <div class="value">%s</div>
        <div class="label">REQUEST ID</div>
        <div class="value">%s</div>
        <div class="label">TIMESTAMP</div>
        <div class="value">%s</div>
        <div class="label">LISTENING ON</div>
        <div class="value">%s</div>
    </div>
</body>
</html>`, serverID, color, color, color, serverID, hostname, requestID, timestamp, port)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}
