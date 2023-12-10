package clevelandart

import (
	"testing"
)

func TestGetArtworkByID(t *testing.T) {

	params := ArtworkParams{
		Indent: 0,
	}

	artworkID := "151904"
	want := "1927.1984"
	art, err := GetArtworkByID(artworkID, params)

	if want != art.Data.AccessionNumber || err != nil {
		t.Fatalf(`GetArtworkByID("151904") = %q, %v, want match for %#q, nil`, art.Data.AccessionNumber, err, want)
	}
}
