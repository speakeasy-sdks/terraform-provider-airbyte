// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceRecreationRecreation string

const (
	SourceRecreationRecreationRecreation SourceRecreationRecreation = "recreation"
)

func (e SourceRecreationRecreation) ToPointer() *SourceRecreationRecreation {
	return &e
}

func (e *SourceRecreationRecreation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "recreation":
		*e = SourceRecreationRecreation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceRecreationRecreation: %v", v)
	}
}

type SourceRecreation struct {
	// API Key
	Apikey         string                     `json:"apikey"`
	QueryCampsites *string                    `json:"query_campsites,omitempty"`
	SourceType     SourceRecreationRecreation `json:"sourceType"`
}
