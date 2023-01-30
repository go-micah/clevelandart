package main

import (
	"fmt"
	"os"

	"github.com/go-micah/clevelandart"
)

func main() {

	args := os.Args[1:]
	artworkID := args[0]

	art, err := clevelandart.GetArtworkByID(artworkID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*art)
	}
}
