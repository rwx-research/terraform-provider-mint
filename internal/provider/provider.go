package provider

import (
	"context"
	"os"

	"github.com/rwx-research/terraform-provider-mint/internal/api"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure MintProvider satisfies various the provider interface.
var _ provider.Provider = &MintProvider{}

type MintProvider struct {
	version string
}

// MintProviderModel describes the provider data model.
type MintProviderModel struct {
	Host        types.String `tfsdk:"host"`
	AccessToken types.String `tfsdk:"access_token"`
}

func (p *MintProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "mint"
	resp.Version = p.version
}

func (p *MintProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional: true,
			},
			"access_token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *MintProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config MintProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unkown Mint Host",
			"The provider cannot create the Mint API client as there is an unknown configuration value for the Mint host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the MINT_HOST environment variable.",
		)
	}
	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("access_token"),
			"Unkown Mint Access Token",
			"The provider cannot create the Mint API client as there is an unknown configuration value for the Mint access token. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the RWX_ACCESS_TOKEN environment variable.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	host := os.Getenv("MINT_HOST")
	accessToken := os.Getenv("RWX_ACCESS_TOKEN")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}
	if !config.AccessToken.IsNull() {
		accessToken = config.AccessToken.ValueString()
	}

	if host == "" {
		host = "cloud.rwx.com"
	}
	if accessToken == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("access_token"),
			"Missing Mint Access Token",
			"The provider cannot create the Mint API client as there is a missing or empty value for the Mint access token. "+
				"Set the access token value in the configuration or use the RWX_ACCESS_TOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	client, err := api.NewClient(api.Config{Host: host, AccessToken: accessToken, Version: p.version})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create Mint API client",
			"An unexpected error occurred when creating the Mint API client. "+
				"If the error is not clear, please contact us at support@rwx.com.\n\n"+
				"Original Error: "+err.Error(),
		)
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *MintProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewSecretResource,
		NewVariableResource,
	}
}

func (p *MintProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &MintProvider{
			version: version,
		}
	}
}
