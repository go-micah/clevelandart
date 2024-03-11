package main

import (
	"fmt"
	"log"

	"github.com/go-micah/clevelandart"
)

func main() {
	artwork, err := clevelandart.GetArtwork("1952.233")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(artwork.Description)
	fmt.Println(artwork.Images.Web.Url)
}
