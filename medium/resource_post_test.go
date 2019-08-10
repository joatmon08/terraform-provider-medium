package medium

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/joatmon08/terraform-provider-medium/readmedium"
)

func TestAccResourcePostMarkdown(t *testing.T) {
	var post readmedium.Story

	title := fmt.Sprintf("Unit Testing on %s", time.Now().Format("Jan 02, 2006 at 15:04:05"))
	content := "# Hello World!"
	contentFormat := "markdown"
	publishStatus := "draft"

	expectedPost := readmedium.Story{
		Payload: readmedium.Payload{
			Value: readmedium.Value{
				Title: title,
				Virtuals: readmedium.Virtuals{
					Tags: []string{},
				},
			},
		},
	}
	
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMediumPostDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMediumPostConfig_basic(title, content, contentFormat, publishStatus),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMediumPostExists("medium_post.test", &post),
					testAccCheckMediumPostAttributes(&post, &expectedPost),
					resource.TestCheckResourceAttr("medium_post.test", "title", title),
					resource.TestCheckResourceAttr("medium_post.test", "content", content),
					resource.TestCheckResourceAttr("medium_post.test", "content_format", contentFormat),
					resource.TestCheckResourceAttr("medium_post.test", "publish_status", publishStatus),
					resource.TestCheckResourceAttr("medium_post.test", "has_unpublished_edits", "true"),
				),
			},
		},
	})
}

func testAccCheckMediumPostConfig_basic(title string, content string, contentFormat string, publishStatus string) string {
	return fmt.Sprintf(`
	resource "medium_post" "test" {
		title          = "%s"
		content        = "%s"
		content_format = "%s"
		publish_status = "%s"
	}
	`, title, content, contentFormat, publishStatus)
}

func testAccCheckMediumPostDestroy(s *terraform.State) error {
	return nil
}

func testAccCheckMediumPostExists(n string, post *readmedium.Story) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no post ID is set")
		}

		var config Config
		if err := config.LoadAndValidate(); err != nil {
			return err
		}
		
		foundStory, err := config.StoryEndpoint.GetStory(config.User.ID, rs.Primary.ID)
		if err != nil {
			return err
		}

		if foundStory.Payload.Value.ID != rs.Primary.ID {
			return fmt.Errorf("post not found")
		}

		*post = *foundStory

		return nil
	}
}

func testAccCheckMediumPostAttributes(post *readmedium.Story, expectedPost *readmedium.Story) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if (len(post.Payload.Value.Virtuals.Tags) != len(expectedPost.Payload.Value.Virtuals.Tags)) {
			return fmt.Errorf("tags do not match: expected %s, got %s", expectedPost.Payload.Value.Virtuals.Tags, post.Payload.Value.Virtuals.Tags)
		}

		var config Config
		if err := config.LoadAndValidate(); err != nil {
			return fmt.Errorf("could not load configuration: %s", err)
		}

		story, err := config.StoryEndpoint.GetStory(config.User.ID, post.Payload.Value.ID)

		if (story.Payload.Value.MediumURL != post.Payload.Value.MediumURL) {
			return fmt.Errorf("urls do not match: expected %s, got %s", post.Payload.Value.MediumURL, story.Payload.Value.MediumURL)
		}

		return nil
	}
}