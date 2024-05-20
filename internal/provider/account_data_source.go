package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/netlifyapi"
)

var (
	_ datasource.DataSource              = &accountDataSource{}
	_ datasource.DataSourceWithConfigure = &accountDataSource{}
)

func NewAccountDataSource() datasource.DataSource {
	return &accountDataSource{}
}

type accountDataSource struct {
	data NetlifyProviderData
}

type accountDataSourceModel struct {
	ID   types.String `tfsdk:"id"`
	Slug types.String `tfsdk:"slug"`
	Name types.String `tfsdk:"name"`
}

func (d *accountDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(NetlifyProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected NetlifyProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.data = data
}

func (d *accountDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account"
}

func (d *accountDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.AtLeastOneOf(path.MatchRoot("slug")),
				},
			},
			"slug": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *accountDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config accountDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var account *netlifyapi.Account
	if !config.ID.IsUnknown() && !config.ID.IsNull() {
		var err error
		account, _, err = d.data.client.AccountsAPI.GetAccount(ctx, config.ID.ValueString()).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error reading Netlify account", fmt.Sprintf("Could not read Netlify account ID %q: %q",
				config.ID.ValueString(), err.Error()))
			return
		}
	} else {
		accounts, _, err := d.data.client.AccountsAPI.ListAccountsForUser(ctx).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error reading Netlify account", fmt.Sprintf("Could not list Netlify accounts: %q", err.Error()))
			return
		}
		slugString := config.Slug.ValueString()
		for _, a := range accounts {
			if a.Slug == slugString {
				acc := a
				account = &acc
				break
			}
		}
		if account == nil {
			resp.Diagnostics.AddError("Error reading Netlify account", fmt.Sprintf("Could not find Netlify account with slug %q", slugString))
			return
		}
	}

	config.ID = types.StringValue(account.Id)
	config.Slug = types.StringValue(account.Slug)
	config.Name = types.StringValue(account.Name)

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
