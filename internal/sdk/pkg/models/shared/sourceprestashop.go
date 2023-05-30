// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
	"encoding/json"
	"fmt"
)

type SourcePrestashopPrestashop string

const (
	SourcePrestashopPrestashopPrestashop SourcePrestashopPrestashop = "prestashop"
)

func (e SourcePrestashopPrestashop) ToPointer() *SourcePrestashopPrestashop {
	return &e
}

func (e *SourcePrestashopPrestashop) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "prestashop":
		*e = SourcePrestashopPrestashop(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePrestashopPrestashop: %v", v)
	}
}

type SourcePrestashop struct {
	// Your PrestaShop access key. See <a href="https://devdocs.prestashop.com/1.7/webservice/tutorials/creating-access/#create-an-access-key"> the docs </a> for info on how to obtain this.
	AccessKey  string                     `json:"access_key"`
	SourceType SourcePrestashopPrestashop `json:"sourceType"`
	// The Start date in the format YYYY-MM-DD.
	StartDate types.Date `json:"start_date"`
	// Shop URL without trailing slash.
	URL string `json:"url"`
}
