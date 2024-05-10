---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netlify_dns_zone Resource - netlify"
subcategory: ""
description: |-
  
---

# netlify_dns_zone (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_slug` (String)
- `name` (String)

### Read-Only

- `account_id` (String)
- `dns_servers` (List of String)
- `domain` (Block, Read-only) (see [below for nested schema](#nestedblock--domain))
- `id` (String) The ID of this resource.
- `last_updated` (String)

<a id="nestedblock--domain"></a>
### Nested Schema for `domain`

Read-Only:

- `auto_renew` (Boolean)
- `auto_renew_at` (String)
- `expires_at` (String)
- `id` (String)
- `name` (String)
- `registered_at` (String)
- `renewal_price` (String)