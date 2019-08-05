package medium

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/joatmon08/terraform-provider-medium/readmedium"
)

func TestAccResourceImage_basic(t *testing.T) {
	var image readmedium.Image

	contentType := "image/png"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMediumImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMediumImageConfig_basic(contentType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMediumImageExists("medium_image.test", &image),
					testAccCheckMediumImageAttributes(&image, contentType),
					resource.TestCheckResourceAttr("medium_image.test", "content_type", contentType),
				),
			},
		},
	})
}

func testAccCheckMediumImageConfig_basic(contentType string) string {
	return fmt.Sprintf(`
	resource "medium_image" "test" {
		file_path    = "./resources/test.png"
		content_type = "%s"
	}
	`, contentType)
}

func testAccCheckMediumImageDestroy(s *terraform.State) error {
	return nil
}

func testAccCheckMediumImageExists(n string, image *readmedium.Image) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no image ID is set")
		}

		foundImage, err := readmedium.GetImage(rs.Primary.ID)

		if err != nil {
			return err
		}

		if foundImage.URL == rs.Primary.ID {
			return fmt.Errorf("image not found: %s", foundImage.URL)
		}

		if strings.Contains("http", foundImage.URL) {
			return fmt.Errorf("image ID is not a URL")
		}

		*image = *foundImage

		return nil
	}
}

func testAccCheckMediumImageAttributes(image *readmedium.Image, contentType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if image.ContentType != contentType {
			return fmt.Errorf("content type does not match: %s", image.ContentType)
		}

		return nil
	}
}
