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

func TestSerializeParams(t *testing.T) {

	params := ArtworkInputParams{
		Indent: 0,
		Limit:  1,
	}

	got, err := serializeParams(params)
	want := "has_image=0&indent=0&limit=1&q=&skip=0"

	if got != want || err != nil {
		t.Errorf("got %q, %v want %q, nil", got, err, want)
	}
}
func TestSearchArtworks(t *testing.T) {

	params := ArtworkInputParams{
		Indent: 0,
		Limit:  1,
	}

	art, err := SearchArtworks(params)

	got := art.Info.Parameters.Limit
	want := 1

	if got != want || err != nil {
		t.Errorf("got %q, %v want %q, nil", got, err, want)
	}
}

func TestGetRandomArtwork(t *testing.T) {

	hasImage := true
	art, err := GetRandomArtwork(hasImage)

	got := art.Id

	if got <= 0 || err != nil {
		t.Errorf("got %v, %v want > 0, nil", got, err)
	}
}
