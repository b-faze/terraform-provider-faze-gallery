package fazegallery

import (
	"context"
	"time"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	gc "github.com/b-faze/faze-gallery-client-go"
)

func resourceVisualisation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVisualisationCreate,
		ReadContext:   resourceVisualisationRead,
		UpdateContext: resourceVisualisationUpdate,
		DeleteContext: resourceVisualisationDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceVisualisationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*gc.APIClient).VisualisationsApi

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	vis := gc.Visualisation{
		Name: d.Get("name").(string),
	}

	visBody := gc.VisualisationsApiVisualisationsPostOpts{
		Body: optional.NewInterface(vis),
	}
	visResult, _, err := c.VisualisationsPost(context.TODO(), &visBody)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(visResult.Id)

	resourceVisualisationRead(ctx, d, m)

	return diags
}

func resourceVisualisationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*gc.APIClient).VisualisationsApi

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	vis, _, err := c.VisualisationsIdGet(context.TODO(), d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", vis.Name); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceVisualisationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*gc.APIClient).VisualisationsApi

	vis := gc.Visualisation{
		Name: d.Get("name").(string),
	}
	visBody := gc.VisualisationsApiVisualisationsIdPutOpts{
		Body: optional.NewInterface(vis),
	}
	_, _, err := c.VisualisationsIdPut(context.TODO(), d.Id(), &visBody)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceVisualisationRead(ctx, d, m)
}

func resourceVisualisationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*gc.APIClient).VisualisationsApi

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, _, err := c.VisualisationsIdDelete(context.TODO(), d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
