package main 

import "fmt"
import "os"
import "os/exec"
import "log"
import "html"
import "encoding/json"

func main()  {
	if len(os.Args) != 2{
		if len(os.Args) < 2{log.Fatal("you must give an argnent <www.example.com>")}
		log.Fatal("too many arguments")
	}
	cmd , err := exec.Command("python3", "crawler.py", os.Args[1]).Output()
	if err != nil{
		log.Fatal(err)
	}

	var urls []string
	if err := json.Unmarshal(cmd, &urls); err != nil {
		log.Fatalf("json invalide : %v", err)
	}
	
	for i, u := range urls {
		fmt.Println(i, html.UnescapeString(u))
	}
}
