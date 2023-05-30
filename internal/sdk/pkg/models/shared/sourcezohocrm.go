// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// SourceZohoCrmDataCenterLocation - Please choose the region of your Data Center location. More info by this <a href="https://www.zoho.com/crm/developer/docs/api/v2/multi-dc.html">Link</a>
type SourceZohoCrmDataCenterLocation string

const (
	SourceZohoCrmDataCenterLocationUs SourceZohoCrmDataCenterLocation = "US"
	SourceZohoCrmDataCenterLocationAu SourceZohoCrmDataCenterLocation = "AU"
	SourceZohoCrmDataCenterLocationEu SourceZohoCrmDataCenterLocation = "EU"
	SourceZohoCrmDataCenterLocationIn SourceZohoCrmDataCenterLocation = "IN"
	SourceZohoCrmDataCenterLocationCn SourceZohoCrmDataCenterLocation = "CN"
	SourceZohoCrmDataCenterLocationJp SourceZohoCrmDataCenterLocation = "JP"
)

func (e SourceZohoCrmDataCenterLocation) ToPointer() *SourceZohoCrmDataCenterLocation {
	return &e
}

func (e *SourceZohoCrmDataCenterLocation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US":
		fallthrough
	case "AU":
		fallthrough
	case "EU":
		fallthrough
	case "IN":
		fallthrough
	case "CN":
		fallthrough
	case "JP":
		*e = SourceZohoCrmDataCenterLocation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZohoCrmDataCenterLocation: %v", v)
	}
}

// SourceZohoCRMZohoCRMEdition - Choose your Edition of Zoho CRM to determine API Concurrency Limits
type SourceZohoCRMZohoCRMEdition string

const (
	SourceZohoCRMZohoCRMEditionFree         SourceZohoCRMZohoCRMEdition = "Free"
	SourceZohoCRMZohoCRMEditionStandard     SourceZohoCRMZohoCRMEdition = "Standard"
	SourceZohoCRMZohoCRMEditionProfessional SourceZohoCRMZohoCRMEdition = "Professional"
	SourceZohoCRMZohoCRMEditionEnterprise   SourceZohoCRMZohoCRMEdition = "Enterprise"
	SourceZohoCRMZohoCRMEditionUltimate     SourceZohoCRMZohoCRMEdition = "Ultimate"
)

func (e SourceZohoCRMZohoCRMEdition) ToPointer() *SourceZohoCRMZohoCRMEdition {
	return &e
}

func (e *SourceZohoCRMZohoCRMEdition) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Free":
		fallthrough
	case "Standard":
		fallthrough
	case "Professional":
		fallthrough
	case "Enterprise":
		fallthrough
	case "Ultimate":
		*e = SourceZohoCRMZohoCRMEdition(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZohoCRMZohoCRMEdition: %v", v)
	}
}

// SourceZohoCrmEnvironment - Please choose the environment
type SourceZohoCrmEnvironment string

const (
	SourceZohoCrmEnvironmentProduction SourceZohoCrmEnvironment = "Production"
	SourceZohoCrmEnvironmentDeveloper  SourceZohoCrmEnvironment = "Developer"
	SourceZohoCrmEnvironmentSandbox    SourceZohoCrmEnvironment = "Sandbox"
)

func (e SourceZohoCrmEnvironment) ToPointer() *SourceZohoCrmEnvironment {
	return &e
}

func (e *SourceZohoCrmEnvironment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Production":
		fallthrough
	case "Developer":
		fallthrough
	case "Sandbox":
		*e = SourceZohoCrmEnvironment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZohoCrmEnvironment: %v", v)
	}
}

type SourceZohoCrmZohoCrm string

const (
	SourceZohoCrmZohoCrmZohoCrm SourceZohoCrmZohoCrm = "zoho-crm"
)

func (e SourceZohoCrmZohoCrm) ToPointer() *SourceZohoCrmZohoCrm {
	return &e
}

func (e *SourceZohoCrmZohoCrm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "zoho-crm":
		*e = SourceZohoCrmZohoCrm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZohoCrmZohoCrm: %v", v)
	}
}

type SourceZohoCrm struct {
	// OAuth2.0 Client ID
	ClientID string `json:"client_id"`
	// OAuth2.0 Client Secret
	ClientSecret string `json:"client_secret"`
	// Please choose the region of your Data Center location. More info by this <a href="https://www.zoho.com/crm/developer/docs/api/v2/multi-dc.html">Link</a>
	DcRegion SourceZohoCrmDataCenterLocation `json:"dc_region"`
	// Choose your Edition of Zoho CRM to determine API Concurrency Limits
	Edition SourceZohoCRMZohoCRMEdition `json:"edition"`
	// Please choose the environment
	Environment SourceZohoCrmEnvironment `json:"environment"`
	// OAuth2.0 Refresh Token
	RefreshToken string               `json:"refresh_token"`
	SourceType   SourceZohoCrmZohoCrm `json:"sourceType"`
	// ISO 8601, for instance: `YYYY-MM-DD`, `YYYY-MM-DD HH:MM:SS+HH:MM`
	StartDatetime *time.Time `json:"start_datetime,omitempty"`
}
