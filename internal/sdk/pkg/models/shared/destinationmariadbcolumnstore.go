// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationMariadbColumnstoreMariadbColumnstore string

const (
	DestinationMariadbColumnstoreMariadbColumnstoreMariadbColumnstore DestinationMariadbColumnstoreMariadbColumnstore = "mariadb-columnstore"
)

func (e DestinationMariadbColumnstoreMariadbColumnstore) ToPointer() *DestinationMariadbColumnstoreMariadbColumnstore {
	return &e
}

func (e *DestinationMariadbColumnstoreMariadbColumnstore) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mariadb-columnstore":
		*e = DestinationMariadbColumnstoreMariadbColumnstore(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMariadbColumnstoreMariadbColumnstore: %v", v)
	}
}

// DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthenticationTunnelMethod string

const (
	DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthenticationTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthenticationTunnelMethod) ToPointer() *DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	TunnelMethod DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

// DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthenticationTunnelMethod string

const (
	DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthenticationTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) ToPointer() *DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	TunnelMethod DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

// DestinationMariadbColumnstoreSSHTunnelMethodNoTunnelTunnelMethod - No ssh tunnel needed to connect to database
type DestinationMariadbColumnstoreSSHTunnelMethodNoTunnelTunnelMethod string

const (
	DestinationMariadbColumnstoreSSHTunnelMethodNoTunnelTunnelMethodNoTunnel DestinationMariadbColumnstoreSSHTunnelMethodNoTunnelTunnelMethod = "NO_TUNNEL"
)

func (e DestinationMariadbColumnstoreSSHTunnelMethodNoTunnelTunnelMethod) ToPointer() *DestinationMariadbColumnstoreSSHTunnelMethodNoTunnelTunnelMethod {
	return &e
}

func (e *DestinationMariadbColumnstoreSSHTunnelMethodNoTunnelTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationMariadbColumnstoreSSHTunnelMethodNoTunnelTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMariadbColumnstoreSSHTunnelMethodNoTunnelTunnelMethod: %v", v)
	}
}

// DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	TunnelMethod DestinationMariadbColumnstoreSSHTunnelMethodNoTunnelTunnelMethod `json:"tunnel_method"`
}

type DestinationMariadbColumnstoreSSHTunnelMethodType string

const (
	DestinationMariadbColumnstoreSSHTunnelMethodTypeDestinationMariadbColumnstoreSSHTunnelMethodNoTunnel               DestinationMariadbColumnstoreSSHTunnelMethodType = "destination-mariadb-columnstore_SSH Tunnel Method_No Tunnel"
	DestinationMariadbColumnstoreSSHTunnelMethodTypeDestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication   DestinationMariadbColumnstoreSSHTunnelMethodType = "destination-mariadb-columnstore_SSH Tunnel Method_SSH Key Authentication"
	DestinationMariadbColumnstoreSSHTunnelMethodTypeDestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication DestinationMariadbColumnstoreSSHTunnelMethodType = "destination-mariadb-columnstore_SSH Tunnel Method_Password Authentication"
)

type DestinationMariadbColumnstoreSSHTunnelMethod struct {
	DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel               *DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel
	DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication   *DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication
	DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication *DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication

	Type DestinationMariadbColumnstoreSSHTunnelMethodType
}

func CreateDestinationMariadbColumnstoreSSHTunnelMethodDestinationMariadbColumnstoreSSHTunnelMethodNoTunnel(destinationMariadbColumnstoreSSHTunnelMethodNoTunnel DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel) DestinationMariadbColumnstoreSSHTunnelMethod {
	typ := DestinationMariadbColumnstoreSSHTunnelMethodTypeDestinationMariadbColumnstoreSSHTunnelMethodNoTunnel

	return DestinationMariadbColumnstoreSSHTunnelMethod{
		DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel: &destinationMariadbColumnstoreSSHTunnelMethodNoTunnel,
		Type: typ,
	}
}

func CreateDestinationMariadbColumnstoreSSHTunnelMethodDestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication(destinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication) DestinationMariadbColumnstoreSSHTunnelMethod {
	typ := DestinationMariadbColumnstoreSSHTunnelMethodTypeDestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication

	return DestinationMariadbColumnstoreSSHTunnelMethod{
		DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication: &destinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationMariadbColumnstoreSSHTunnelMethodDestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication(destinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication) DestinationMariadbColumnstoreSSHTunnelMethod {
	typ := DestinationMariadbColumnstoreSSHTunnelMethodTypeDestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication

	return DestinationMariadbColumnstoreSSHTunnelMethod{
		DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication: &destinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationMariadbColumnstoreSSHTunnelMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationMariadbColumnstoreSSHTunnelMethodNoTunnel := new(DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMariadbColumnstoreSSHTunnelMethodNoTunnel); err == nil {
		u.DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel = destinationMariadbColumnstoreSSHTunnelMethodNoTunnel
		u.Type = DestinationMariadbColumnstoreSSHTunnelMethodTypeDestinationMariadbColumnstoreSSHTunnelMethodNoTunnel
		return nil
	}

	destinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication := new(DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication); err == nil {
		u.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication = destinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication
		u.Type = DestinationMariadbColumnstoreSSHTunnelMethodTypeDestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	destinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication := new(DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication); err == nil {
		u.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication = destinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication
		u.Type = DestinationMariadbColumnstoreSSHTunnelMethodTypeDestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMariadbColumnstoreSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel != nil {
		return json.Marshal(u.DestinationMariadbColumnstoreSSHTunnelMethodNoTunnel)
	}

	if u.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication != nil {
		return json.Marshal(u.DestinationMariadbColumnstoreSSHTunnelMethodSSHKeyAuthentication)
	}

	if u.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication != nil {
		return json.Marshal(u.DestinationMariadbColumnstoreSSHTunnelMethodPasswordAuthentication)
	}

	return nil, nil
}

type DestinationMariadbColumnstore struct {
	// Name of the database.
	Database        string                                          `json:"database"`
	DestinationType DestinationMariadbColumnstoreMariadbColumnstore `json:"destinationType"`
	// The Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The Password associated with the username.
	Password *string `json:"password,omitempty"`
	// The Port of the database.
	Port int64 `json:"port"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationMariadbColumnstoreSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The Username which is used to access the database.
	Username string `json:"username"`
}
