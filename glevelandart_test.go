package clevelandart

import (
	"fmt"
	"testing"
)

func TestGetArtworkByID(t *testing.T) {

	params := ArtworkParams{
		Indent: 0,
	}

	artworkID := 151904

	art, err := GetArtworkByID(artworkID, params)

	got := art.Data.AccessionNumber
	want := "1927.1984"

	if got != want || err != nil {
		t.Errorf("got %q, %v want %q, nil", got, err, want)
	}
}

func TestSearchArtworks(t *testing.T) {

	params := ArtworkParams{
		Indent: 0,
		Limit:  1,
	}

	query := ""
	art, err := SearchArtworks(query, params)

	got := art.Info.Parameters.Limit
	want := 1

	if got != want || err != nil {
		t.Errorf("got %q, %v want %q, nil", got, err, want)
	}

}

func TestSerializeParams(t *testing.T) {

	params := ArtworkParams{
		Indent: 0,
		Limit:  1,
	}

	got, err := serializeParams(params)
	want := "indent=0&limit=1&q="

	if got != want || err != nil {
		t.Errorf("got %q, %v want %q, nil", got, err, want)
	}
}

func TestGetRandomArtwork(t *testing.T) {

	hasImage := true
	art, err := GetRandomArtwork(hasImage)

	got := art.Data.ID

	fmt.Print(art.Data)
	
	if got <= 0 || err != nil {
		t.Errorf("got %v, %v want > 0, nil", got, err)
	}
}
