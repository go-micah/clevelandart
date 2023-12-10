// Package clevelandart offers users an easy way to fetch data from the Cleveland Museum of Art Open API
package clevelandart

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// endpoint is the current endpoint for Cleveland's Open API
const endpoint = "https://openaccess-api.clevelandart.org/api/"

// Artwork is a structured data type for an individual artwork record
type Artwork struct {
	Data Data `json:"data"`
}

// Artworks is a structured data type for multiple artwork records
type Artworks struct {
	Info struct {
		Total int `json:"total"`
	} `json:"info"`
	Data []Data `json:"data"`
}

// Data is a structured data type representing an artwork
type Data struct {
	ID                       int    `json:"id"`
	AccessionNumber          string `json:"accession_number,omitempty"`
	ShareLicenseStatus       string `json:"share_license_status,omitempty"`
	Tombstone                string `json:"tombstone,omitempty"`
	CurrentLocation          string `json:"current_location,omitempty"`
	Title                    string `json:"title,omitempty"`
	TitleInOriginalLanguage  string `json:"title_in_original_language,omitempty"`
	Series                   string `json:"series,omitempty"`
	SeriesInOriginalLanguage string `json:"series_in_original_language,omitempty"`
	CreationDate             string `json:"creation_date,omitempty"`
	CreationDateEarliest     int    `json:"creation_date_earliest,omitempty"`
	CreationDateLatest       int    `json:"creation_date_latest,omitempty"`
}

// GetArtworkByID returns a single artwork by its ID
func GetArtworkByID(id string) (*Artwork, error) {

	resp, err := http.Get(fmt.Sprintf(endpoint+"artworks/%s", id))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", "could not fetch data", resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("could not read data")
	}

	result := Artwork{}

	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}

	return &result, nil

}
