// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	customTypes "airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceGoogleSearchConsoleResourceModel) ToCreateSDKType() *shared.SourceGoogleSearchConsoleCreateRequest {
	var authorization shared.SourceGoogleSearchConsoleAuthenticationType
	var sourceGoogleSearchConsoleAuthenticationTypeOAuth *shared.SourceGoogleSearchConsoleAuthenticationTypeOAuth
	if r.Configuration.Authorization.SourceGoogleSearchConsoleAuthenticationTypeOAuth != nil {
		accessToken := new(string)
		if !r.Configuration.Authorization.SourceGoogleSearchConsoleAuthenticationTypeOAuth.AccessToken.IsUnknown() && !r.Configuration.Authorization.SourceGoogleSearchConsoleAuthenticationTypeOAuth.AccessToken.IsNull() {
			*accessToken = r.Configuration.Authorization.SourceGoogleSearchConsoleAuthenticationTypeOAuth.AccessToken.ValueString()
		} else {
			accessToken = nil
		}
		authType := shared.SourceGoogleSearchConsoleAuthenticationTypeOAuthAuthType(r.Configuration.Authorization.SourceGoogleSearchConsoleAuthenticationTypeOAuth.AuthType.ValueString())
		clientID := r.Configuration.Authorization.SourceGoogleSearchConsoleAuthenticationTypeOAuth.ClientID.ValueString()
		clientSecret := r.Configuration.Authorization.SourceGoogleSearchConsoleAuthenticationTypeOAuth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Authorization.SourceGoogleSearchConsoleAuthenticationTypeOAuth.RefreshToken.ValueString()
		sourceGoogleSearchConsoleAuthenticationTypeOAuth = &shared.SourceGoogleSearchConsoleAuthenticationTypeOAuth{
			AccessToken:  accessToken,
			AuthType:     authType,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
		}
	}
	if sourceGoogleSearchConsoleAuthenticationTypeOAuth != nil {
		authorization = shared.SourceGoogleSearchConsoleAuthenticationType{
			SourceGoogleSearchConsoleAuthenticationTypeOAuth: sourceGoogleSearchConsoleAuthenticationTypeOAuth,
		}
	}
	var sourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication *shared.SourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication
	if r.Configuration.Authorization.SourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication != nil {
		authType1 := shared.SourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthenticationAuthType(r.Configuration.Authorization.SourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication.AuthType.ValueString())
		email := r.Configuration.Authorization.SourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication.Email.ValueString()
		serviceAccountInfo := r.Configuration.Authorization.SourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication.ServiceAccountInfo.ValueString()
		sourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication = &shared.SourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication{
			AuthType:           authType1,
			Email:              email,
			ServiceAccountInfo: serviceAccountInfo,
		}
	}
	if sourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication != nil {
		authorization = shared.SourceGoogleSearchConsoleAuthenticationType{
			SourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication: sourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication,
		}
	}
	customReports := new(string)
	if !r.Configuration.CustomReports.IsUnknown() && !r.Configuration.CustomReports.IsNull() {
		*customReports = r.Configuration.CustomReports.ValueString()
	} else {
		customReports = nil
	}
	endDate := new(customTypes.Date)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		endDate = customTypes.MustNewDateFromString(r.Configuration.EndDate.ValueString())
	} else {
		endDate = nil
	}
	siteUrls := make([]string, 0)
	for _, siteUrlsItem := range r.Configuration.SiteUrls {
		siteUrls = append(siteUrls, siteUrlsItem.ValueString())
	}
	sourceType := shared.SourceGoogleSearchConsoleGoogleSearchConsole(r.Configuration.SourceType.ValueString())
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	configuration := shared.SourceGoogleSearchConsole{
		Authorization: authorization,
		CustomReports: customReports,
		EndDate:       endDate,
		SiteUrls:      siteUrls,
		SourceType:    sourceType,
		StartDate:     startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleSearchConsoleCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleSearchConsoleResourceModel) ToDeleteSDKType() *shared.SourceGoogleSearchConsoleCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGoogleSearchConsoleResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
