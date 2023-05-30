// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceLokaliseLokalise string

const (
	SourceLokaliseLokaliseLokalise SourceLokaliseLokalise = "lokalise"
)

func (e SourceLokaliseLokalise) ToPointer() *SourceLokaliseLokalise {
	return &e
}

func (e *SourceLokaliseLokalise) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "lokalise":
		*e = SourceLokaliseLokalise(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLokaliseLokalise: %v", v)
	}
}

type SourceLokalise struct {
	// Lokalise API Key with read-access. Available at Profile settings > API tokens. See <a href="https://docs.lokalise.com/en/articles/1929556-api-tokens">here</a>.
	APIKey string `json:"api_key"`
	// Lokalise project ID. Available at Project Settings > General.
	ProjectID  string                 `json:"project_id"`
	SourceType SourceLokaliseLokalise `json:"sourceType"`
}
