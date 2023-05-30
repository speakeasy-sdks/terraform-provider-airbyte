// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceZendeskChatAuthorizationMethodAccessTokenCredentials string

const (
	SourceZendeskChatAuthorizationMethodAccessTokenCredentialsAccessToken SourceZendeskChatAuthorizationMethodAccessTokenCredentials = "access_token"
)

func (e SourceZendeskChatAuthorizationMethodAccessTokenCredentials) ToPointer() *SourceZendeskChatAuthorizationMethodAccessTokenCredentials {
	return &e
}

func (e *SourceZendeskChatAuthorizationMethodAccessTokenCredentials) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceZendeskChatAuthorizationMethodAccessTokenCredentials(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskChatAuthorizationMethodAccessTokenCredentials: %v", v)
	}
}

type SourceZendeskChatAuthorizationMethodAccessToken struct {
	// The Access Token to make authenticated requests.
	AccessToken string                                                     `json:"access_token"`
	Credentials SourceZendeskChatAuthorizationMethodAccessTokenCredentials `json:"credentials"`
}

type SourceZendeskChatAuthorizationMethodOAuth20Credentials string

const (
	SourceZendeskChatAuthorizationMethodOAuth20CredentialsOauth20 SourceZendeskChatAuthorizationMethodOAuth20Credentials = "oauth2.0"
)

func (e SourceZendeskChatAuthorizationMethodOAuth20Credentials) ToPointer() *SourceZendeskChatAuthorizationMethodOAuth20Credentials {
	return &e
}

func (e *SourceZendeskChatAuthorizationMethodOAuth20Credentials) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceZendeskChatAuthorizationMethodOAuth20Credentials(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskChatAuthorizationMethodOAuth20Credentials: %v", v)
	}
}

type SourceZendeskChatAuthorizationMethodOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken *string `json:"access_token,omitempty"`
	// The Client ID of your OAuth application
	ClientID *string `json:"client_id,omitempty"`
	// The Client Secret of your OAuth application.
	ClientSecret *string                                                `json:"client_secret,omitempty"`
	Credentials  SourceZendeskChatAuthorizationMethodOAuth20Credentials `json:"credentials"`
	// Refresh Token to obtain new Access Token, when it's expired.
	RefreshToken *string `json:"refresh_token,omitempty"`
}

type SourceZendeskChatAuthorizationMethodType string

const (
	SourceZendeskChatAuthorizationMethodTypeSourceZendeskChatAuthorizationMethodOAuth20     SourceZendeskChatAuthorizationMethodType = "source-zendesk-chat_Authorization Method_OAuth2.0"
	SourceZendeskChatAuthorizationMethodTypeSourceZendeskChatAuthorizationMethodAccessToken SourceZendeskChatAuthorizationMethodType = "source-zendesk-chat_Authorization Method_Access Token"
)

type SourceZendeskChatAuthorizationMethod struct {
	SourceZendeskChatAuthorizationMethodOAuth20     *SourceZendeskChatAuthorizationMethodOAuth20
	SourceZendeskChatAuthorizationMethodAccessToken *SourceZendeskChatAuthorizationMethodAccessToken

	Type SourceZendeskChatAuthorizationMethodType
}

func CreateSourceZendeskChatAuthorizationMethodSourceZendeskChatAuthorizationMethodOAuth20(sourceZendeskChatAuthorizationMethodOAuth20 SourceZendeskChatAuthorizationMethodOAuth20) SourceZendeskChatAuthorizationMethod {
	typ := SourceZendeskChatAuthorizationMethodTypeSourceZendeskChatAuthorizationMethodOAuth20

	return SourceZendeskChatAuthorizationMethod{
		SourceZendeskChatAuthorizationMethodOAuth20: &sourceZendeskChatAuthorizationMethodOAuth20,
		Type: typ,
	}
}

func CreateSourceZendeskChatAuthorizationMethodSourceZendeskChatAuthorizationMethodAccessToken(sourceZendeskChatAuthorizationMethodAccessToken SourceZendeskChatAuthorizationMethodAccessToken) SourceZendeskChatAuthorizationMethod {
	typ := SourceZendeskChatAuthorizationMethodTypeSourceZendeskChatAuthorizationMethodAccessToken

	return SourceZendeskChatAuthorizationMethod{
		SourceZendeskChatAuthorizationMethodAccessToken: &sourceZendeskChatAuthorizationMethodAccessToken,
		Type: typ,
	}
}

func (u *SourceZendeskChatAuthorizationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceZendeskChatAuthorizationMethodOAuth20 := new(SourceZendeskChatAuthorizationMethodOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceZendeskChatAuthorizationMethodOAuth20); err == nil {
		u.SourceZendeskChatAuthorizationMethodOAuth20 = sourceZendeskChatAuthorizationMethodOAuth20
		u.Type = SourceZendeskChatAuthorizationMethodTypeSourceZendeskChatAuthorizationMethodOAuth20
		return nil
	}

	sourceZendeskChatAuthorizationMethodAccessToken := new(SourceZendeskChatAuthorizationMethodAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceZendeskChatAuthorizationMethodAccessToken); err == nil {
		u.SourceZendeskChatAuthorizationMethodAccessToken = sourceZendeskChatAuthorizationMethodAccessToken
		u.Type = SourceZendeskChatAuthorizationMethodTypeSourceZendeskChatAuthorizationMethodAccessToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceZendeskChatAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceZendeskChatAuthorizationMethodOAuth20 != nil {
		return json.Marshal(u.SourceZendeskChatAuthorizationMethodOAuth20)
	}

	if u.SourceZendeskChatAuthorizationMethodAccessToken != nil {
		return json.Marshal(u.SourceZendeskChatAuthorizationMethodAccessToken)
	}

	return nil, nil
}

type SourceZendeskChatZendeskChat string

const (
	SourceZendeskChatZendeskChatZendeskChat SourceZendeskChatZendeskChat = "zendesk-chat"
)

func (e SourceZendeskChatZendeskChat) ToPointer() *SourceZendeskChatZendeskChat {
	return &e
}

func (e *SourceZendeskChatZendeskChat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "zendesk-chat":
		*e = SourceZendeskChatZendeskChat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskChatZendeskChat: %v", v)
	}
}

type SourceZendeskChat struct {
	Credentials *SourceZendeskChatAuthorizationMethod `json:"credentials,omitempty"`
	SourceType  SourceZendeskChatZendeskChat          `json:"sourceType"`
	// The date from which you'd like to replicate data for Zendesk Chat API, in the format YYYY-MM-DDT00:00:00Z.
	StartDate time.Time `json:"start_date"`
	// Required if you access Zendesk Chat from a Zendesk Support subdomain.
	Subdomain *string `json:"subdomain,omitempty"`
}
