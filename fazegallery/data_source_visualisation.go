package fazegallery

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVisualisation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceImagesRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vis_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vis_album": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vis_pipeline_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vis_depth": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vis_variation": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vis_tree_size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vis_image_size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceVisualisationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	origin := "http://localhost:5000"
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/images", origin), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	images := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&images)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("images", images); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
