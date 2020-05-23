<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_virtual_machine_data_disk_attachment_invalid_create_option

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

Allowed values are:
- FromImage
- Empty
- Attach

## Example

```hcl
resource "azurerm_virtual_machine_data_disk_attachment" "foo" {
  create_option = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "..." is an invalid value as create_option (azurerm_virtual_machine_data_disk_attachment_invalid_create_option)

  on template.tf line 2:
  2:   create_option = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_virtual_machine_data_disk_attachment_invalid_create_option.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/compute.json