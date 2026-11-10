package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"os/exec"
	"sync"
)

func crawl(url string, data *data) {
	data.mu.Lock()
	already := data.buffer[url]
	data.buffer[url] = true
	data.mu.Unlock()

	if already {
		return
	}

	cmd := exec.Command("python3", "crawler.py", url)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("python3 failed for %s: %v\n output:\n%s", url, err, output)
	}

	var urls []string
	if err := json.Unmarshal(output, &urls); err != nil {
		log.Fatalf("json invalide %s: %v\n output:\n%s", url, err, output)
	}

	var wait sync.WaitGroup
	for i, u := range urls {
		wait.Add(1)
		defer wait.Done()
		go crawl(html.UnescapeString(u), data)
		//fetched[html.UnescapeString(u)] = true
		fmt.Println(i, html.UnescapeString(u), "raw: ", u)
	}
	wait.Wait()
}
