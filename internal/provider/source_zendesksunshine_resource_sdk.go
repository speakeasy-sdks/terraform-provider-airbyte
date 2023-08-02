// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceZendeskSunshineResourceModel) ToCreateSDKType() *shared.SourceZendeskSunshineCreateRequest {
	var credentials *shared.SourceZendeskSunshineAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceZendeskSunshineAuthorizationMethodOAuth20 *shared.SourceZendeskSunshineAuthorizationMethodOAuth20
		if r.Configuration.Credentials.SourceZendeskSunshineAuthorizationMethodAPIToken != nil {
			authMethod := shared.SourceZendeskSunshineAuthorizationMethodOAuth20AuthMethod(r.Configuration.Credentials.SourceZendeskSunshineAuthorizationMethodAPIToken.AuthMethod.ValueString())
			sourceZendeskSunshineAuthorizationMethodOAuth20 = &shared.SourceZendeskSunshineAuthorizationMethodOAuth20{
				AuthMethod: authMethod,
			}
		}
		if sourceZendeskSunshineAuthorizationMethodOAuth20 != nil {
			credentials = &shared.SourceZendeskSunshineAuthorizationMethod{
				SourceZendeskSunshineAuthorizationMethodOAuth20: sourceZendeskSunshineAuthorizationMethodOAuth20,
			}
		}
		var sourceZendeskSunshineAuthorizationMethodAPIToken *shared.SourceZendeskSunshineAuthorizationMethodAPIToken
		if r.Configuration.Credentials.SourceZendeskSunshineAuthorizationMethodOAuth20 != nil {
			authMethod1 := shared.SourceZendeskSunshineAuthorizationMethodAPITokenAuthMethod(r.Configuration.Credentials.SourceZendeskSunshineAuthorizationMethodOAuth20.AuthMethod.ValueString())
			sourceZendeskSunshineAuthorizationMethodAPIToken = &shared.SourceZendeskSunshineAuthorizationMethodAPIToken{
				AuthMethod: authMethod1,
			}
		}
		if sourceZendeskSunshineAuthorizationMethodAPIToken != nil {
			credentials = &shared.SourceZendeskSunshineAuthorizationMethod{
				SourceZendeskSunshineAuthorizationMethodAPIToken: sourceZendeskSunshineAuthorizationMethodAPIToken,
			}
		}
	}
	sourceType := shared.SourceZendeskSunshineZendeskSunshine(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	subdomain := r.Configuration.Subdomain.ValueString()
	configuration := shared.SourceZendeskSunshine{
		Credentials: credentials,
		SourceType:  sourceType,
		StartDate:   startDate,
		Subdomain:   subdomain,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZendeskSunshineCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZendeskSunshineResourceModel) ToGetSDKType() *shared.SourceZendeskSunshineCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZendeskSunshineResourceModel) ToUpdateSDKType() *shared.SourceZendeskSunshinePutRequest {
	var credentials *shared.SourceZendeskSunshineUpdateAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceZendeskSunshineUpdateAuthorizationMethodOAuth20 *shared.SourceZendeskSunshineUpdateAuthorizationMethodOAuth20
		if r.Configuration.Credentials.SourceZendeskSunshineAuthorizationMethodAPIToken != nil {
			authMethod := shared.SourceZendeskSunshineUpdateAuthorizationMethodOAuth20AuthMethod(r.Configuration.Credentials.SourceZendeskSunshineAuthorizationMethodAPIToken.AuthMethod.ValueString())
			sourceZendeskSunshineUpdateAuthorizationMethodOAuth20 = &shared.SourceZendeskSunshineUpdateAuthorizationMethodOAuth20{
				AuthMethod: authMethod,
			}
		}
		if sourceZendeskSunshineUpdateAuthorizationMethodOAuth20 != nil {
			credentials = &shared.SourceZendeskSunshineUpdateAuthorizationMethod{
				SourceZendeskSunshineUpdateAuthorizationMethodOAuth20: sourceZendeskSunshineUpdateAuthorizationMethodOAuth20,
			}
		}
		var sourceZendeskSunshineUpdateAuthorizationMethodAPIToken *shared.SourceZendeskSunshineUpdateAuthorizationMethodAPIToken
		if r.Configuration.Credentials.SourceZendeskSunshineAuthorizationMethodOAuth20 != nil {
			authMethod1 := shared.SourceZendeskSunshineUpdateAuthorizationMethodAPITokenAuthMethod(r.Configuration.Credentials.SourceZendeskSunshineAuthorizationMethodOAuth20.AuthMethod.ValueString())
			sourceZendeskSunshineUpdateAuthorizationMethodAPIToken = &shared.SourceZendeskSunshineUpdateAuthorizationMethodAPIToken{
				AuthMethod: authMethod1,
			}
		}
		if sourceZendeskSunshineUpdateAuthorizationMethodAPIToken != nil {
			credentials = &shared.SourceZendeskSunshineUpdateAuthorizationMethod{
				SourceZendeskSunshineUpdateAuthorizationMethodAPIToken: sourceZendeskSunshineUpdateAuthorizationMethodAPIToken,
			}
		}
	}
	startDate := r.Configuration.StartDate.ValueString()
	subdomain := r.Configuration.Subdomain.ValueString()
	configuration := shared.SourceZendeskSunshineUpdate{
		Credentials: credentials,
		StartDate:   startDate,
		Subdomain:   subdomain,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZendeskSunshinePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZendeskSunshineResourceModel) ToDeleteSDKType() *shared.SourceZendeskSunshineCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZendeskSunshineResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceZendeskSunshineResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
