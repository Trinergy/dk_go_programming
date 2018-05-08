// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const prefix string = "http://"

func main() {
	for _, url := range os.Args[1:] {
		// 1.8 Add http:// prefix if not included in url
		if !strings.HasPrefix(prefix, url) {
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// 1.7 Use io.Copy instead of ioutil.ReadAll
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		// 1.9 Return status code from request
		fmt.Printf("Bytes Copied: %v STATUS CODE: %v", b, resp.StatusCode)
	}
}
