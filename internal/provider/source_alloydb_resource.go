// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

	"airbyte/internal/sdk/pkg/models/operations"
	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceAlloydbResource{}
var _ resource.ResourceWithImportState = &SourceAlloydbResource{}

func NewSourceAlloydbResource() resource.Resource {
	return &SourceAlloydbResource{}
}

// SourceAlloydbResource defines the resource implementation.
type SourceAlloydbResource struct {
	client *sdk.SDK
}

// SourceAlloydbResourceModel describes the resource data model.
type SourceAlloydbResourceModel struct {
	Configuration SourceAlloydb `tfsdk:"configuration"`
	Name          types.String  `tfsdk:"name"`
	SecretID      types.String  `tfsdk:"secret_id"`
	SourceID      types.String  `tfsdk:"source_id"`
	SourceType    types.String  `tfsdk:"source_type"`
	WorkspaceID   types.String  `tfsdk:"workspace_id"`
}

func (r *SourceAlloydbResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_alloydb"
}

func (r *SourceAlloydbResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceAlloydb Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
				Required: true,
				Attributes: map[string]schema.Attribute{
					"database": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required: true,
					},
					"host": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required: true,
					},
					"jdbc_url_params": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Optional: true,
					},
					"password": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Optional: true,
					},
					"port": schema.Int64Attribute{
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
						Required: true,
					},
					"replication_method": schema.SingleNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplace(),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_alloydb_replication_method_standard": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Standard",
											),
										},
									},
								},
								Description: `Standard replication requires no setup on the DB side but will not be able to represent deletions incrementally.`,
							},
							"source_alloydb_replication_method_logical_replication_cdc_": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"initial_waiting_seconds": schema.Int64Attribute{
										Computed: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"lsn_commit_behaviour": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"While reading Data",
												"After loading Data in the destination",
											),
										},
										Description: `Determines when Airbtye should flush the LSN of processed WAL logs in the source database. ` + "`" + `After loading Data in the destination` + "`" + ` is default. If ` + "`" + `While reading Data` + "`" + ` is selected, in case of a downstream failure (while loading data into the destination), next sync would result in a full sync.`,
									},
									"method": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"CDC",
											),
										},
									},
									"plugin": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"pgoutput",
											),
										},
										Description: `A logical decoding plugin installed on the PostgreSQL server.`,
									},
									"publication": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
									},
									"replication_slot": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
									},
									"additional_properties": schema.StringAttribute{
										Computed: true,
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Logical replication uses the Postgres write-ahead log (WAL) to detect inserts, updates, and deletes. This needs to be configured on the source database itself. Only available on Postgres 10 and above. Read the <a href="https://docs.airbyte.com/integrations/sources/postgres">docs</a>.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"schemas": schema.ListAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.RequiresReplace(),
						},
						Optional:    true,
						ElementType: types.StringType,
					},
					"source_type": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"alloydb",
							),
						},
					},
					"ssl_mode": schema.SingleNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplace(),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_alloydb_ssl_modes_disable": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"disable",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Computed: true,
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Disables encryption of communication between Airbyte and source database.`,
							},
							"source_alloydb_ssl_modes_allow": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"allow",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Computed: true,
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Enables encryption only when required by the source database.`,
							},
							"source_alloydb_ssl_modes_prefer": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"prefer",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Computed: true,
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Allows unencrypted connection only if the source database does not support encryption.`,
							},
							"source_alloydb_ssl_modes_require": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"require",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Computed: true,
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Always require encryption. If the source database server does not support encryption, connection will fail.`,
							},
							"source_alloydb_ssl_modes_verify_ca": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
									},
									"client_certificate": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"client_key": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"client_key_password": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"mode": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"verify-ca",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Computed: true,
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Always require encryption and verifies that the source database server has a valid SSL certificate.`,
							},
							"source_alloydb_ssl_modes_verify_full": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
									},
									"client_certificate": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"client_key": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"client_key_password": schema.StringAttribute{
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Optional: true,
									},
									"mode": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"verify-full",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Computed: true,
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `This is the most secure mode. Always require encryption and verifies the identity of the source database server.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"tunnel_method": schema.SingleNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplace(),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_alloydb_ssh_tunnel_method_no_tunnel": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_method": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
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
							"source_alloydb_ssh_tunnel_method_ssh_key_authentication": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ssh_key": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
									},
									"tunnel_host": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
									},
									"tunnel_method": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_KEY_AUTH",
											),
										},
										Description: `Connect through a jump server tunnel host using username and ssh key`,
									},
									"tunnel_port": schema.Int64Attribute{
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.RequiresReplace(),
										},
										Required: true,
									},
									"tunnel_user": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_alloydb_ssh_tunnel_method_password_authentication": schema.SingleNestedAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplace(),
								},
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_host": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
									},
									"tunnel_method": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_PASSWORD_AUTH",
											),
										},
										Description: `Connect through a jump server tunnel host using username and password authentication`,
									},
									"tunnel_port": schema.Int64Attribute{
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.RequiresReplace(),
										},
										Required: true,
									},
									"tunnel_user": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										Required: true,
									},
									"tunnel_user_password": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Required: true,
					},
				},
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
			"secret_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional: true,
			},
			"source_id": schema.StringAttribute{
				Computed: true,
			},
			"source_type": schema.StringAttribute{
				Computed: true,
			},
			"workspace_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
		},
	}
}

func (r *SourceAlloydbResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceAlloydbResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceAlloydbResourceModel
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
	res, err := r.client.Sources.CreateSourceAlloydb(ctx, request)
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
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceAlloydbResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceAlloydbResourceModel
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

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceAlloydbResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceAlloydbResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceAlloydbResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceAlloydbResourceModel
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

	sourceID := data.SourceID.ValueString()
	request := operations.DeleteSourceAlloydbRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceAlloydb(ctx, request)
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

func (r *SourceAlloydbResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
