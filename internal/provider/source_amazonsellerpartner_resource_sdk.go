// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceAmazonSellerPartnerResourceModel) ToCreateSDKType() *shared.SourceAmazonSellerPartnerCreateRequest {
	appID := r.Configuration.AppID.ValueString()
	authType := new(shared.SourceAmazonSellerPartnerAuthType)
	if !r.Configuration.AuthType.IsUnknown() && !r.Configuration.AuthType.IsNull() {
		*authType = shared.SourceAmazonSellerPartnerAuthType(r.Configuration.AuthType.ValueString())
	} else {
		authType = nil
	}
	awsAccessKey := new(string)
	if !r.Configuration.AwsAccessKey.IsUnknown() && !r.Configuration.AwsAccessKey.IsNull() {
		*awsAccessKey = r.Configuration.AwsAccessKey.ValueString()
	} else {
		awsAccessKey = nil
	}
	awsEnvironment := shared.SourceAmazonSellerPartnerAWSEnvironment(r.Configuration.AwsEnvironment.ValueString())
	awsSecretKey := new(string)
	if !r.Configuration.AwsSecretKey.IsUnknown() && !r.Configuration.AwsSecretKey.IsNull() {
		*awsSecretKey = r.Configuration.AwsSecretKey.ValueString()
	} else {
		awsSecretKey = nil
	}
	lwaAppID := r.Configuration.LwaAppID.ValueString()
	lwaClientSecret := r.Configuration.LwaClientSecret.ValueString()
	maxWaitSeconds := new(int64)
	if !r.Configuration.MaxWaitSeconds.IsUnknown() && !r.Configuration.MaxWaitSeconds.IsNull() {
		*maxWaitSeconds = r.Configuration.MaxWaitSeconds.ValueInt64()
	} else {
		maxWaitSeconds = nil
	}
	periodInDays := new(int64)
	if !r.Configuration.PeriodInDays.IsUnknown() && !r.Configuration.PeriodInDays.IsNull() {
		*periodInDays = r.Configuration.PeriodInDays.ValueInt64()
	} else {
		periodInDays = nil
	}
	refreshToken := r.Configuration.RefreshToken.ValueString()
	region := shared.SourceAmazonSellerPartnerAWSRegion(r.Configuration.Region.ValueString())
	replicationEndDate := new(string)
	if !r.Configuration.ReplicationEndDate.IsUnknown() && !r.Configuration.ReplicationEndDate.IsNull() {
		*replicationEndDate = r.Configuration.ReplicationEndDate.ValueString()
	} else {
		replicationEndDate = nil
	}
	replicationStartDate := r.Configuration.ReplicationStartDate.ValueString()
	reportOptions := new(string)
	if !r.Configuration.ReportOptions.IsUnknown() && !r.Configuration.ReportOptions.IsNull() {
		*reportOptions = r.Configuration.ReportOptions.ValueString()
	} else {
		reportOptions = nil
	}
	roleArn := new(string)
	if !r.Configuration.RoleArn.IsUnknown() && !r.Configuration.RoleArn.IsNull() {
		*roleArn = r.Configuration.RoleArn.ValueString()
	} else {
		roleArn = nil
	}
	sourceType := shared.SourceAmazonSellerPartnerAmazonSellerPartner(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceAmazonSellerPartner{
		AppID:                appID,
		AuthType:             authType,
		AwsAccessKey:         awsAccessKey,
		AwsEnvironment:       awsEnvironment,
		AwsSecretKey:         awsSecretKey,
		LwaAppID:             lwaAppID,
		LwaClientSecret:      lwaClientSecret,
		MaxWaitSeconds:       maxWaitSeconds,
		PeriodInDays:         periodInDays,
		RefreshToken:         refreshToken,
		Region:               region,
		ReplicationEndDate:   replicationEndDate,
		ReplicationStartDate: replicationStartDate,
		ReportOptions:        reportOptions,
		RoleArn:              roleArn,
		SourceType:           sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAmazonSellerPartnerCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAmazonSellerPartnerResourceModel) ToDeleteSDKType() *shared.SourceAmazonSellerPartnerCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAmazonSellerPartnerResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
