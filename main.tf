resource "medium_post" "my-post" {
    title = "Writing a Terraform Medium Provider"
    content = file("./content.md")
    publish_status = "public"
    tags = ["terraform", "medium", "automation"]
}