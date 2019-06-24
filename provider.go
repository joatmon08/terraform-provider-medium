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
	}
}
