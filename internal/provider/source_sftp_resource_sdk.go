// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSftpResourceModel) ToCreateSDKType() *shared.SourceSftpCreateRequest {
	var credentials *shared.SourceSftpAuthenticationWildcard
	var sourceSftpAuthenticationWildcardPasswordAuthentication *shared.SourceSftpAuthenticationWildcardPasswordAuthentication
	if r.Configuration.Credentials.SourceSftpAuthenticationWildcardPasswordAuthentication != nil {
		authMethod := shared.SourceSftpAuthenticationWildcardPasswordAuthenticationAuthMethod(r.Configuration.Credentials.SourceSftpAuthenticationWildcardPasswordAuthentication.AuthMethod.ValueString())
		authUserPassword := r.Configuration.Credentials.SourceSftpAuthenticationWildcardPasswordAuthentication.AuthUserPassword.ValueString()
		sourceSftpAuthenticationWildcardPasswordAuthentication = &shared.SourceSftpAuthenticationWildcardPasswordAuthentication{
			AuthMethod:       authMethod,
			AuthUserPassword: authUserPassword,
		}
	}
	if sourceSftpAuthenticationWildcardPasswordAuthentication != nil {
		credentials = &shared.SourceSftpAuthenticationWildcard{
			SourceSftpAuthenticationWildcardPasswordAuthentication: sourceSftpAuthenticationWildcardPasswordAuthentication,
		}
	}
	var sourceSftpAuthenticationWildcardSSHKeyAuthentication *shared.SourceSftpAuthenticationWildcardSSHKeyAuthentication
	if r.Configuration.Credentials.SourceSftpAuthenticationWildcardSSHKeyAuthentication != nil {
		authMethod1 := shared.SourceSftpAuthenticationWildcardSSHKeyAuthenticationAuthMethod(r.Configuration.Credentials.SourceSftpAuthenticationWildcardSSHKeyAuthentication.AuthMethod.ValueString())
		authSSHKey := r.Configuration.Credentials.SourceSftpAuthenticationWildcardSSHKeyAuthentication.AuthSSHKey.ValueString()
		sourceSftpAuthenticationWildcardSSHKeyAuthentication = &shared.SourceSftpAuthenticationWildcardSSHKeyAuthentication{
			AuthMethod: authMethod1,
			AuthSSHKey: authSSHKey,
		}
	}
	if sourceSftpAuthenticationWildcardSSHKeyAuthentication != nil {
		credentials = &shared.SourceSftpAuthenticationWildcard{
			SourceSftpAuthenticationWildcardSSHKeyAuthentication: sourceSftpAuthenticationWildcardSSHKeyAuthentication,
		}
	}
	filePattern := new(string)
	if !r.Configuration.FilePattern.IsUnknown() && !r.Configuration.FilePattern.IsNull() {
		*filePattern = r.Configuration.FilePattern.ValueString()
	} else {
		filePattern = nil
	}
	fileTypes := new(string)
	if !r.Configuration.FileTypes.IsUnknown() && !r.Configuration.FileTypes.IsNull() {
		*fileTypes = r.Configuration.FileTypes.ValueString()
	} else {
		fileTypes = nil
	}
	folderPath := new(string)
	if !r.Configuration.FolderPath.IsUnknown() && !r.Configuration.FolderPath.IsNull() {
		*folderPath = r.Configuration.FolderPath.ValueString()
	} else {
		folderPath = nil
	}
	host := r.Configuration.Host.ValueString()
	port := r.Configuration.Port.ValueInt64()
	sourceType := shared.SourceSftpSftp(r.Configuration.SourceType.ValueString())
	user := r.Configuration.User.ValueString()
	configuration := shared.SourceSftp{
		Credentials: credentials,
		FilePattern: filePattern,
		FileTypes:   fileTypes,
		FolderPath:  folderPath,
		Host:        host,
		Port:        port,
		SourceType:  sourceType,
		User:        user,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSftpCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSftpResourceModel) ToDeleteSDKType() *shared.SourceSftpCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSftpResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
