package clevelandart

import (
	"testing"
)

func TestGetArtworkByID(t *testing.T) {
	artworkID := "151904"
	want := "1927.1984"
	art, err := GetArtworkByID(artworkID)

	if want != art.Data.AccessionNumber || err != nil {
		t.Fatalf(`GetArtworkByID("151904") = %q, %v, want match for %#q, nil`, art.Data.AccessionNumber, err, want)
	}
}
