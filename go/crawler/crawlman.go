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

	cmd, err := exec.Command("python3", "crawler.py", url).Output()
	if err != nil {
		log.Fatal(err)
	}

	var urls []string
	if err := json.Unmarshal(cmd, &urls); err != nil {
		log.Fatalf("json invalide : %v", err)
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
