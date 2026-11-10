package main

import (
	"log"
	"os"
)

var buffer map[string]bool

func main() {
	if len(os.Args) != 2 {
		if len(os.Args) < 2 {
			log.Fatal("you must give an argnent <https://example.com>")
		}
		log.Fatal("too many arguments")
	}
	crawl(os.Args[1], buffer)
}
