package main

import (
	"fmt"
	"log"
	"github.com/nakabonne/netsurfer"
)

func main() {
	// Obtain the URL of the organic page
	urls, err := netsurfer.OrganicSearch("e10ulen", 3)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Success!")
	for _, url := range urls {
		// Retrieve the title
		title, err := netsurfer.GetTitle(url.String())
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(title)
}
}
