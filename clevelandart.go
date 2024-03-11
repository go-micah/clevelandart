// Package clevelandart provides types and functions for working with the Cleveland Museum of Art's Open Access API.
package clevelandart

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// endpoint is the current endpoint for Cleveland's Open API
const endpoint = "https://openaccess-api.clevelandart.org/api/"

// GetArtworkResponse represents the response from an API call to retrieve a single object
type GetArtworkResponse struct {
	Artwork Artwork `json:"data"`
}

// Artwork represents a single object in the Cleveland Museum of Art's Open Access API.
type Artwork struct {
	Id                       int      `json:"id"`
	AccessionNumber          string   `json:"accession_number"`
	ShareLicenseStatus       string   `json:"share_license_status,omitempty"`
	Tombstone                string   `json:"tombstone,omitempty"`
	CurrentLocation          string   `json:"current_location,omitempty"`
	Title                    string   `json:"title,omitempty"`
	TitleInOriginalLanguage  string   `json:"title_in_original_language,omitempty"`
	Series                   string   `json:"series,omitempty"`
	SeriesInOriginalLanguage string   `json:"series_in_original_language,omitempty"`
	CreationDate             string   `json:"creation_date,omitempty"`
	CreationDateEarliest     int      `json:"creation_date_earliest,omitempty"`
	CreationDateLatest       int      `json:"creation_date_latest,omitempty"`
	ArtistsTags              []string `json:"artists_tags,omitempty"`
	Creators                 []struct {
		Description string `json:"description,omitempty"`
		Role        string `json:"role,omitempty"`
		Biography   string `json:"biography,omitempty"`
	} `json:"creators,omitempty"`
	Culture     []string `json:"culture,omitempty"`
	Technique   string   `json:"technique,omitempty"`
	Department  string   `json:"department,omitempty"`
	Description string   `json:"description,omitempty"`
	URL         string   `json:"url,omitempty"`
	JSON        string
}

// GetArtwork retrieves an object from the Cleveland Museum of Art's Open Access API
func GetArtwork(id string) (*Artwork, error) {

	// an id can be either an Id or an Accession Number
	resp, err := http.Get(fmt.Sprintf(endpoint+"artworks/%v", id))
	if err != nil {
		return nil, fmt.Errorf("error communicating with Cleveland, %v", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", "could not fetch data", resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "could not read data", err)
	}

	var a GetArtworkResponse

	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "could not unmarshal data", err)
	}

	// save the raw json reponse to the Artwork struct
	a.Artwork.JSON = string(b)

	return &a.Artwork, nil
}
