package medium

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcePost() *schema.Resource {
	return &schema.Resource{
		Create: resourcePostCreate,
		Read:   resourcePostRead,
		Update: resourcePostUpdate,
		Delete: resourcePostDelete,

		Schema: map[string]*schema.Schema{
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePostCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId("abcd")
	return resourcePostRead(d, m)
}

func resourcePostRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePostUpdate(d *schema.ResourceData, m interface{}) error {
	return resourcePostRead(d, m)
}

func resourcePostDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
