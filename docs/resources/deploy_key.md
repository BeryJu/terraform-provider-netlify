---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netlify_deploy_key Resource - netlify"
subcategory: ""
description: |-
  Deploy key for Git repositories. Avoid creating this resource directly if possible. Read more https://docs.netlify.com/git/repo-permissions-linking/#deploy-keys
---

# netlify_deploy_key (Resource)

Deploy key for Git repositories. Avoid creating this resource directly if possible. [Read more](https://docs.netlify.com/git/repo-permissions-linking/#deploy-keys)

## Example Usage

```terraform
resource "netlify_deploy_key" "common" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `last_updated` (String)
- `public_key` (String)

## Import

Import is supported using the following syntax:

```shell
# Import a deploy key by its ID
terraform import netlify_deploy_key.common 6600abcdef1234567890abcd
```