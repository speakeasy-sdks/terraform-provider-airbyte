// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceEmailoctopusEmailoctopus string

const (
	SourceEmailoctopusEmailoctopusEmailoctopus SourceEmailoctopusEmailoctopus = "emailoctopus"
)

func (e SourceEmailoctopusEmailoctopus) ToPointer() *SourceEmailoctopusEmailoctopus {
	return &e
}

func (e *SourceEmailoctopusEmailoctopus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "emailoctopus":
		*e = SourceEmailoctopusEmailoctopus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceEmailoctopusEmailoctopus: %v", v)
	}
}

type SourceEmailoctopus struct {
	// EmailOctopus API Key. See the <a href="https://help.emailoctopus.com/article/165-how-to-create-and-delete-api-keys">docs</a> for information on how to generate this key.
	APIKey     string                         `json:"api_key"`
	SourceType SourceEmailoctopusEmailoctopus `json:"sourceType"`
}
