// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceMyHoursResourceModel) ToCreateSDKType() *shared.SourceMyHoursCreateRequest {
	email := r.Configuration.Email.ValueString()
	logsBatchSize := new(int64)
	if !r.Configuration.LogsBatchSize.IsUnknown() && !r.Configuration.LogsBatchSize.IsNull() {
		*logsBatchSize = r.Configuration.LogsBatchSize.ValueInt64()
	} else {
		logsBatchSize = nil
	}
	password := r.Configuration.Password.ValueString()
	sourceType := shared.SourceMyHoursMyHours(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceMyHours{
		Email:         email,
		LogsBatchSize: logsBatchSize,
		Password:      password,
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
	out := shared.SourceMyHoursCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMyHoursResourceModel) ToDeleteSDKType() *shared.SourceMyHoursCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMyHoursResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
