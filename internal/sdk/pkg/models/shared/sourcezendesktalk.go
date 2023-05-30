// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceZendeskTalkAuthenticationOAuth20AuthType string

const (
	SourceZendeskTalkAuthenticationOAuth20AuthTypeOauth20 SourceZendeskTalkAuthenticationOAuth20AuthType = "oauth2.0"
)

func (e SourceZendeskTalkAuthenticationOAuth20AuthType) ToPointer() *SourceZendeskTalkAuthenticationOAuth20AuthType {
	return &e
}

func (e *SourceZendeskTalkAuthenticationOAuth20AuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceZendeskTalkAuthenticationOAuth20AuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskTalkAuthenticationOAuth20AuthType: %v", v)
	}
}

// SourceZendeskTalkAuthenticationOAuth20 - Zendesk service provides two authentication methods. Choose between: `OAuth2.0` or `API token`.
type SourceZendeskTalkAuthenticationOAuth20 struct {
	// The value of the API token generated. See the <a href="https://docs.airbyte.com/integrations/sources/zendesk-talk">docs</a> for more information.
	AccessToken string                                          `json:"access_token"`
	AuthType    *SourceZendeskTalkAuthenticationOAuth20AuthType `json:"auth_type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _SourceZendeskTalkAuthenticationOAuth20 SourceZendeskTalkAuthenticationOAuth20

func (c *SourceZendeskTalkAuthenticationOAuth20) UnmarshalJSON(bs []byte) error {
	data := _SourceZendeskTalkAuthenticationOAuth20{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = SourceZendeskTalkAuthenticationOAuth20(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "access_token")
	delete(additionalFields, "auth_type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c SourceZendeskTalkAuthenticationOAuth20) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_SourceZendeskTalkAuthenticationOAuth20(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}

type SourceZendeskTalkAuthenticationAPITokenAuthType string

const (
	SourceZendeskTalkAuthenticationAPITokenAuthTypeAPIToken SourceZendeskTalkAuthenticationAPITokenAuthType = "api_token"
)

func (e SourceZendeskTalkAuthenticationAPITokenAuthType) ToPointer() *SourceZendeskTalkAuthenticationAPITokenAuthType {
	return &e
}

func (e *SourceZendeskTalkAuthenticationAPITokenAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_token":
		*e = SourceZendeskTalkAuthenticationAPITokenAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskTalkAuthenticationAPITokenAuthType: %v", v)
	}
}

// SourceZendeskTalkAuthenticationAPIToken - Zendesk service provides two authentication methods. Choose between: `OAuth2.0` or `API token`.
type SourceZendeskTalkAuthenticationAPIToken struct {
	// The value of the API token generated. See the <a href="https://docs.airbyte.com/integrations/sources/zendesk-talk">docs</a> for more information.
	APIToken string                                           `json:"api_token"`
	AuthType *SourceZendeskTalkAuthenticationAPITokenAuthType `json:"auth_type,omitempty"`
	// The user email for your Zendesk account.
	Email string `json:"email"`

	AdditionalProperties interface{} `json:"-"`
}
type _SourceZendeskTalkAuthenticationAPIToken SourceZendeskTalkAuthenticationAPIToken

func (c *SourceZendeskTalkAuthenticationAPIToken) UnmarshalJSON(bs []byte) error {
	data := _SourceZendeskTalkAuthenticationAPIToken{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = SourceZendeskTalkAuthenticationAPIToken(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "api_token")
	delete(additionalFields, "auth_type")
	delete(additionalFields, "email")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c SourceZendeskTalkAuthenticationAPIToken) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_SourceZendeskTalkAuthenticationAPIToken(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}

type SourceZendeskTalkAuthenticationType string

const (
	SourceZendeskTalkAuthenticationTypeSourceZendeskTalkAuthenticationAPIToken SourceZendeskTalkAuthenticationType = "source-zendesk-talk_Authentication_API Token"
	SourceZendeskTalkAuthenticationTypeSourceZendeskTalkAuthenticationOAuth20  SourceZendeskTalkAuthenticationType = "source-zendesk-talk_Authentication_OAuth2.0"
)

type SourceZendeskTalkAuthentication struct {
	SourceZendeskTalkAuthenticationAPIToken *SourceZendeskTalkAuthenticationAPIToken
	SourceZendeskTalkAuthenticationOAuth20  *SourceZendeskTalkAuthenticationOAuth20

	Type SourceZendeskTalkAuthenticationType
}

func CreateSourceZendeskTalkAuthenticationSourceZendeskTalkAuthenticationAPIToken(sourceZendeskTalkAuthenticationAPIToken SourceZendeskTalkAuthenticationAPIToken) SourceZendeskTalkAuthentication {
	typ := SourceZendeskTalkAuthenticationTypeSourceZendeskTalkAuthenticationAPIToken

	return SourceZendeskTalkAuthentication{
		SourceZendeskTalkAuthenticationAPIToken: &sourceZendeskTalkAuthenticationAPIToken,
		Type:                                    typ,
	}
}

func CreateSourceZendeskTalkAuthenticationSourceZendeskTalkAuthenticationOAuth20(sourceZendeskTalkAuthenticationOAuth20 SourceZendeskTalkAuthenticationOAuth20) SourceZendeskTalkAuthentication {
	typ := SourceZendeskTalkAuthenticationTypeSourceZendeskTalkAuthenticationOAuth20

	return SourceZendeskTalkAuthentication{
		SourceZendeskTalkAuthenticationOAuth20: &sourceZendeskTalkAuthenticationOAuth20,
		Type:                                   typ,
	}
}

func (u *SourceZendeskTalkAuthentication) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceZendeskTalkAuthenticationAPIToken := new(SourceZendeskTalkAuthenticationAPIToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceZendeskTalkAuthenticationAPIToken); err == nil {
		u.SourceZendeskTalkAuthenticationAPIToken = sourceZendeskTalkAuthenticationAPIToken
		u.Type = SourceZendeskTalkAuthenticationTypeSourceZendeskTalkAuthenticationAPIToken
		return nil
	}

	sourceZendeskTalkAuthenticationOAuth20 := new(SourceZendeskTalkAuthenticationOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceZendeskTalkAuthenticationOAuth20); err == nil {
		u.SourceZendeskTalkAuthenticationOAuth20 = sourceZendeskTalkAuthenticationOAuth20
		u.Type = SourceZendeskTalkAuthenticationTypeSourceZendeskTalkAuthenticationOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceZendeskTalkAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceZendeskTalkAuthenticationAPIToken != nil {
		return json.Marshal(u.SourceZendeskTalkAuthenticationAPIToken)
	}

	if u.SourceZendeskTalkAuthenticationOAuth20 != nil {
		return json.Marshal(u.SourceZendeskTalkAuthenticationOAuth20)
	}

	return nil, nil
}

type SourceZendeskTalkZendeskTalk string

const (
	SourceZendeskTalkZendeskTalkZendeskTalk SourceZendeskTalkZendeskTalk = "zendesk-talk"
)

func (e SourceZendeskTalkZendeskTalk) ToPointer() *SourceZendeskTalkZendeskTalk {
	return &e
}

func (e *SourceZendeskTalkZendeskTalk) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "zendesk-talk":
		*e = SourceZendeskTalkZendeskTalk(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskTalkZendeskTalk: %v", v)
	}
}

type SourceZendeskTalk struct {
	// Zendesk service provides two authentication methods. Choose between: `OAuth2.0` or `API token`.
	Credentials *SourceZendeskTalkAuthentication `json:"credentials,omitempty"`
	SourceType  SourceZendeskTalkZendeskTalk     `json:"sourceType"`
	// The date from which you'd like to replicate data for Zendesk Talk API, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
	StartDate time.Time `json:"start_date"`
	// This is your Zendesk subdomain that can be found in your account URL. For example, in https://{MY_SUBDOMAIN}.zendesk.com/, where MY_SUBDOMAIN is the value of your subdomain.
	Subdomain string `json:"subdomain"`
}
