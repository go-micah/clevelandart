package clevelandart

import "testing"

func TestGetArtworkById(t *testing.T) {

	id := "151904"

	art, err := GetArtwork(id)

	got := art.AccessionNumber
	want := "1927.1984"

	if got != want || err != nil {
		t.Errorf("got %q, %v want %q, nil", got, err, want)
	}
}

func TestGetArtworkByAccessionNumber(t *testing.T) {

	accessionNumber := "1927.1984"

	art, err := GetArtwork(accessionNumber)

	got := art.Id
	want := 151904

	if got != want || err != nil {
		t.Errorf("got %q, %v want %q, nil", got, err, want)
	}
}
