package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"os/exec"
)

func crawl(url string, fetched map[string]bool) {
	if fetched[url] {
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

	for i, u := range urls {
		fetched[html.UnescapeString(u)] = true
		fmt.Println(i, html.UnescapeString(u), "raw: ", u)
	}
}
