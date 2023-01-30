package main

import (
	"fmt"

	"github.com/go-micah/clevelandart"
)

func main() {
	art, _ := clevelandart.GetArtworkByID("162080")
	fmt.Println(*art)
}
