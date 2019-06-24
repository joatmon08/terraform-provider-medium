package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/joatmon08/terraform-provider-medium/medium"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"medium_post": medium.ResourcePost(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var config medium.Config
	if err := config.LoadAndValidate(); err != nil {
		return nil, err
	}
	return &config, nil
}