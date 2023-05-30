// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourcePaystackResourceModel) ToCreateSDKType() *shared.SourcePaystackCreateRequest {
	lookbackWindowDays := new(int64)
	if !r.Configuration.LookbackWindowDays.IsUnknown() && !r.Configuration.LookbackWindowDays.IsNull() {
		*lookbackWindowDays = r.Configuration.LookbackWindowDays.ValueInt64()
	} else {
		lookbackWindowDays = nil
	}
	secretKey := r.Configuration.SecretKey.ValueString()
	sourceType := shared.SourcePaystackPaystack(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourcePaystack{
		LookbackWindowDays: lookbackWindowDays,
		SecretKey:          secretKey,
		SourceType:         sourceType,
		StartDate:          startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePaystackCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePaystackResourceModel) ToDeleteSDKType() *shared.SourcePaystackCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePaystackResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
