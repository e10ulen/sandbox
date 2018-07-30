package main

import (
	"fmt"
	"log"

	"github.com/comail/colog"
	"github.com/nakabonne/netsurfer"
)

func main() {
	colog.Register()
	// Obtain the URL of the organic page
	urls, err := netsurfer.OrganicSearch("木村月深", 3)
	if err != nil {
		log.Print("e: ", err)
	}
	fmt.Println("Success!")
	for _, url := range urls {
		// Retrieve the title
		title, err := netsurfer.GetTitle(url.String())
		if err != nil {
			log.Print("p: ", err)
		}
		log.Print("d: ", title)
	}
}
