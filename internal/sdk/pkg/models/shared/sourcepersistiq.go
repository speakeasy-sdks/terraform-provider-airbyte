// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourcePersistiqPersistiq string

const (
	SourcePersistiqPersistiqPersistiq SourcePersistiqPersistiq = "persistiq"
)

func (e SourcePersistiqPersistiq) ToPointer() *SourcePersistiqPersistiq {
	return &e
}

func (e *SourcePersistiqPersistiq) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "persistiq":
		*e = SourcePersistiqPersistiq(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePersistiqPersistiq: %v", v)
	}
}

type SourcePersistiq struct {
	// PersistIq API Key. See the <a href="https://apidocs.persistiq.com/#authentication">docs</a> for more information on where to find that key.
	APIKey     string                   `json:"api_key"`
	SourceType SourcePersistiqPersistiq `json:"sourceType"`
}
