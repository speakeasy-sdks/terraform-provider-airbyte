// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

	speakeasy_stringplanmodifier "airbyte/internal/planmodifiers/stringplanmodifier"
	"airbyte/internal/sdk/pkg/models/operations"
	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DestinationRedisResource{}
var _ resource.ResourceWithImportState = &DestinationRedisResource{}

func NewDestinationRedisResource() resource.Resource {
	return &DestinationRedisResource{}
}

// DestinationRedisResource defines the resource implementation.
type DestinationRedisResource struct {
	client *sdk.SDK
}

// DestinationRedisResourceModel describes the resource data model.
type DestinationRedisResourceModel struct {
	Configuration   DestinationRedis `tfsdk:"configuration"`
	DestinationID   types.String     `tfsdk:"destination_id"`
	DestinationType types.String     `tfsdk:"destination_type"`
	Name            types.String     `tfsdk:"name"`
	WorkspaceID     types.String     `tfsdk:"workspace_id"`
}

func (r *DestinationRedisResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_redis"
}

func (r *DestinationRedisResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationRedis Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"cache_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"hash",
							),
						},
						Description: `Redis cache type to store data in.`,
					},
					"destination_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"redis",
							),
						},
					},
					"host": schema.StringAttribute{
						Required: true,
					},
					"password": schema.StringAttribute{
						Optional: true,
					},
					"port": schema.Int64Attribute{
						Required: true,
					},
					"ssl": schema.BoolAttribute{
						Optional: true,
					},
					"ssl_mode": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"destination_redis_ssl_modes_disable": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"disable",
											),
										},
									},
								},
								Description: `Disable SSL.`,
							},
							"destination_redis_ssl_modes_verify_full": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										Required: true,
									},
									"client_certificate": schema.StringAttribute{
										Required: true,
									},
									"client_key": schema.StringAttribute{
										Required: true,
									},
									"client_key_password": schema.StringAttribute{
										Optional: true,
									},
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"verify-full",
											),
										},
									},
								},
								Description: `Verify-full SSL mode.`,
							},
							"destination_redis_update_ssl_modes_disable": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"disable",
											),
										},
									},
								},
								Description: `Disable SSL.`,
							},
							"destination_redis_update_ssl_modes_verify_full": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										Required: true,
									},
									"client_certificate": schema.StringAttribute{
										Required: true,
									},
									"client_key": schema.StringAttribute{
										Required: true,
									},
									"client_key_password": schema.StringAttribute{
										Optional: true,
									},
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"verify-full",
											),
										},
									},
								},
								Description: `Verify-full SSL mode.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"tunnel_method": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"destination_redis_ssh_tunnel_method_no_tunnel": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"NO_TUNNEL",
											),
										},
										Description: `No ssh tunnel needed to connect to database`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"destination_redis_ssh_tunnel_method_password_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_host": schema.StringAttribute{
										Required: true,
									},
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_PASSWORD_AUTH",
											),
										},
										Description: `Connect through a jump server tunnel host using username and password authentication`,
									},
									"tunnel_port": schema.Int64Attribute{
										Required: true,
									},
									"tunnel_user": schema.StringAttribute{
										Required: true,
									},
									"tunnel_user_password": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"destination_redis_ssh_tunnel_method_ssh_key_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ssh_key": schema.StringAttribute{
										Required: true,
									},
									"tunnel_host": schema.StringAttribute{
										Required: true,
									},
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_KEY_AUTH",
											),
										},
										Description: `Connect through a jump server tunnel host using username and ssh key`,
									},
									"tunnel_port": schema.Int64Attribute{
										Required: true,
									},
									"tunnel_user": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"destination_redis_update_ssh_tunnel_method_no_tunnel": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"NO_TUNNEL",
											),
										},
										Description: `No ssh tunnel needed to connect to database`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"destination_redis_update_ssh_tunnel_method_password_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_host": schema.StringAttribute{
										Required: true,
									},
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_PASSWORD_AUTH",
											),
										},
										Description: `Connect through a jump server tunnel host using username and password authentication`,
									},
									"tunnel_port": schema.Int64Attribute{
										Required: true,
									},
									"tunnel_user": schema.StringAttribute{
										Required: true,
									},
									"tunnel_user_password": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"destination_redis_update_ssh_tunnel_method_ssh_key_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ssh_key": schema.StringAttribute{
										Required: true,
									},
									"tunnel_host": schema.StringAttribute{
										Required: true,
									},
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_KEY_AUTH",
											),
										},
										Description: `Connect through a jump server tunnel host using username and ssh key`,
									},
									"tunnel_port": schema.Int64Attribute{
										Required: true,
									},
									"tunnel_user": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"username": schema.StringAttribute{
						Required: true,
					},
				},
			},
			"destination_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.TrustRead(),
				},
			},
			"destination_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.TrustRead(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"workspace_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *DestinationRedisResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *DestinationRedisResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationRedisResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToCreateSDKType()
	res, err := r.client.Destinations.CreateDestinationRedis(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationRedisResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationRedisResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	destinationID := data.DestinationID.ValueString()
	request := operations.GetDestinationRedisRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationRedis(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationRedisResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationRedisResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	destinationRedisPutRequest := data.ToUpdateSDKType()
	destinationID := data.DestinationID.ValueString()
	request := operations.PutDestinationRedisRequest{
		DestinationRedisPutRequest: destinationRedisPutRequest,
		DestinationID:              destinationID,
	}
	res, err := r.client.Destinations.PutDestinationRedis(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	destinationId1 := data.DestinationID.ValueString()
	getRequest := operations.GetDestinationRedisRequest{
		DestinationID: destinationId1,
	}
	getResponse, err := r.client.Destinations.GetDestinationRedis(ctx, getRequest)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if getResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", getResponse))
		return
	}
	if getResponse.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", getResponse.StatusCode), debugResponse(getResponse.RawResponse))
		return
	}
	if getResponse.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(getResponse.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationRedisResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationRedisResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	destinationID := data.DestinationID.ValueString()
	request := operations.DeleteDestinationRedisRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.DeleteDestinationRedis(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *DestinationRedisResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("destination_id"), req, resp)
}
