# Writing a Terraform Medium Provider

I keep my experimental, personal blogging on Medium as future references for myself.
My typical workflow for blogs involves:

- Creating a new Medium story
- Keeping the tab open as I write
- Copy code and reformat repeatedly
- Accidentally close the tab
- Re-save on an airplane because the wireless drops
- And repeat.

With all of my struggles, I decided to write my own Terraform provider for
Medium. I figured this was a good way to "dogfood" the Terraform provider
learning process and finally fix my workflow struggle with Medium. Ideally, I
wanted my new workflow to be:

- Write the blog in Markdown.
- Add it to my Terraform configuration.
- Run `terraform plan`.
- Run `terraform apply`.
- Repeat until final draft.
- Change a configuration to "publish".
- Run `terraform apply`.

The new workflow lets me write my blogs in Markdown, offline and with version
control. Even better, if I want someone to review or collaborate, I add them as
a collaborator to the Markdown on Github and they can submit a pull request for edits.

# Exploring the Medium API

The Medium API can be a little tricky to use. It only allows write-only
operations, so yes to post and image creation but no to deletion and in-place
updates. I had to request developer access by submitting a support ticket.

Problematically, I wanted blog metadata with update timestamps, publication
status, and more. Given that the API does not have read operations, I had to
seek a public endpoint that returned a JSON with said metadata. After resorting
to my favorite search engine and a few experiments, I found that if I added the
`format=json` query parameter to a call to the Medium public endpoint, I
received my blog and draft metadata.

# Caveats

When you delete, it doesn't delete the post.