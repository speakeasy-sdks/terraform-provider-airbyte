// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceWikipediaPageviewsWikipediaPageviews string

const (
	SourceWikipediaPageviewsWikipediaPageviewsWikipediaPageviews SourceWikipediaPageviewsWikipediaPageviews = "wikipedia-pageviews"
)

func (e SourceWikipediaPageviewsWikipediaPageviews) ToPointer() *SourceWikipediaPageviewsWikipediaPageviews {
	return &e
}

func (e *SourceWikipediaPageviewsWikipediaPageviews) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "wikipedia-pageviews":
		*e = SourceWikipediaPageviewsWikipediaPageviews(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceWikipediaPageviewsWikipediaPageviews: %v", v)
	}
}

type SourceWikipediaPageviews struct {
	// If you want to filter by access method, use one of desktop, mobile-app or mobile-web. If you are interested in pageviews regardless of access method, use all-access.
	Access string `json:"access"`
	// If you want to filter by agent type, use one of user, automated or spider. If you are interested in pageviews regardless of agent type, use all-agents.
	Agent string `json:"agent"`
	// The title of any article in the specified project. Any spaces should be replaced with underscores. It also should be URI-encoded, so that non-URI-safe characters like %, / or ? are accepted.
	Article string `json:"article"`
	// The ISO 3166-1 alpha-2 code of a country for which to retrieve top articles.
	Country string `json:"country"`
	// The date of the last day to include, in YYYYMMDD or YYYYMMDDHH format.
	End string `json:"end"`
	// If you want to filter by project, use the domain of any Wikimedia project.
	Project    string                                     `json:"project"`
	SourceType SourceWikipediaPageviewsWikipediaPageviews `json:"sourceType"`
	// The date of the first day to include, in YYYYMMDD or YYYYMMDDHH format.
	Start string `json:"start"`
}
