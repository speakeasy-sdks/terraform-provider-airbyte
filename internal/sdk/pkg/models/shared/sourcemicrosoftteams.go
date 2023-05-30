// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftAuthType string

const (
	SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftAuthTypeToken SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftAuthType = "Token"
)

func (e SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftAuthType) ToPointer() *SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftAuthType {
	return &e
}

func (e *SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Token":
		*e = SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftAuthType: %v", v)
	}
}

// SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft - Choose how to authenticate to Microsoft
type SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft struct {
	AuthType *SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftAuthType `json:"auth_type,omitempty"`
	// The Client ID of your Microsoft Teams developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Microsoft Teams developer application.
	ClientSecret string `json:"client_secret"`
	// A globally unique identifier (GUID) that is different than your organization name or domain. Follow these steps to obtain: open one of the Teams where you belong inside the Teams Application -> Click on the … next to the Team title -> Click on Get link to team -> Copy the link to the team and grab the tenant ID form the URL
	TenantID string `json:"tenant_id"`
}

type SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20AuthType string

const (
	SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20AuthTypeClient SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20AuthType = "Client"
)

func (e SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20AuthType) ToPointer() *SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20AuthType {
	return &e
}

func (e *SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20AuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20AuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20AuthType: %v", v)
	}
}

// SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 - Choose how to authenticate to Microsoft
type SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 struct {
	AuthType *SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20AuthType `json:"auth_type,omitempty"`
	// The Client ID of your Microsoft Teams developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Microsoft Teams developer application.
	ClientSecret string `json:"client_secret"`
	// A Refresh Token to renew the expired Access Token.
	RefreshToken string `json:"refresh_token"`
	// A globally unique identifier (GUID) that is different than your organization name or domain. Follow these steps to obtain: open one of the Teams where you belong inside the Teams Application -> Click on the … next to the Team title -> Click on Get link to team -> Copy the link to the team and grab the tenant ID form the URL
	TenantID string `json:"tenant_id"`
}

type SourceMicrosoftTeamsAuthenticationMechanismType string

const (
	SourceMicrosoftTeamsAuthenticationMechanismTypeSourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 SourceMicrosoftTeamsAuthenticationMechanismType = "source-microsoft-teams_Authentication mechanism_Authenticate via Microsoft (OAuth 2.0)"
	SourceMicrosoftTeamsAuthenticationMechanismTypeSourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft        SourceMicrosoftTeamsAuthenticationMechanismType = "source-microsoft-teams_Authentication mechanism_Authenticate via Microsoft"
)

type SourceMicrosoftTeamsAuthenticationMechanism struct {
	SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 *SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20
	SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft        *SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft

	Type SourceMicrosoftTeamsAuthenticationMechanismType
}

func CreateSourceMicrosoftTeamsAuthenticationMechanismSourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20(sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20) SourceMicrosoftTeamsAuthenticationMechanism {
	typ := SourceMicrosoftTeamsAuthenticationMechanismTypeSourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20

	return SourceMicrosoftTeamsAuthenticationMechanism{
		SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20: &sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20,
		Type: typ,
	}
}

func CreateSourceMicrosoftTeamsAuthenticationMechanismSourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft(sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft) SourceMicrosoftTeamsAuthenticationMechanism {
	typ := SourceMicrosoftTeamsAuthenticationMechanismTypeSourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft

	return SourceMicrosoftTeamsAuthenticationMechanism{
		SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft: &sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft,
		Type: typ,
	}
}

func (u *SourceMicrosoftTeamsAuthenticationMechanism) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 := new(SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20); err == nil {
		u.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 = sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20
		u.Type = SourceMicrosoftTeamsAuthenticationMechanismTypeSourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20
		return nil
	}

	sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft := new(SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft); err == nil {
		u.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft = sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft
		u.Type = SourceMicrosoftTeamsAuthenticationMechanismTypeSourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMicrosoftTeamsAuthenticationMechanism) MarshalJSON() ([]byte, error) {
	if u.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 != nil {
		return json.Marshal(u.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20)
	}

	if u.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft != nil {
		return json.Marshal(u.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft)
	}

	return nil, nil
}

type SourceMicrosoftTeamsMicrosoftTeams string

const (
	SourceMicrosoftTeamsMicrosoftTeamsMicrosoftTeams SourceMicrosoftTeamsMicrosoftTeams = "microsoft-teams"
)

func (e SourceMicrosoftTeamsMicrosoftTeams) ToPointer() *SourceMicrosoftTeamsMicrosoftTeams {
	return &e
}

func (e *SourceMicrosoftTeamsMicrosoftTeams) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "microsoft-teams":
		*e = SourceMicrosoftTeamsMicrosoftTeams(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMicrosoftTeamsMicrosoftTeams: %v", v)
	}
}

type SourceMicrosoftTeams struct {
	// Choose how to authenticate to Microsoft
	Credentials *SourceMicrosoftTeamsAuthenticationMechanism `json:"credentials,omitempty"`
	// Specifies the length of time over which the Team Device Report stream is aggregated. The supported values are: D7, D30, D90, and D180.
	Period     string                             `json:"period"`
	SourceType SourceMicrosoftTeamsMicrosoftTeams `json:"sourceType"`
}
