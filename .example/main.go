package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	MEMCACHE_ENDPOINT = "http://127.0.0.1:8080"
	HOSTNAME          = "127.0.0.1"
	PORT              = 3000

	// Fibonacci sequence length
	N = 46
)

var MEMCACHE_CACHE_ID string = ""

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf("%s:%d", HOSTNAME, PORT), nil)
}

// Fibonacci sequence generator
// Long running process
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Generate Fibonacci sequence
func generateFibonacciSequence(n int) string {
	var sequence []int
	for i := 1; i < n; i++ {
		sequence = append(sequence, fibonacci(i))
	}

	return string(fmt.Sprintf("%v", sequence))
}

// Handler
func handler(w http.ResponseWriter, r *http.Request) {
	if MEMCACHE_CACHE_ID == "" {
		data := map[string]string{"value": generateFibonacciSequence(N)}
		json_data, err := json.Marshal(data)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		res, err := http.Post(MEMCACHE_ENDPOINT, "application/json", bytes.NewBuffer(json_data))
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		defer res.Body.Close()

		// convert response body to string
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		bodyString := buf.String()

		MEMCACHE_CACHE_ID = bodyString
	}

	res, err := http.Get(fmt.Sprintf("%s/%s", MEMCACHE_ENDPOINT, MEMCACHE_CACHE_ID))
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	defer res.Body.Close()

	// convert response body to string
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	bodyString := buf.String()

	fmt.Fprintf(w, "%s", bodyString)
}
