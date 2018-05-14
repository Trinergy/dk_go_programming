// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	// 1.10 Print stats output
	f, err := os.Create("fetch_stats.txt")
	defer f.Close()
	if err != nil {
		ch <- fmt.Sprintf("while creating file: %v", err)
	}

	// 1.10 Print output into a log file
	flog, err := os.Create("fetch_log.txt")
	defer f.Close()
	if err != nil {
		ch <- fmt.Sprintf("while creating file: %v", err)
	}

	for _, url := range os.Args[1:] {
		go fetch(url, ch, flog) // start a goroutine
		go fetch(url, ch, flog) // 1.10 make a double request to see if it caches
	}
	for range os.Args[1:] {
		output := <-ch
		io.WriteString(f, output)
		fmt.Println(output) // receive from channel ch
	}

	// 1.10 Lazy receive double fetch
	for range os.Args[1:] {
		output := <-ch
		io.WriteString(f, output)
		fmt.Println(output) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, output io.Writer) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	if output == nil {
		output = ioutil.Discard
	}
	nbytes, err := io.Copy(output, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s\n", secs, nbytes, url)
}
