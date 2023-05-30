// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DestinationMeilisearchMeilisearch string

const (
	DestinationMeilisearchMeilisearchMeilisearch DestinationMeilisearchMeilisearch = "meilisearch"
)

func (e DestinationMeilisearchMeilisearch) ToPointer() *DestinationMeilisearchMeilisearch {
	return &e
}

func (e *DestinationMeilisearchMeilisearch) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "meilisearch":
		*e = DestinationMeilisearchMeilisearch(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMeilisearchMeilisearch: %v", v)
	}
}

type DestinationMeilisearch struct {
	// MeiliSearch API Key. See the <a href="https://docs.airbyte.com/integrations/destinations/meilisearch">docs</a> for more information on how to obtain this key.
	APIKey          *string                           `json:"api_key,omitempty"`
	DestinationType DestinationMeilisearchMeilisearch `json:"destinationType"`
	// Hostname of the MeiliSearch instance.
	Host string `json:"host"`
}
