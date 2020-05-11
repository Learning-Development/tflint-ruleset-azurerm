// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermIothubRouteInvalidSourceRule checks the pattern is valid
type AzurermIothubRouteInvalidSourceRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermIothubRouteInvalidSourceRule returns new rule with default attributes
func NewAzurermIothubRouteInvalidSourceRule() *AzurermIothubRouteInvalidSourceRule {
	return &AzurermIothubRouteInvalidSourceRule{
		resourceType:  "azurerm_iothub_route",
		attributeName: "source",
		enum: []string{
			"Invalid",
			"DeviceMessages",
			"TwinChangeEvents",
			"DeviceLifecycleEvents",
			"DeviceJobLifecycleEvents",
		},
	}
}

// Name returns the rule name
func (r *AzurermIothubRouteInvalidSourceRule) Name() string {
	return "azurerm_iothub_route_invalid_source"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermIothubRouteInvalidSourceRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermIothubRouteInvalidSourceRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermIothubRouteInvalidSourceRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermIothubRouteInvalidSourceRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as source`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}