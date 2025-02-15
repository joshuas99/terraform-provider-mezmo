---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mezmo_honeycomb_logs_destination Resource - terraform-provider-mezmo"
subcategory: ""
description: |-
  Send log data to Honeycomb
---

# mezmo_honeycomb_logs_destination (Resource)

Send log data to Honeycomb



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `api_key` (String, Sensitive) Honeycomb API key
- `dataset` (String) The name of the targeted dataset
- `pipeline_id` (String) The uuid of the pipeline

### Optional

- `ack_enabled` (Boolean) Acknowledge data from the source when it reaches the destination
- `description` (String) A user-defined value describing the destination
- `inputs` (List of String) The ids of the input components
- `title` (String) A user-defined title for the destination

### Read-Only

- `generation_id` (Number) An internal field used for component versioning
- `id` (String) The uuid of the destination


