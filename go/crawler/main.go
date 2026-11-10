package main

import (
	"log"
	"os"
	"sync"
)

type data struct {
	mu     sync.Mutex
	buffer map[string]bool
}

var buff = data{buffer: make(map[string]bool)}

func main() {
	if len(os.Args) != 2 {
		if len(os.Args) < 2 {
			log.Fatal("you must give an argnent <https://example.com>")
		}
		log.Fatal("too many arguments")
	}
	crawl(os.Args[1], &buff)
}
