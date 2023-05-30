// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationKinesisResourceModel) ToCreateSDKType() *shared.DestinationKinesisCreateRequest {
	accessKey := r.Configuration.AccessKey.ValueString()
	bufferSize := r.Configuration.BufferSize.ValueInt64()
	destinationType := shared.DestinationKinesisKinesis(r.Configuration.DestinationType.ValueString())
	endpoint := r.Configuration.Endpoint.ValueString()
	privateKey := r.Configuration.PrivateKey.ValueString()
	region := r.Configuration.Region.ValueString()
	shardCount := r.Configuration.ShardCount.ValueInt64()
	configuration := shared.DestinationKinesis{
		AccessKey:       accessKey,
		BufferSize:      bufferSize,
		DestinationType: destinationType,
		Endpoint:        endpoint,
		PrivateKey:      privateKey,
		Region:          region,
		ShardCount:      shardCount,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationKinesisCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationKinesisResourceModel) ToDeleteSDKType() *shared.DestinationKinesisCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationKinesisResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
