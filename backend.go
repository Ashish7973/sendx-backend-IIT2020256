package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

var crawledPages = make(map[string]CrawledPage)
var mu sync.Mutex

type CrawledPage struct {
	URL     string
	Content string
	Time    time.Time
}

var payingWorkers = make(chan struct{}, 5)
var nonPayingWorkers = make(chan struct{}, 2)

func crawlPage(url string, isPaying bool) ([]byte, error) {
    fmt.Println("Initiating crawling for URL:", url)

    var workers chan struct{}
    if isPaying {
        workers = payingWorkers
    } else {
        workers = nonPayingWorkers
    }

    workers <- struct{}{} // Acquire token
    defer func() { <-workers }()

    maxRetries := 3
    retries := 0
    var body []byte
    var err error

    for retries < maxRetries {
        fmt.Printf("Attempt %d for URL: %s\n", retries+1, url)
        resp, err := http.Get(url)
        if err == nil {
            defer resp.Body.Close()
            body, err = ioutil.ReadAll(resp.Body)
            if err == nil {
                break
            }
        }
        retries++
        time.Sleep(time.Second) // Adding a delay before retrying
    }

    if err != nil {
        return nil, err
    }

    mu.Lock()
    defer mu.Unlock()

    if cached, found := crawledPages[url]; found && time.Since(cached.Time).Minutes() <= 60 && (!isPaying || isPaying) {
        fmt.Printf("Retrieving content from cache for URL: %s\n", url)
        return []byte(cached.Content), nil
    }

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    crawledPages[url] = CrawledPage{
        URL:     url,
        Content: string(body),
        Time:    time.Now(),
    }

    fmt.Printf("Crawling completed for URL: %s\n", url)
    return body, nil
}


func crawlHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	url := queryParams.Get("url")
	isPaying := queryParams.Get("isPaying") == "true"

	if url == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received request to crawl URL: %s\n", url)
	crawledContent, err := crawlPage(url, isPaying)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to crawl the page: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", crawledContent)
}

func main() {
	http.HandleFunc("/crawl", crawlHandler)
	http.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Println("Server is listening on port 5500")
	err := http.ListenAndServe(":5500", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
