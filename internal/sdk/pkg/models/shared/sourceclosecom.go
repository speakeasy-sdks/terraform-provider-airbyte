// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SourceCloseComCloseCom string

const (
	SourceCloseComCloseComCloseCom SourceCloseComCloseCom = "close-com"
)

func (e SourceCloseComCloseCom) ToPointer() *SourceCloseComCloseCom {
	return &e
}

func (e *SourceCloseComCloseCom) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "close-com":
		*e = SourceCloseComCloseCom(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceCloseComCloseCom: %v", v)
	}
}

type SourceCloseCom struct {
	// Close.com API key (usually starts with 'api_'; find yours <a href="https://app.close.com/settings/api/">here</a>).
	APIKey     string                 `json:"api_key"`
	SourceType SourceCloseComCloseCom `json:"sourceType"`
	// The start date to sync data. Leave blank for full sync. Format: YYYY-MM-DD.
	StartDate *time.Time `json:"start_date,omitempty"`
}
