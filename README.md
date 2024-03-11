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

```
> This rare, dramatic object served as the back support of a litter carried by human porters, a mode of transport reserved for honored members of many societies without draft animals or wheeled vehicles. The simple, bold figures—perhaps a Chimú lord and four officials—all wear wide collars, tunics, and crescent headdresses that are either brightly painted or covered with golden but now-corroded sheet metal. The holes at the bottom probably served as lashing points for a beam that supported the litter's seat.
```