resource "medium_image" "draft" {
  file_path    = "./images/draft.png"
  content_type = "image/png"
}

resource "medium_post" "my-post" {
  title          = "Writing a Terraform Medium Provider"
  content        = templatefile("./content.md", { image_url = medium_image.draft.url })
  content_format = "markdown"
  publish_status = "draft"
}