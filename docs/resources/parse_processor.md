---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mezmo_parse_processor Resource - terraform-provider-mezmo"
subcategory: ""
description: |-
  Parse a specified field using the chosen parser
---

# mezmo_parse_processor (Resource)

Parse a specified field using the chosen parser



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `field` (String) The JSON field whose value should be parsed.
- `parser` (String) The kind of parser to use against the input value from "field".
- `pipeline_id` (String) The uuid of the pipeline

### Optional

- `apache_log_options` (Attributes) Options for apache log parser (see [below for nested schema](#nestedatt--apache_log_options))
- `cef_log_options` (Attributes) Options for common event format (CEF) log parser (see [below for nested schema](#nestedatt--cef_log_options))
- `csv_row_options` (Attributes) Options for CSV row parser (see [below for nested schema](#nestedatt--csv_row_options))
- `description` (String) A user-defined value describing the processor
- `grok_parser_options` (Attributes) Options for grok parser (see [below for nested schema](#nestedatt--grok_parser_options))
- `inputs` (List of String) The ids of the input components
- `key_value_log_options` (Attributes) Options for key value log parser (see [below for nested schema](#nestedatt--key_value_log_options))
- `nginx_log_options` (Attributes) Options for nginx log parser (see [below for nested schema](#nestedatt--nginx_log_options))
- `regex_parser_options` (Attributes) Options for regex parser (see [below for nested schema](#nestedatt--regex_parser_options))
- `target_field` (String) The field into which the parsed value should be inserted. Leave blank to insert the parsed data into the original field.
- `timestamp_parser_options` (Attributes) Options for timestamp log parser (see [below for nested schema](#nestedatt--timestamp_parser_options))
- `title` (String) A user-defined title for the processor

### Read-Only

- `generation_id` (Number) An internal field used for component versioning
- `id` (String) The uuid of the processor

<a id="nestedatt--apache_log_options"></a>
### Nested Schema for `apache_log_options`

Required:

- `format` (String) The log format.

Optional:

- `custom_timestamp_format` (String) Custom timestamp format, according to strftime, for log entries.
- `timestamp_format` (String) The timestamp format of log entries.


<a id="nestedatt--cef_log_options"></a>
### Nested Schema for `cef_log_options`

Optional:

- `translate_custom_fields` (Boolean) Translate custom fields in log.


<a id="nestedatt--csv_row_options"></a>
### Nested Schema for `csv_row_options`

Optional:

- `field_names` (List of String) The name of the CSV fields that take the value in the same order they appear in data


<a id="nestedatt--grok_parser_options"></a>
### Nested Schema for `grok_parser_options`

Required:

- `pattern` (String) The grok pattern. Must be composed of community expressions.


<a id="nestedatt--key_value_log_options"></a>
### Nested Schema for `key_value_log_options`

Optional:

- `field_delimiter` (String) One or more characters that separate each key/value pair.
- `key_value_delimiter` (String) One or more characters that separate each key and value.


<a id="nestedatt--nginx_log_options"></a>
### Nested Schema for `nginx_log_options`

Required:

- `format` (String) The log format.

Optional:

- `custom_timestamp_format` (String) Custom timestamp format, according to strftime, for log entries.
- `timestamp_format` (String) The timestamp format of log entries.


<a id="nestedatt--regex_parser_options"></a>
### Nested Schema for `regex_parser_options`

Required:

- `pattern` (String) The regex pattern.


<a id="nestedatt--timestamp_parser_options"></a>
### Nested Schema for `timestamp_parser_options`

Required:

- `format` (String) The timestamp format.

Optional:

- `custom_format` (String) Custom timestamp format, according to strftime, for log entries.


