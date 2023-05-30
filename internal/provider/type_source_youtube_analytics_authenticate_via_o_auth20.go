// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceYoutubeAnalyticsAuthenticateViaOAuth20 struct {
	ClientID             types.String `tfsdk:"client_id"`
	ClientSecret         types.String `tfsdk:"client_secret"`
	RefreshToken         types.String `tfsdk:"refresh_token"`
	AdditionalProperties types.String `tfsdk:"additional_properties"`
}
