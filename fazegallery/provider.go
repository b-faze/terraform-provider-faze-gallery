package fazegallery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	gc "github.com/b-faze/faze-gallery-client-go"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"fazegallery_visualisation": resourceVisualisation(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"fazegallery_images": dataSourceImages(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	clientConfig := gc.NewConfiguration()
	clientConfig.BasePath = "http://localhost:19090"
	c := gc.NewAPIClient(clientConfig)

	return c, diags
}
