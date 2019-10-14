package main

import (
	"fmt"
	"sort"
	"strings"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	keys := make([]string, 0, len(r.Header))
	for name, _ := range r.Header {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	for _, name := range keys {
		values := r.Header[name]
		// Loop over all values for the name.
		for _, value := range values {
			if strings.HasPrefix(name, "X-") {
				fmt.Fprintf(w, "Name: %s, Value: %s\n", name, value)
			}
		}
	}
}

// Sample  "echo" service displaying UserID and UserRole HTTP request headers
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
