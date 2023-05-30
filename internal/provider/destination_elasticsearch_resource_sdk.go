// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationElasticsearchResourceModel) ToCreateSDKType() *shared.DestinationElasticsearchCreateRequest {
	var authenticationMethod *shared.DestinationElasticsearchAuthenticationMethod
	var destinationElasticsearchAuthenticationMethodAPIKeySecret *shared.DestinationElasticsearchAuthenticationMethodAPIKeySecret
	if r.Configuration.AuthenticationMethod.DestinationElasticsearchAuthenticationMethodAPIKeySecret != nil {
		apiKeyID := r.Configuration.AuthenticationMethod.DestinationElasticsearchAuthenticationMethodAPIKeySecret.APIKeyID.ValueString()
		apiKeySecret := r.Configuration.AuthenticationMethod.DestinationElasticsearchAuthenticationMethodAPIKeySecret.APIKeySecret.ValueString()
		method := shared.DestinationElasticsearchAuthenticationMethodAPIKeySecretMethod(r.Configuration.AuthenticationMethod.DestinationElasticsearchAuthenticationMethodAPIKeySecret.Method.ValueString())
		destinationElasticsearchAuthenticationMethodAPIKeySecret = &shared.DestinationElasticsearchAuthenticationMethodAPIKeySecret{
			APIKeyID:     apiKeyID,
			APIKeySecret: apiKeySecret,
			Method:       method,
		}
	}
	if destinationElasticsearchAuthenticationMethodAPIKeySecret != nil {
		authenticationMethod = &shared.DestinationElasticsearchAuthenticationMethod{
			DestinationElasticsearchAuthenticationMethodAPIKeySecret: destinationElasticsearchAuthenticationMethodAPIKeySecret,
		}
	}
	var destinationElasticsearchAuthenticationMethodUsernamePassword *shared.DestinationElasticsearchAuthenticationMethodUsernamePassword
	if r.Configuration.AuthenticationMethod.DestinationElasticsearchAuthenticationMethodUsernamePassword != nil {
		method1 := shared.DestinationElasticsearchAuthenticationMethodUsernamePasswordMethod(r.Configuration.AuthenticationMethod.DestinationElasticsearchAuthenticationMethodUsernamePassword.Method.ValueString())
		password := r.Configuration.AuthenticationMethod.DestinationElasticsearchAuthenticationMethodUsernamePassword.Password.ValueString()
		username := r.Configuration.AuthenticationMethod.DestinationElasticsearchAuthenticationMethodUsernamePassword.Username.ValueString()
		destinationElasticsearchAuthenticationMethodUsernamePassword = &shared.DestinationElasticsearchAuthenticationMethodUsernamePassword{
			Method:   method1,
			Password: password,
			Username: username,
		}
	}
	if destinationElasticsearchAuthenticationMethodUsernamePassword != nil {
		authenticationMethod = &shared.DestinationElasticsearchAuthenticationMethod{
			DestinationElasticsearchAuthenticationMethodUsernamePassword: destinationElasticsearchAuthenticationMethodUsernamePassword,
		}
	}
	caCertificate := new(string)
	if !r.Configuration.CaCertificate.IsUnknown() && !r.Configuration.CaCertificate.IsNull() {
		*caCertificate = r.Configuration.CaCertificate.ValueString()
	} else {
		caCertificate = nil
	}
	destinationType := shared.DestinationElasticsearchElasticsearch(r.Configuration.DestinationType.ValueString())
	endpoint := r.Configuration.Endpoint.ValueString()
	upsert := new(bool)
	if !r.Configuration.Upsert.IsUnknown() && !r.Configuration.Upsert.IsNull() {
		*upsert = r.Configuration.Upsert.ValueBool()
	} else {
		upsert = nil
	}
	configuration := shared.DestinationElasticsearch{
		AuthenticationMethod: authenticationMethod,
		CaCertificate:        caCertificate,
		DestinationType:      destinationType,
		Endpoint:             endpoint,
		Upsert:               upsert,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationElasticsearchCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationElasticsearchResourceModel) ToDeleteSDKType() *shared.DestinationElasticsearchCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationElasticsearchResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
