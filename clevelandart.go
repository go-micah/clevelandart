// Package clevelandart offers users an easy way to fetch data from the Cleveland Museum of Art Open API
package clevelandart

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/google/go-querystring/query"
)

// endpoint is the current endpoint for Cleveland's Open API
const endpoint = "https://openaccess-api.clevelandart.org/api/"

// ArtworkParams is a struct representing the paramesters available to query artworks
type ArtworkParams struct {
	Indent   int    `url:"indent"`
	Query    string `url:"q"`
	Limit    int    `url:"limit"`
	Skip     int    `url:"skip"`
	HasImage int    `url:"has_image"`
}

// Artwork is a structured data type for an individual artwork record
type Artwork struct {
	Data Data `json:"data"`
}

// Artworks is a structured data type for multiple artwork records
type Artworks struct {
	Info struct {
		Total      int `json:"total"`
		Parameters struct {
			Indent int `json:"indent"`
			Limit  int `json:"limit"`
		} `json:"parameters"`
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
	Creators                 []struct {
		Description string `json:"description"`
		Role        string `json:"role"`
		Biography   string `json:"biography"`
	} `json:"creators"`
	Culture     []string `json:"culture"`
	Description string   `json:"description"`
}

// GetArtworkByID returns a single artwork by its ID
func GetArtworkByID(id int, params ArtworkParams) (*Artwork, error) {

	indent := params.Indent

	resp, err := http.Get(fmt.Sprintf(endpoint+"artworks/%v?indent=%v", id, indent))

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

// SearchArtworks returns a slice of Artwork objects
func SearchArtworks(query string, params ArtworkParams) (*Artworks, error) {

	queryString, err := serializeParams(params)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(fmt.Sprintf(endpoint+"artworks/?%v", queryString))

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

	results := Artworks{}

	err = json.Unmarshal(b, &results)
	if err != nil {
		return nil, errors.New("could not unmarshal data")
	}

	return &results, nil

}

// serializedParams converts ArtworkParams to a query string parameter and returns as a string
func serializeParams(params ArtworkParams) (string, error) {
	v, err := query.Values(params)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

func GetRandomArtwork(hasImage bool) (*Artwork, error) {

	// Call SearchArtworks to get the total number of artworks in the collection

	params := ArtworkParams{
		Indent: 1,
		Limit:  1,
	}

	if hasImage {
		params.HasImage = 1
	}

	art, err := SearchArtworks("", params)
	if err != nil {
		return nil, err
	}

	// Call SearchArtworks using a randomly generated index as the Skip param
	max := art.Info.Total

	params.Skip = rand.Intn(max)
	randomArt, err := SearchArtworks("", params)
	if err != nil {
		return nil, err
	}

	result := Artwork{}
	result.Data = randomArt.Data[0]

	return &result, nil
}
