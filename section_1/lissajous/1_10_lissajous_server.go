package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler prints the lissajous gif at the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var cycleInput int
	if cycles := q.Get("cycles"); cycles != "" {
		cycleInput, err := strconv.Atoi(q.Get("cycles"))
		if err != nil {
			fmt.Printf("Could not convert cycleInput: %v", cycleInput)
			os.Exit(1)
		}
	}

	lissajous(w, &cycleInput)
}
