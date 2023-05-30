// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceGridlyGridly string

const (
	SourceGridlyGridlyGridly SourceGridlyGridly = "gridly"
)

func (e SourceGridlyGridly) ToPointer() *SourceGridlyGridly {
	return &e
}

func (e *SourceGridlyGridly) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "gridly":
		*e = SourceGridlyGridly(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGridlyGridly: %v", v)
	}
}

type SourceGridly struct {
	APIKey string `json:"api_key"`
	// ID of a grid, or can be ID of a branch
	GridID     string             `json:"grid_id"`
	SourceType SourceGridlyGridly `json:"sourceType"`
}
