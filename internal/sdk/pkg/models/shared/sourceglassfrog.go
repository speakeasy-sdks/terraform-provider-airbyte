// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceGlassfrogGlassfrog string

const (
	SourceGlassfrogGlassfrogGlassfrog SourceGlassfrogGlassfrog = "glassfrog"
)

func (e SourceGlassfrogGlassfrog) ToPointer() *SourceGlassfrogGlassfrog {
	return &e
}

func (e *SourceGlassfrogGlassfrog) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "glassfrog":
		*e = SourceGlassfrogGlassfrog(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGlassfrogGlassfrog: %v", v)
	}
}

type SourceGlassfrog struct {
	// API key provided by Glassfrog
	APIKey     string                   `json:"api_key"`
	SourceType SourceGlassfrogGlassfrog `json:"sourceType"`
}
