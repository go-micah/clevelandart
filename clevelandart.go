package clevelandart

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const endpoint = "https://openaccess-api.clevelandart.org/api/"

type Artwork struct {
	Data struct {
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
	} `json:"data"`
}

// GetArtworkByID get a single artwork by its ID
func GetArtworkByID(id string) (*Artwork, error) {

	resp, err := http.Get(fmt.Sprintf(endpoint+"artworks/%s", id))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", "could not fetch data", resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
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
