resource "medium_post" "my-post" {
    title = "Writing a Terraform Medium Provider"
    content = file("./content.md")
    publish_status = "draft"
    tags = ["terraform", "medium", "automation"]
}