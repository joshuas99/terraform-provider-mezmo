---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mezmo_http_source Resource - terraform-provider-mezmo"
subcategory: ""
description: |-
  Represents an HTTP source.
---

# mezmo_http_source (Resource)

Represents an HTTP source.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `pipeline_id` (String) The uuid of the pipeline

### Optional

- `capture_metadata` (Boolean) Enable the inclusion of all http headers and query string parameters that were sent from the source
- `decoding` (String) The decoding method for converting frames into data events.
- `description` (String) A user-defined value describing the source component
- `gateway_route_id` (String) The uuid of a pre-existing source to be used as the input for this component. This can only be provided on resource creation (not update).
- `title` (String) A user-defined title for the source component

### Read-Only

- `generation_id` (Number) An internal field used for component versioning
- `id` (String) The uuid of the source component


