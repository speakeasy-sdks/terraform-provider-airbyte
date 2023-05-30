// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationElasticsearchAuthenticationMethodUsernamePasswordMethod string

const (
	DestinationElasticsearchAuthenticationMethodUsernamePasswordMethodBasic DestinationElasticsearchAuthenticationMethodUsernamePasswordMethod = "basic"
)

func (e DestinationElasticsearchAuthenticationMethodUsernamePasswordMethod) ToPointer() *DestinationElasticsearchAuthenticationMethodUsernamePasswordMethod {
	return &e
}

func (e *DestinationElasticsearchAuthenticationMethodUsernamePasswordMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		*e = DestinationElasticsearchAuthenticationMethodUsernamePasswordMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationElasticsearchAuthenticationMethodUsernamePasswordMethod: %v", v)
	}
}

// DestinationElasticsearchAuthenticationMethodUsernamePassword - Basic auth header with a username and password
type DestinationElasticsearchAuthenticationMethodUsernamePassword struct {
	Method DestinationElasticsearchAuthenticationMethodUsernamePasswordMethod `json:"method"`
	// Basic auth password to access a secure Elasticsearch server
	Password string `json:"password"`
	// Basic auth username to access a secure Elasticsearch server
	Username string `json:"username"`
}

type DestinationElasticsearchAuthenticationMethodAPIKeySecretMethod string

const (
	DestinationElasticsearchAuthenticationMethodAPIKeySecretMethodSecret DestinationElasticsearchAuthenticationMethodAPIKeySecretMethod = "secret"
)

func (e DestinationElasticsearchAuthenticationMethodAPIKeySecretMethod) ToPointer() *DestinationElasticsearchAuthenticationMethodAPIKeySecretMethod {
	return &e
}

func (e *DestinationElasticsearchAuthenticationMethodAPIKeySecretMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "secret":
		*e = DestinationElasticsearchAuthenticationMethodAPIKeySecretMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationElasticsearchAuthenticationMethodAPIKeySecretMethod: %v", v)
	}
}

// DestinationElasticsearchAuthenticationMethodAPIKeySecret - Use a api key and secret combination to authenticate
type DestinationElasticsearchAuthenticationMethodAPIKeySecret struct {
	// The Key ID to used when accessing an enterprise Elasticsearch instance.
	APIKeyID string `json:"apiKeyId"`
	// The secret associated with the API Key ID.
	APIKeySecret string                                                         `json:"apiKeySecret"`
	Method       DestinationElasticsearchAuthenticationMethodAPIKeySecretMethod `json:"method"`
}

type DestinationElasticsearchAuthenticationMethodType string

const (
	DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchAuthenticationMethodAPIKeySecret     DestinationElasticsearchAuthenticationMethodType = "destination-elasticsearch_Authentication Method_Api Key/Secret"
	DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchAuthenticationMethodUsernamePassword DestinationElasticsearchAuthenticationMethodType = "destination-elasticsearch_Authentication Method_Username/Password"
)

type DestinationElasticsearchAuthenticationMethod struct {
	DestinationElasticsearchAuthenticationMethodAPIKeySecret     *DestinationElasticsearchAuthenticationMethodAPIKeySecret
	DestinationElasticsearchAuthenticationMethodUsernamePassword *DestinationElasticsearchAuthenticationMethodUsernamePassword

	Type DestinationElasticsearchAuthenticationMethodType
}

func CreateDestinationElasticsearchAuthenticationMethodDestinationElasticsearchAuthenticationMethodAPIKeySecret(destinationElasticsearchAuthenticationMethodAPIKeySecret DestinationElasticsearchAuthenticationMethodAPIKeySecret) DestinationElasticsearchAuthenticationMethod {
	typ := DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchAuthenticationMethodAPIKeySecret

	return DestinationElasticsearchAuthenticationMethod{
		DestinationElasticsearchAuthenticationMethodAPIKeySecret: &destinationElasticsearchAuthenticationMethodAPIKeySecret,
		Type: typ,
	}
}

func CreateDestinationElasticsearchAuthenticationMethodDestinationElasticsearchAuthenticationMethodUsernamePassword(destinationElasticsearchAuthenticationMethodUsernamePassword DestinationElasticsearchAuthenticationMethodUsernamePassword) DestinationElasticsearchAuthenticationMethod {
	typ := DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchAuthenticationMethodUsernamePassword

	return DestinationElasticsearchAuthenticationMethod{
		DestinationElasticsearchAuthenticationMethodUsernamePassword: &destinationElasticsearchAuthenticationMethodUsernamePassword,
		Type: typ,
	}
}

func (u *DestinationElasticsearchAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationElasticsearchAuthenticationMethodAPIKeySecret := new(DestinationElasticsearchAuthenticationMethodAPIKeySecret)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationElasticsearchAuthenticationMethodAPIKeySecret); err == nil {
		u.DestinationElasticsearchAuthenticationMethodAPIKeySecret = destinationElasticsearchAuthenticationMethodAPIKeySecret
		u.Type = DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchAuthenticationMethodAPIKeySecret
		return nil
	}

	destinationElasticsearchAuthenticationMethodUsernamePassword := new(DestinationElasticsearchAuthenticationMethodUsernamePassword)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationElasticsearchAuthenticationMethodUsernamePassword); err == nil {
		u.DestinationElasticsearchAuthenticationMethodUsernamePassword = destinationElasticsearchAuthenticationMethodUsernamePassword
		u.Type = DestinationElasticsearchAuthenticationMethodTypeDestinationElasticsearchAuthenticationMethodUsernamePassword
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationElasticsearchAuthenticationMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationElasticsearchAuthenticationMethodAPIKeySecret != nil {
		return json.Marshal(u.DestinationElasticsearchAuthenticationMethodAPIKeySecret)
	}

	if u.DestinationElasticsearchAuthenticationMethodUsernamePassword != nil {
		return json.Marshal(u.DestinationElasticsearchAuthenticationMethodUsernamePassword)
	}

	return nil, nil
}

type DestinationElasticsearchElasticsearch string

const (
	DestinationElasticsearchElasticsearchElasticsearch DestinationElasticsearchElasticsearch = "elasticsearch"
)

func (e DestinationElasticsearchElasticsearch) ToPointer() *DestinationElasticsearchElasticsearch {
	return &e
}

func (e *DestinationElasticsearchElasticsearch) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "elasticsearch":
		*e = DestinationElasticsearchElasticsearch(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationElasticsearchElasticsearch: %v", v)
	}
}

type DestinationElasticsearch struct {
	// The type of authentication to be used
	AuthenticationMethod *DestinationElasticsearchAuthenticationMethod `json:"authenticationMethod,omitempty"`
	// CA certificate
	CaCertificate   *string                               `json:"ca_certificate,omitempty"`
	DestinationType DestinationElasticsearchElasticsearch `json:"destinationType"`
	// The full url of the Elasticsearch server
	Endpoint string `json:"endpoint"`
	// If a primary key identifier is defined in the source, an upsert will be performed using the primary key value as the elasticsearch doc id. Does not support composite primary keys.
	Upsert *bool `json:"upsert,omitempty"`
}
