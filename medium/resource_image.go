package medium

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

var (
	imageContentTypes = []string{
		"image/jpeg",
		"image/png",
		"image/gif",
		"image/tiff",
	}
)

func ResourceImage() *schema.Resource {
	return &schema.Resource{
		Create: ResourceImageCreate,
		Read:   ResourceImageRead,
		Update: ResourceImageUpdate,
		Delete: ResourceImageDelete,

		Schema: map[string]*schema.Schema{
			"file_path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"content_type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(imageContentTypes, false),
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceImageCreate(d *schema.ResourceData, m interface{}) error {
	var builder PostBuilder
	config := m.(*Config)
	filePath := d.Get("file_path").(string)
	contentType := d.Get("content_type").(string)

	builder.BuildImage(filePath, contentType)
	image, err := config.Client.UploadImage(*builder.ImageOptions)
	if err != nil {
		return err
	}
	d.SetId(image.MD5)
	d.Set("url", image.URL)
	return ResourceImageRead(d, m)
}

func ResourceImageRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func ResourceImageUpdate(d *schema.ResourceData, m interface{}) error {
	return ResourceImageCreate(d, m)
}

func ResourceImageDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
