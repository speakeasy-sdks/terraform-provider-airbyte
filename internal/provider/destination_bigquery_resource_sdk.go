// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationBigqueryResourceModel) ToCreateSDKType() *shared.DestinationBigqueryCreateRequest {
	bigQueryClientBufferSizeMb := new(int64)
	if !r.Configuration.BigQueryClientBufferSizeMb.IsUnknown() && !r.Configuration.BigQueryClientBufferSizeMb.IsNull() {
		*bigQueryClientBufferSizeMb = r.Configuration.BigQueryClientBufferSizeMb.ValueInt64()
	} else {
		bigQueryClientBufferSizeMb = nil
	}
	credentialsJSON := new(string)
	if !r.Configuration.CredentialsJSON.IsUnknown() && !r.Configuration.CredentialsJSON.IsNull() {
		*credentialsJSON = r.Configuration.CredentialsJSON.ValueString()
	} else {
		credentialsJSON = nil
	}
	datasetID := r.Configuration.DatasetID.ValueString()
	datasetLocation := shared.DestinationBigqueryDatasetLocation(r.Configuration.DatasetLocation.ValueString())
	destinationType := shared.DestinationBigqueryBigquery(r.Configuration.DestinationType.ValueString())
	var loadingMethod *shared.DestinationBigqueryLoadingMethod
	var destinationBigqueryLoadingMethodStandardInserts *shared.DestinationBigqueryLoadingMethodStandardInserts
	if r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodStandardInserts != nil {
		method := shared.DestinationBigqueryLoadingMethodStandardInsertsMethod(r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodStandardInserts.Method.ValueString())
		destinationBigqueryLoadingMethodStandardInserts = &shared.DestinationBigqueryLoadingMethodStandardInserts{
			Method: method,
		}
	}
	if destinationBigqueryLoadingMethodStandardInserts != nil {
		loadingMethod = &shared.DestinationBigqueryLoadingMethod{
			DestinationBigqueryLoadingMethodStandardInserts: destinationBigqueryLoadingMethodStandardInserts,
		}
	}
	var destinationBigqueryLoadingMethodGCSStaging *shared.DestinationBigqueryLoadingMethodGCSStaging
	if r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodGCSStaging != nil {
		var credential shared.DestinationBigqueryLoadingMethodGCSStagingCredential
		var destinationBigqueryLoadingMethodGCSStagingCredentialHMACKey *shared.DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey
		if r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodGCSStaging.Credential.DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey != nil {
			credentialType := shared.DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKeyCredentialType(r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodGCSStaging.Credential.DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey.CredentialType.ValueString())
			hmacKeyAccessID := r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodGCSStaging.Credential.DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey.HmacKeyAccessID.ValueString()
			hmacKeySecret := r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodGCSStaging.Credential.DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey.HmacKeySecret.ValueString()
			destinationBigqueryLoadingMethodGCSStagingCredentialHMACKey = &shared.DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey{
				CredentialType:  credentialType,
				HmacKeyAccessID: hmacKeyAccessID,
				HmacKeySecret:   hmacKeySecret,
			}
		}
		if destinationBigqueryLoadingMethodGCSStagingCredentialHMACKey != nil {
			credential = shared.DestinationBigqueryLoadingMethodGCSStagingCredential{
				DestinationBigqueryLoadingMethodGCSStagingCredentialHMACKey: destinationBigqueryLoadingMethodGCSStagingCredentialHMACKey,
			}
		}
		fileBufferCount := new(int64)
		if !r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodGCSStaging.FileBufferCount.IsUnknown() && !r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodGCSStaging.FileBufferCount.IsNull() {
			*fileBufferCount = r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodGCSStaging.FileBufferCount.ValueInt64()
		} else {
			fileBufferCount = nil
		}
		gcsBucketName := r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodGCSStaging.GcsBucketName.ValueString()
		gcsBucketPath := r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodGCSStaging.GcsBucketPath.ValueString()
		keepFilesInGcsBucket := new(shared.DestinationBigqueryLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing)
		if !r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodGCSStaging.KeepFilesInGcsBucket.IsUnknown() && !r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodGCSStaging.KeepFilesInGcsBucket.IsNull() {
			*keepFilesInGcsBucket = shared.DestinationBigqueryLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing(r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodGCSStaging.KeepFilesInGcsBucket.ValueString())
		} else {
			keepFilesInGcsBucket = nil
		}
		method1 := shared.DestinationBigqueryLoadingMethodGCSStagingMethod(r.Configuration.LoadingMethod.DestinationBigqueryLoadingMethodGCSStaging.Method.ValueString())
		destinationBigqueryLoadingMethodGCSStaging = &shared.DestinationBigqueryLoadingMethodGCSStaging{
			Credential:           credential,
			FileBufferCount:      fileBufferCount,
			GcsBucketName:        gcsBucketName,
			GcsBucketPath:        gcsBucketPath,
			KeepFilesInGcsBucket: keepFilesInGcsBucket,
			Method:               method1,
		}
	}
	if destinationBigqueryLoadingMethodGCSStaging != nil {
		loadingMethod = &shared.DestinationBigqueryLoadingMethod{
			DestinationBigqueryLoadingMethodGCSStaging: destinationBigqueryLoadingMethodGCSStaging,
		}
	}
	projectID := r.Configuration.ProjectID.ValueString()
	transformationPriority := new(shared.DestinationBigqueryTransformationQueryRunType)
	if !r.Configuration.TransformationPriority.IsUnknown() && !r.Configuration.TransformationPriority.IsNull() {
		*transformationPriority = shared.DestinationBigqueryTransformationQueryRunType(r.Configuration.TransformationPriority.ValueString())
	} else {
		transformationPriority = nil
	}
	configuration := shared.DestinationBigquery{
		BigQueryClientBufferSizeMb: bigQueryClientBufferSizeMb,
		CredentialsJSON:            credentialsJSON,
		DatasetID:                  datasetID,
		DatasetLocation:            datasetLocation,
		DestinationType:            destinationType,
		LoadingMethod:              loadingMethod,
		ProjectID:                  projectID,
		TransformationPriority:     transformationPriority,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationBigqueryCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationBigqueryResourceModel) ToDeleteSDKType() *shared.DestinationBigqueryCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationBigqueryResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
