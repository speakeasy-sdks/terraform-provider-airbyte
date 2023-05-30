// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SourceTypeformTypeform string

const (
	SourceTypeformTypeformTypeform SourceTypeformTypeform = "typeform"
)

func (e SourceTypeformTypeform) ToPointer() *SourceTypeformTypeform {
	return &e
}

func (e *SourceTypeformTypeform) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "typeform":
		*e = SourceTypeformTypeform(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceTypeformTypeform: %v", v)
	}
}

type SourceTypeform struct {
	// When this parameter is set, the connector will replicate data only from the input forms. Otherwise, all forms in your Typeform account will be replicated. You can find form IDs in your form URLs. For example, in the URL "https://mysite.typeform.com/to/u6nXL7" the form_id is u6nXL7. You can find form URLs on Share panel
	FormIds    []string               `json:"form_ids,omitempty"`
	SourceType SourceTypeformTypeform `json:"sourceType"`
	// UTC date and time in the format: YYYY-MM-DDTHH:mm:ss[Z]. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
	// The API Token for a Typeform account.
	Token string `json:"token"`
}
