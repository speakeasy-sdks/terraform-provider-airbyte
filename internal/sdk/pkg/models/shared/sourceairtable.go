// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceAirtableAuthenticationPersonalAccessTokenAuthMethod string

const (
	SourceAirtableAuthenticationPersonalAccessTokenAuthMethodAPIKey SourceAirtableAuthenticationPersonalAccessTokenAuthMethod = "api_key"
)

func (e SourceAirtableAuthenticationPersonalAccessTokenAuthMethod) ToPointer() *SourceAirtableAuthenticationPersonalAccessTokenAuthMethod {
	return &e
}

func (e *SourceAirtableAuthenticationPersonalAccessTokenAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_key":
		*e = SourceAirtableAuthenticationPersonalAccessTokenAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAirtableAuthenticationPersonalAccessTokenAuthMethod: %v", v)
	}
}

type SourceAirtableAuthenticationPersonalAccessToken struct {
	// The Personal Access Token for the Airtable account. See the <a href="https://airtable.com/developers/web/guides/personal-access-tokens">Support Guide</a> for more information on how to obtain this token.
	APIKey     string                                                     `json:"api_key"`
	AuthMethod *SourceAirtableAuthenticationPersonalAccessTokenAuthMethod `json:"auth_method,omitempty"`
}

type SourceAirtableAuthenticationOAuth20AuthMethod string

const (
	SourceAirtableAuthenticationOAuth20AuthMethodOauth20 SourceAirtableAuthenticationOAuth20AuthMethod = "oauth2.0"
)

func (e SourceAirtableAuthenticationOAuth20AuthMethod) ToPointer() *SourceAirtableAuthenticationOAuth20AuthMethod {
	return &e
}

func (e *SourceAirtableAuthenticationOAuth20AuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceAirtableAuthenticationOAuth20AuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAirtableAuthenticationOAuth20AuthMethod: %v", v)
	}
}

type SourceAirtableAuthenticationOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken *string                                        `json:"access_token,omitempty"`
	AuthMethod  *SourceAirtableAuthenticationOAuth20AuthMethod `json:"auth_method,omitempty"`
	// The client ID of the Airtable developer application.
	ClientID string `json:"client_id"`
	// The client secret the Airtable developer application.
	ClientSecret string `json:"client_secret"`
	// The key to refresh the expired access token.
	RefreshToken string `json:"refresh_token"`
	// The date-time when the access token should be refreshed.
	TokenExpiryDate *time.Time `json:"token_expiry_date,omitempty"`
}

type SourceAirtableAuthenticationType string

const (
	SourceAirtableAuthenticationTypeSourceAirtableAuthenticationOAuth20             SourceAirtableAuthenticationType = "source-airtable_Authentication_OAuth2.0"
	SourceAirtableAuthenticationTypeSourceAirtableAuthenticationPersonalAccessToken SourceAirtableAuthenticationType = "source-airtable_Authentication_Personal Access Token"
)

type SourceAirtableAuthentication struct {
	SourceAirtableAuthenticationOAuth20             *SourceAirtableAuthenticationOAuth20
	SourceAirtableAuthenticationPersonalAccessToken *SourceAirtableAuthenticationPersonalAccessToken

	Type SourceAirtableAuthenticationType
}

func CreateSourceAirtableAuthenticationSourceAirtableAuthenticationOAuth20(sourceAirtableAuthenticationOAuth20 SourceAirtableAuthenticationOAuth20) SourceAirtableAuthentication {
	typ := SourceAirtableAuthenticationTypeSourceAirtableAuthenticationOAuth20

	return SourceAirtableAuthentication{
		SourceAirtableAuthenticationOAuth20: &sourceAirtableAuthenticationOAuth20,
		Type:                                typ,
	}
}

func CreateSourceAirtableAuthenticationSourceAirtableAuthenticationPersonalAccessToken(sourceAirtableAuthenticationPersonalAccessToken SourceAirtableAuthenticationPersonalAccessToken) SourceAirtableAuthentication {
	typ := SourceAirtableAuthenticationTypeSourceAirtableAuthenticationPersonalAccessToken

	return SourceAirtableAuthentication{
		SourceAirtableAuthenticationPersonalAccessToken: &sourceAirtableAuthenticationPersonalAccessToken,
		Type: typ,
	}
}

func (u *SourceAirtableAuthentication) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceAirtableAuthenticationOAuth20 := new(SourceAirtableAuthenticationOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceAirtableAuthenticationOAuth20); err == nil {
		u.SourceAirtableAuthenticationOAuth20 = sourceAirtableAuthenticationOAuth20
		u.Type = SourceAirtableAuthenticationTypeSourceAirtableAuthenticationOAuth20
		return nil
	}

	sourceAirtableAuthenticationPersonalAccessToken := new(SourceAirtableAuthenticationPersonalAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceAirtableAuthenticationPersonalAccessToken); err == nil {
		u.SourceAirtableAuthenticationPersonalAccessToken = sourceAirtableAuthenticationPersonalAccessToken
		u.Type = SourceAirtableAuthenticationTypeSourceAirtableAuthenticationPersonalAccessToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceAirtableAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceAirtableAuthenticationOAuth20 != nil {
		return json.Marshal(u.SourceAirtableAuthenticationOAuth20)
	}

	if u.SourceAirtableAuthenticationPersonalAccessToken != nil {
		return json.Marshal(u.SourceAirtableAuthenticationPersonalAccessToken)
	}

	return nil, nil
}

type SourceAirtableAirtable string

const (
	SourceAirtableAirtableAirtable SourceAirtableAirtable = "airtable"
)

func (e SourceAirtableAirtable) ToPointer() *SourceAirtableAirtable {
	return &e
}

func (e *SourceAirtableAirtable) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "airtable":
		*e = SourceAirtableAirtable(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAirtableAirtable: %v", v)
	}
}

type SourceAirtable struct {
	Credentials *SourceAirtableAuthentication `json:"credentials,omitempty"`
	SourceType  SourceAirtableAirtable        `json:"sourceType"`
}
