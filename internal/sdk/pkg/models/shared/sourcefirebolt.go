// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceFireboltFirebolt string

const (
	SourceFireboltFireboltFirebolt SourceFireboltFirebolt = "firebolt"
)

func (e SourceFireboltFirebolt) ToPointer() *SourceFireboltFirebolt {
	return &e
}

func (e *SourceFireboltFirebolt) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "firebolt":
		*e = SourceFireboltFirebolt(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFireboltFirebolt: %v", v)
	}
}

type SourceFirebolt struct {
	// Firebolt account to login.
	Account *string `json:"account,omitempty"`
	// The database to connect to.
	Database string `json:"database"`
	// Engine name or url to connect to.
	Engine *string `json:"engine,omitempty"`
	// The host name of your Firebolt database.
	Host *string `json:"host,omitempty"`
	// Firebolt password.
	Password   string                 `json:"password"`
	SourceType SourceFireboltFirebolt `json:"sourceType"`
	// Firebolt email address you use to login.
	Username string `json:"username"`
}
