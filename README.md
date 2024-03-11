# clevelandart

This package contains types and functions to make it easy to work with the [Open Access API](https://openaccess-api.clevelandart.org/) at the Cleveland Museum of Art.

## Example

There is an example program in `cmd/main.go` which you can use to get started like this:

```go
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
}
```