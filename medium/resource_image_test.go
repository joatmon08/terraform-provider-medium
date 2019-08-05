package medium

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/joatmon08/terraform-provider-medium/testimage"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/joatmon08/terraform-provider-medium/readmedium"
)

func TestAccResourceImage(t *testing.T) {
	var image readmedium.Image
	testImage, err := ioutil.ReadFile("resources/test.png")
	if err != nil {
		t.Fatalf("could not read test image: %s", err)
	}

	contentType := "image/png"
	md5 := testimage.GetBase64MD5(testImage)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMediumImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMediumImageConfig(contentType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMediumImageExists("medium_image.test", &image),
					testAccCheckMediumImageAttributes(&image, md5, contentType),
					resource.TestCheckResourceAttr("medium_image.test", "content_type", contentType),
					resource.TestCheckResourceAttr("medium_image.test", "md5", md5),
				),
			},
		},
	})
}

func testAccCheckMediumImageConfig(contentType string) string {
	return fmt.Sprintf(`
	resource "medium_image" "test" {
		file_path    = "./resources/test.png"
		content_type = "%s"
	}
	`, contentType)
}

func testAccCheckMediumImageAttributes(image *readmedium.Image, md5 string, contentType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if image.MD5 != md5 {
			return fmt.Errorf("image content does not match: %s", image.MD5)
		}

		if image.ContentType != contentType {
			return fmt.Errorf("image content type does not match: %s", image.ContentType)
		}

		return nil
	}
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

		if foundImage.URL != rs.Primary.ID {
			return fmt.Errorf("image not found: %s", foundImage.URL)
		}

		if strings.Contains("http", foundImage.URL) {
			return fmt.Errorf("image ID is not a URL")
		}

		*image = *foundImage

		return nil
	}
}