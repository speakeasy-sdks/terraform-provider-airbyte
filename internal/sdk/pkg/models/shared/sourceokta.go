// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceOktaAuthorizationMethodAPITokenAuthType string

const (
	SourceOktaAuthorizationMethodAPITokenAuthTypeAPIToken SourceOktaAuthorizationMethodAPITokenAuthType = "api_token"
)

func (e SourceOktaAuthorizationMethodAPITokenAuthType) ToPointer() *SourceOktaAuthorizationMethodAPITokenAuthType {
	return &e
}

func (e *SourceOktaAuthorizationMethodAPITokenAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_token":
		*e = SourceOktaAuthorizationMethodAPITokenAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOktaAuthorizationMethodAPITokenAuthType: %v", v)
	}
}

type SourceOktaAuthorizationMethodAPIToken struct {
	// An Okta token. See the <a href="https://docs.airbyte.com/integrations/sources/okta">docs</a> for instructions on how to generate it.
	APIToken string                                        `json:"api_token"`
	AuthType SourceOktaAuthorizationMethodAPITokenAuthType `json:"auth_type"`
}

type SourceOktaAuthorizationMethodOAuth20AuthType string

const (
	SourceOktaAuthorizationMethodOAuth20AuthTypeOauth20 SourceOktaAuthorizationMethodOAuth20AuthType = "oauth2.0"
)

func (e SourceOktaAuthorizationMethodOAuth20AuthType) ToPointer() *SourceOktaAuthorizationMethodOAuth20AuthType {
	return &e
}

func (e *SourceOktaAuthorizationMethodOAuth20AuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceOktaAuthorizationMethodOAuth20AuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOktaAuthorizationMethodOAuth20AuthType: %v", v)
	}
}

type SourceOktaAuthorizationMethodOAuth20 struct {
	AuthType SourceOktaAuthorizationMethodOAuth20AuthType `json:"auth_type"`
	// The Client ID of your OAuth application.
	ClientID string `json:"client_id"`
	// The Client Secret of your OAuth application.
	ClientSecret string `json:"client_secret"`
	// Refresh Token to obtain new Access Token, when it's expired.
	RefreshToken string `json:"refresh_token"`
}

type SourceOktaAuthorizationMethodType string

const (
	SourceOktaAuthorizationMethodTypeSourceOktaAuthorizationMethodOAuth20  SourceOktaAuthorizationMethodType = "source-okta_Authorization Method_OAuth2.0"
	SourceOktaAuthorizationMethodTypeSourceOktaAuthorizationMethodAPIToken SourceOktaAuthorizationMethodType = "source-okta_Authorization Method_API Token"
)

type SourceOktaAuthorizationMethod struct {
	SourceOktaAuthorizationMethodOAuth20  *SourceOktaAuthorizationMethodOAuth20
	SourceOktaAuthorizationMethodAPIToken *SourceOktaAuthorizationMethodAPIToken

	Type SourceOktaAuthorizationMethodType
}

func CreateSourceOktaAuthorizationMethodSourceOktaAuthorizationMethodOAuth20(sourceOktaAuthorizationMethodOAuth20 SourceOktaAuthorizationMethodOAuth20) SourceOktaAuthorizationMethod {
	typ := SourceOktaAuthorizationMethodTypeSourceOktaAuthorizationMethodOAuth20

	return SourceOktaAuthorizationMethod{
		SourceOktaAuthorizationMethodOAuth20: &sourceOktaAuthorizationMethodOAuth20,
		Type:                                 typ,
	}
}

func CreateSourceOktaAuthorizationMethodSourceOktaAuthorizationMethodAPIToken(sourceOktaAuthorizationMethodAPIToken SourceOktaAuthorizationMethodAPIToken) SourceOktaAuthorizationMethod {
	typ := SourceOktaAuthorizationMethodTypeSourceOktaAuthorizationMethodAPIToken

	return SourceOktaAuthorizationMethod{
		SourceOktaAuthorizationMethodAPIToken: &sourceOktaAuthorizationMethodAPIToken,
		Type:                                  typ,
	}
}

func (u *SourceOktaAuthorizationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceOktaAuthorizationMethodOAuth20 := new(SourceOktaAuthorizationMethodOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceOktaAuthorizationMethodOAuth20); err == nil {
		u.SourceOktaAuthorizationMethodOAuth20 = sourceOktaAuthorizationMethodOAuth20
		u.Type = SourceOktaAuthorizationMethodTypeSourceOktaAuthorizationMethodOAuth20
		return nil
	}

	sourceOktaAuthorizationMethodAPIToken := new(SourceOktaAuthorizationMethodAPIToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceOktaAuthorizationMethodAPIToken); err == nil {
		u.SourceOktaAuthorizationMethodAPIToken = sourceOktaAuthorizationMethodAPIToken
		u.Type = SourceOktaAuthorizationMethodTypeSourceOktaAuthorizationMethodAPIToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceOktaAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceOktaAuthorizationMethodOAuth20 != nil {
		return json.Marshal(u.SourceOktaAuthorizationMethodOAuth20)
	}

	if u.SourceOktaAuthorizationMethodAPIToken != nil {
		return json.Marshal(u.SourceOktaAuthorizationMethodAPIToken)
	}

	return nil, nil
}

type SourceOktaOkta string

const (
	SourceOktaOktaOkta SourceOktaOkta = "okta"
)

func (e SourceOktaOkta) ToPointer() *SourceOktaOkta {
	return &e
}

func (e *SourceOktaOkta) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "okta":
		*e = SourceOktaOkta(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOktaOkta: %v", v)
	}
}

type SourceOkta struct {
	Credentials *SourceOktaAuthorizationMethod `json:"credentials,omitempty"`
	// The Okta domain. See the <a href="https://docs.airbyte.com/integrations/sources/okta">docs</a> for instructions on how to find it.
	Domain     *string        `json:"domain,omitempty"`
	SourceType SourceOktaOkta `json:"sourceType"`
	// UTC date and time in the format YYYY-MM-DDTHH:MM:SSZ. Any data before this date will not be replicated.
	StartDate *string `json:"start_date,omitempty"`
}
