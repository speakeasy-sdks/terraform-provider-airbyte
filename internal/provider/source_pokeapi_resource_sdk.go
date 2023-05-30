// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourcePokeapiResourceModel) ToCreateSDKType() *shared.SourcePokeapiCreateRequest {
	pokemonName := r.Configuration.PokemonName.ValueString()
	sourceType := shared.SourcePokeapiPokeapi(r.Configuration.SourceType.ValueString())
	configuration := shared.SourcePokeapi{
		PokemonName: pokemonName,
		SourceType:  sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePokeapiCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePokeapiResourceModel) ToDeleteSDKType() *shared.SourcePokeapiCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePokeapiResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
