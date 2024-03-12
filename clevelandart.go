// Package clevelandart provides types and functions for working with the Cleveland Museum of Art's Open Access API.
package clevelandart

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/google/go-querystring/query"
)

// Endpoint is the current endpoint for Cleveland's Open API
const endpoint = "https://openaccess-api.clevelandart.org/api/"

// A GetArtworkResponse represents the response from an API call to retrieve a single object
type getArtworkResponse struct {
	Artwork Artwork `json:"data"`
}

// An Artwork represents a single object in the Cleveland Museum of Art's Open Access API.
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
	Culture      []string `json:"culture,omitempty"`
	Technique    string   `json:"technique,omitempty"`
	Department   string   `json:"department,omitempty"`
	Collection   string   `json:"collection,omitempty"`
	Type         string   `json:"type,omitempty"`
	Measurements string   `json:"measurements,omitempty"`
	Description  string   `json:"description,omitempty"`
	Url          string   `json:"url,omitempty"`
	Images       struct {
		Web struct {
			Url      string `json:"url,omitempty"`
			Filename string `json:"filename,omitempty"`
			Filesize string `json:"filesize,omitempty"`
			Width    string `json:"width,omitempty"`
			Height   string `json:"height,omitempty"`
		} `json:"web,omitempty"`
		Print struct {
			Url      string `json:"url,omitempty"`
			Filename string `json:"filename,omitempty"`
			Filesize string `json:"filesize,omitempty"`
			Width    string `json:"width,omitempty"`
			Height   string `json:"height,omitempty"`
		} `json:"print,omitempty"`
		Full struct {
			Url      string `json:"url,omitempty"`
			Filename string `json:"filename,omitempty"`
			Filesize string `json:"filesize,omitempty"`
			Width    string `json:"width,omitempty"`
			Height   string `json:"height,omitempty"`
		} `json:"full,omitempty"`
	} `json:"images,omitempty"`
	JSON string
}

// An ArtworkInputParams is a struct representing the paramesters available to query artworks
type ArtworkInputParams struct {
	Indent   int    `url:"indent"`
	Query    string `url:"q"`
	Limit    int    `url:"limit"`
	Skip     int    `url:"skip"`
	HasImage int    `url:"has_image"`
}

// A SearchResults represents the results of a search with query statistics
type SearchResults struct {
	Info struct {
		Total      int `json:"total"`
		Parameters struct {
			Skip   int    `json:"skip"`
			Limit  int    `json:"limit"`
			Indent int    `json:"indent"`
			Query  string `json:"q"`
			Select string `json:"select"`
			Search string `json:"search"`
		} `json:"parameters"`
	} `json:"info"`
	Artworks []Artwork `json:"data"`
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

	var a getArtworkResponse

	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "could not unmarshal data", err)
	}

	// save the raw json reponse to the Artwork struct
	a.Artwork.JSON = string(b)

	return &a.Artwork, nil
}

// SearchArtworks sends a search query to the API and retrieve an array of Artwork objects along with query statistics
func SearchArtworks(params ArtworkInputParams) (*SearchResults, error) {

	queryString, err := serializeParams(params)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "could not serialize parameters", err)
	}

	resp, err := http.Get(fmt.Sprintf(endpoint+"artworks/?%v", queryString))
	if err != nil {
		return nil, fmt.Errorf("error communicating with Cleveland, %v", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", "could not fetch data", resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "could not read data", resp.Status)
	}

	var results SearchResults

	err = json.Unmarshal(b, &results)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "could not unmarshal data", resp.Status)
	}

	return &results, nil
}

// SerializedParams converts ArtworkParams to a query string parameter and returns as a string
func serializeParams(params ArtworkInputParams) (string, error) {
	v, err := query.Values(params)
	if err != nil {
		return "", err
	}
	return v.Encode(), nil
}

// GetRandomArtwork returns a random artwork from the Cleveland Museum of Art's Open Access API
func GetRandomArtwork(hasImage bool) (*Artwork, error) {

	// Call SearchArtworks to get the total number of artworks in the collection
	params := ArtworkInputParams{
		Indent: 1,
		Limit:  1,
	}

	if hasImage {
		params.HasImage = 1
	}

	art, err := SearchArtworks(params)
	if err != nil {
		return nil, err
	}

	// Call SearchArtworks using a randomly generated index as the Skip param
	max := art.Info.Total

	params.Skip = rand.Intn(max)
	randomArt, err := SearchArtworks(params)
	if err != nil {
		return nil, err
	}

	result := Artwork{}
	result = randomArt.Artworks[0]

	return &result, nil
}
