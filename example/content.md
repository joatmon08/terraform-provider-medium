# Example for Terraform Provider for Medium

This is an example of how to use this Terraform provider for Medium. Note that
due to Medium's API constraints, this only supports:

- Creation of posts
- Uploading of images
- Immutable creation of posts

It does not support deletion of posts (that has to be done manually) or in-place
updates.

# Using Markdown to Medium

Below is how you setup an image.

<figure tabindex="0" contenteditable="false" class="graf graf--figure graf-after--p is-mediaFocused is-selected">
<div class="aspectRatioPlaceholder">
<img class="graf-image" alt="Image of my Draft box" src="${image_url}">
<div class="crosshair u-ignoreBlock">
</div></div>
<figcaption class="imageCaption" contenteditable="true" data-default-value="Type caption for image (optional)">this is my caption</figcaption>
</figure>

For the two types of quotes, you'll need to use HTML tags.

<blockquote class="graf graf--pullquote graf-after--blockquote graf--trailing is-selected">This is one type of quote.</blockquote>

<blockquote class="graf graf--blockquote graf-after--p">Here is another quote with the vertical line.</blockquote>