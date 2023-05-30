// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationPostgresPostgres string

const (
	DestinationPostgresPostgresPostgres DestinationPostgresPostgres = "postgres"
)

func (e DestinationPostgresPostgres) ToPointer() *DestinationPostgresPostgres {
	return &e
}

func (e *DestinationPostgresPostgres) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "postgres":
		*e = DestinationPostgresPostgres(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresPostgres: %v", v)
	}
}

type DestinationPostgresSSLModesVerifyFullMode string

const (
	DestinationPostgresSSLModesVerifyFullModeVerifyFull DestinationPostgresSSLModesVerifyFullMode = "verify-full"
)

func (e DestinationPostgresSSLModesVerifyFullMode) ToPointer() *DestinationPostgresSSLModesVerifyFullMode {
	return &e
}

func (e *DestinationPostgresSSLModesVerifyFullMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify-full":
		*e = DestinationPostgresSSLModesVerifyFullMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesVerifyFullMode: %v", v)
	}
}

// DestinationPostgresSSLModesVerifyFull - Verify-full SSL mode.
type DestinationPostgresSSLModesVerifyFull struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Client certificate
	ClientCertificate string `json:"client_certificate"`
	// Client key
	ClientKey string `json:"client_key"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                                   `json:"client_key_password,omitempty"`
	Mode              DestinationPostgresSSLModesVerifyFullMode `json:"mode"`
}

type DestinationPostgresSSLModesVerifyCaMode string

const (
	DestinationPostgresSSLModesVerifyCaModeVerifyCa DestinationPostgresSSLModesVerifyCaMode = "verify-ca"
)

func (e DestinationPostgresSSLModesVerifyCaMode) ToPointer() *DestinationPostgresSSLModesVerifyCaMode {
	return &e
}

func (e *DestinationPostgresSSLModesVerifyCaMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify-ca":
		*e = DestinationPostgresSSLModesVerifyCaMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesVerifyCaMode: %v", v)
	}
}

// DestinationPostgresSSLModesVerifyCa - Verify-ca SSL mode.
type DestinationPostgresSSLModesVerifyCa struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                                 `json:"client_key_password,omitempty"`
	Mode              DestinationPostgresSSLModesVerifyCaMode `json:"mode"`
}

type DestinationPostgresSSLModesRequireMode string

const (
	DestinationPostgresSSLModesRequireModeRequire DestinationPostgresSSLModesRequireMode = "require"
)

func (e DestinationPostgresSSLModesRequireMode) ToPointer() *DestinationPostgresSSLModesRequireMode {
	return &e
}

func (e *DestinationPostgresSSLModesRequireMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "require":
		*e = DestinationPostgresSSLModesRequireMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesRequireMode: %v", v)
	}
}

// DestinationPostgresSSLModesRequire - Require SSL mode.
type DestinationPostgresSSLModesRequire struct {
	Mode DestinationPostgresSSLModesRequireMode `json:"mode"`
}

type DestinationPostgresSSLModesPreferMode string

const (
	DestinationPostgresSSLModesPreferModePrefer DestinationPostgresSSLModesPreferMode = "prefer"
)

func (e DestinationPostgresSSLModesPreferMode) ToPointer() *DestinationPostgresSSLModesPreferMode {
	return &e
}

func (e *DestinationPostgresSSLModesPreferMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "prefer":
		*e = DestinationPostgresSSLModesPreferMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesPreferMode: %v", v)
	}
}

// DestinationPostgresSSLModesPrefer - Prefer SSL mode.
type DestinationPostgresSSLModesPrefer struct {
	Mode DestinationPostgresSSLModesPreferMode `json:"mode"`
}

type DestinationPostgresSSLModesAllowMode string

const (
	DestinationPostgresSSLModesAllowModeAllow DestinationPostgresSSLModesAllowMode = "allow"
)

func (e DestinationPostgresSSLModesAllowMode) ToPointer() *DestinationPostgresSSLModesAllowMode {
	return &e
}

func (e *DestinationPostgresSSLModesAllowMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "allow":
		*e = DestinationPostgresSSLModesAllowMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesAllowMode: %v", v)
	}
}

// DestinationPostgresSSLModesAllow - Allow SSL mode.
type DestinationPostgresSSLModesAllow struct {
	Mode DestinationPostgresSSLModesAllowMode `json:"mode"`
}

type DestinationPostgresSSLModesDisableMode string

const (
	DestinationPostgresSSLModesDisableModeDisable DestinationPostgresSSLModesDisableMode = "disable"
)

func (e DestinationPostgresSSLModesDisableMode) ToPointer() *DestinationPostgresSSLModesDisableMode {
	return &e
}

func (e *DestinationPostgresSSLModesDisableMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disable":
		*e = DestinationPostgresSSLModesDisableMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesDisableMode: %v", v)
	}
}

// DestinationPostgresSSLModesDisable - Disable SSL.
type DestinationPostgresSSLModesDisable struct {
	Mode DestinationPostgresSSLModesDisableMode `json:"mode"`
}

type DestinationPostgresSSLModesType string

const (
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesDisable    DestinationPostgresSSLModesType = "destination-postgres_SSL modes_disable"
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesAllow      DestinationPostgresSSLModesType = "destination-postgres_SSL modes_allow"
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesPrefer     DestinationPostgresSSLModesType = "destination-postgres_SSL modes_prefer"
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesRequire    DestinationPostgresSSLModesType = "destination-postgres_SSL modes_require"
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyCa   DestinationPostgresSSLModesType = "destination-postgres_SSL modes_verify-ca"
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyFull DestinationPostgresSSLModesType = "destination-postgres_SSL modes_verify-full"
)

type DestinationPostgresSSLModes struct {
	DestinationPostgresSSLModesDisable    *DestinationPostgresSSLModesDisable
	DestinationPostgresSSLModesAllow      *DestinationPostgresSSLModesAllow
	DestinationPostgresSSLModesPrefer     *DestinationPostgresSSLModesPrefer
	DestinationPostgresSSLModesRequire    *DestinationPostgresSSLModesRequire
	DestinationPostgresSSLModesVerifyCa   *DestinationPostgresSSLModesVerifyCa
	DestinationPostgresSSLModesVerifyFull *DestinationPostgresSSLModesVerifyFull

	Type DestinationPostgresSSLModesType
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesDisable(destinationPostgresSSLModesDisable DestinationPostgresSSLModesDisable) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesDisable

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesDisable: &destinationPostgresSSLModesDisable,
		Type:                               typ,
	}
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesAllow(destinationPostgresSSLModesAllow DestinationPostgresSSLModesAllow) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesAllow

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesAllow: &destinationPostgresSSLModesAllow,
		Type:                             typ,
	}
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesPrefer(destinationPostgresSSLModesPrefer DestinationPostgresSSLModesPrefer) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesPrefer

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesPrefer: &destinationPostgresSSLModesPrefer,
		Type:                              typ,
	}
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesRequire(destinationPostgresSSLModesRequire DestinationPostgresSSLModesRequire) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesRequire

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesRequire: &destinationPostgresSSLModesRequire,
		Type:                               typ,
	}
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesVerifyCa(destinationPostgresSSLModesVerifyCa DestinationPostgresSSLModesVerifyCa) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyCa

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesVerifyCa: &destinationPostgresSSLModesVerifyCa,
		Type:                                typ,
	}
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesVerifyFull(destinationPostgresSSLModesVerifyFull DestinationPostgresSSLModesVerifyFull) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyFull

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesVerifyFull: &destinationPostgresSSLModesVerifyFull,
		Type:                                  typ,
	}
}

func (u *DestinationPostgresSSLModes) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationPostgresSSLModesDisable := new(DestinationPostgresSSLModesDisable)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSLModesDisable); err == nil {
		u.DestinationPostgresSSLModesDisable = destinationPostgresSSLModesDisable
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesDisable
		return nil
	}

	destinationPostgresSSLModesAllow := new(DestinationPostgresSSLModesAllow)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSLModesAllow); err == nil {
		u.DestinationPostgresSSLModesAllow = destinationPostgresSSLModesAllow
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesAllow
		return nil
	}

	destinationPostgresSSLModesPrefer := new(DestinationPostgresSSLModesPrefer)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSLModesPrefer); err == nil {
		u.DestinationPostgresSSLModesPrefer = destinationPostgresSSLModesPrefer
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesPrefer
		return nil
	}

	destinationPostgresSSLModesRequire := new(DestinationPostgresSSLModesRequire)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSLModesRequire); err == nil {
		u.DestinationPostgresSSLModesRequire = destinationPostgresSSLModesRequire
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesRequire
		return nil
	}

	destinationPostgresSSLModesVerifyCa := new(DestinationPostgresSSLModesVerifyCa)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSLModesVerifyCa); err == nil {
		u.DestinationPostgresSSLModesVerifyCa = destinationPostgresSSLModesVerifyCa
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyCa
		return nil
	}

	destinationPostgresSSLModesVerifyFull := new(DestinationPostgresSSLModesVerifyFull)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSLModesVerifyFull); err == nil {
		u.DestinationPostgresSSLModesVerifyFull = destinationPostgresSSLModesVerifyFull
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyFull
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationPostgresSSLModes) MarshalJSON() ([]byte, error) {
	if u.DestinationPostgresSSLModesDisable != nil {
		return json.Marshal(u.DestinationPostgresSSLModesDisable)
	}

	if u.DestinationPostgresSSLModesAllow != nil {
		return json.Marshal(u.DestinationPostgresSSLModesAllow)
	}

	if u.DestinationPostgresSSLModesPrefer != nil {
		return json.Marshal(u.DestinationPostgresSSLModesPrefer)
	}

	if u.DestinationPostgresSSLModesRequire != nil {
		return json.Marshal(u.DestinationPostgresSSLModesRequire)
	}

	if u.DestinationPostgresSSLModesVerifyCa != nil {
		return json.Marshal(u.DestinationPostgresSSLModesVerifyCa)
	}

	if u.DestinationPostgresSSLModesVerifyFull != nil {
		return json.Marshal(u.DestinationPostgresSSLModesVerifyFull)
	}

	return nil, nil
}

// DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod string

const (
	DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod) ToPointer() *DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationPostgresSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationPostgresSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	TunnelMethod DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

// DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod string

const (
	DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) ToPointer() *DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationPostgresSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationPostgresSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	TunnelMethod DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

// DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod - No ssh tunnel needed to connect to database
type DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod string

const (
	DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethodNoTunnel DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod = "NO_TUNNEL"
)

func (e DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod) ToPointer() *DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod {
	return &e
}

func (e *DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod: %v", v)
	}
}

// DestinationPostgresSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationPostgresSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	TunnelMethod DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod `json:"tunnel_method"`
}

type DestinationPostgresSSHTunnelMethodType string

const (
	DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodNoTunnel               DestinationPostgresSSHTunnelMethodType = "destination-postgres_SSH Tunnel Method_No Tunnel"
	DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodSSHKeyAuthentication   DestinationPostgresSSHTunnelMethodType = "destination-postgres_SSH Tunnel Method_SSH Key Authentication"
	DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodPasswordAuthentication DestinationPostgresSSHTunnelMethodType = "destination-postgres_SSH Tunnel Method_Password Authentication"
)

type DestinationPostgresSSHTunnelMethod struct {
	DestinationPostgresSSHTunnelMethodNoTunnel               *DestinationPostgresSSHTunnelMethodNoTunnel
	DestinationPostgresSSHTunnelMethodSSHKeyAuthentication   *DestinationPostgresSSHTunnelMethodSSHKeyAuthentication
	DestinationPostgresSSHTunnelMethodPasswordAuthentication *DestinationPostgresSSHTunnelMethodPasswordAuthentication

	Type DestinationPostgresSSHTunnelMethodType
}

func CreateDestinationPostgresSSHTunnelMethodDestinationPostgresSSHTunnelMethodNoTunnel(destinationPostgresSSHTunnelMethodNoTunnel DestinationPostgresSSHTunnelMethodNoTunnel) DestinationPostgresSSHTunnelMethod {
	typ := DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodNoTunnel

	return DestinationPostgresSSHTunnelMethod{
		DestinationPostgresSSHTunnelMethodNoTunnel: &destinationPostgresSSHTunnelMethodNoTunnel,
		Type: typ,
	}
}

func CreateDestinationPostgresSSHTunnelMethodDestinationPostgresSSHTunnelMethodSSHKeyAuthentication(destinationPostgresSSHTunnelMethodSSHKeyAuthentication DestinationPostgresSSHTunnelMethodSSHKeyAuthentication) DestinationPostgresSSHTunnelMethod {
	typ := DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodSSHKeyAuthentication

	return DestinationPostgresSSHTunnelMethod{
		DestinationPostgresSSHTunnelMethodSSHKeyAuthentication: &destinationPostgresSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationPostgresSSHTunnelMethodDestinationPostgresSSHTunnelMethodPasswordAuthentication(destinationPostgresSSHTunnelMethodPasswordAuthentication DestinationPostgresSSHTunnelMethodPasswordAuthentication) DestinationPostgresSSHTunnelMethod {
	typ := DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodPasswordAuthentication

	return DestinationPostgresSSHTunnelMethod{
		DestinationPostgresSSHTunnelMethodPasswordAuthentication: &destinationPostgresSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationPostgresSSHTunnelMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationPostgresSSHTunnelMethodNoTunnel := new(DestinationPostgresSSHTunnelMethodNoTunnel)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSHTunnelMethodNoTunnel); err == nil {
		u.DestinationPostgresSSHTunnelMethodNoTunnel = destinationPostgresSSHTunnelMethodNoTunnel
		u.Type = DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodNoTunnel
		return nil
	}

	destinationPostgresSSHTunnelMethodSSHKeyAuthentication := new(DestinationPostgresSSHTunnelMethodSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSHTunnelMethodSSHKeyAuthentication); err == nil {
		u.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication = destinationPostgresSSHTunnelMethodSSHKeyAuthentication
		u.Type = DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	destinationPostgresSSHTunnelMethodPasswordAuthentication := new(DestinationPostgresSSHTunnelMethodPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationPostgresSSHTunnelMethodPasswordAuthentication); err == nil {
		u.DestinationPostgresSSHTunnelMethodPasswordAuthentication = destinationPostgresSSHTunnelMethodPasswordAuthentication
		u.Type = DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationPostgresSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationPostgresSSHTunnelMethodNoTunnel != nil {
		return json.Marshal(u.DestinationPostgresSSHTunnelMethodNoTunnel)
	}

	if u.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication != nil {
		return json.Marshal(u.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication)
	}

	if u.DestinationPostgresSSHTunnelMethodPasswordAuthentication != nil {
		return json.Marshal(u.DestinationPostgresSSHTunnelMethodPasswordAuthentication)
	}

	return nil, nil
}

type DestinationPostgres struct {
	// Name of the database.
	Database        string                      `json:"database"`
	DestinationType DestinationPostgresPostgres `json:"destinationType"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	Port int64 `json:"port"`
	// The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public".
	Schema string `json:"schema"`
	// SSL connection modes.
	//  <b>disable</b> - Chose this mode to disable encryption of communication between Airbyte and destination database
	//  <b>allow</b> - Chose this mode to enable encryption only when required by the source database
	//  <b>prefer</b> - Chose this mode to allow unencrypted connection only if the source database does not support encryption
	//  <b>require</b> - Chose this mode to always require encryption. If the source database server does not support encryption, connection will fail
	//   <b>verify-ca</b> - Chose this mode to always require encryption and to verify that the source database server has a valid SSL certificate
	//   <b>verify-full</b> - This is the most secure mode. Chose this mode to always require encryption and to verify the identity of the source database server
	//  See more information - <a href="https://jdbc.postgresql.org/documentation/head/ssl-client.html"> in the docs</a>.
	SslMode *DestinationPostgresSSLModes `json:"ssl_mode,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationPostgresSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}
