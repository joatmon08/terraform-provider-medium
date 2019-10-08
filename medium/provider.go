package medium

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"medium_post":  ResourcePost(),
			"medium_image": ResourceImage(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var config Config
	if err := config.LoadAndValidate(); err != nil {
		return nil, err
	}
	return &config, nil
}
