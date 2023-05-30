// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceBambooHrBambooHr string

const (
	SourceBambooHrBambooHrBambooHr SourceBambooHrBambooHr = "bamboo-hr"
)

func (e SourceBambooHrBambooHr) ToPointer() *SourceBambooHrBambooHr {
	return &e
}

func (e *SourceBambooHrBambooHr) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bamboo-hr":
		*e = SourceBambooHrBambooHr(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceBambooHrBambooHr: %v", v)
	}
}

type SourceBambooHr struct {
	// Api key of bamboo hr
	APIKey string `json:"api_key"`
	// Comma-separated list of fields to include in custom reports.
	CustomReportsFields *string `json:"custom_reports_fields,omitempty"`
	// If true, the custom reports endpoint will include the default fields defined here: https://documentation.bamboohr.com/docs/list-of-field-names.
	CustomReportsIncludeDefaultFields *bool                  `json:"custom_reports_include_default_fields,omitempty"`
	SourceType                        SourceBambooHrBambooHr `json:"sourceType"`
	// Sub Domain of bamboo hr
	Subdomain string `json:"subdomain"`
}
