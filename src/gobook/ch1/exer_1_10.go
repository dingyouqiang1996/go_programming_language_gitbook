// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	_ "os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	url := "https://cloud.google.com/bigquery/public-data"
	go fetch(url, ch) // start a goroutine
	fmt.Println(<-ch) // receive from channel ch
	go fetch(url, ch) // start a goroutine
	fmt.Println(<-ch) // receive from channel ch
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}