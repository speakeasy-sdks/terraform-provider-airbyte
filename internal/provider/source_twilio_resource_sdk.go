// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceTwilioResourceModel) ToCreateSDKType() *shared.SourceTwilioCreateRequest {
	accountSid := r.Configuration.AccountSid.ValueString()
	authToken := r.Configuration.AuthToken.ValueString()
	lookbackWindow := new(int64)
	if !r.Configuration.LookbackWindow.IsUnknown() && !r.Configuration.LookbackWindow.IsNull() {
		*lookbackWindow = r.Configuration.LookbackWindow.ValueInt64()
	} else {
		lookbackWindow = nil
	}
	sourceType := shared.SourceTwilioTwilio(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceTwilio{
		AccountSid:     accountSid,
		AuthToken:      authToken,
		LookbackWindow: lookbackWindow,
		SourceType:     sourceType,
		StartDate:      startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTwilioCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTwilioResourceModel) ToDeleteSDKType() *shared.SourceTwilioCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTwilioResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
