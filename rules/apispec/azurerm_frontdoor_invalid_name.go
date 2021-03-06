// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermFrontdoorInvalidNameRule checks the pattern is valid
type AzurermFrontdoorInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermFrontdoorInvalidNameRule returns new rule with default attributes
func NewAzurermFrontdoorInvalidNameRule() *AzurermFrontdoorInvalidNameRule {
	return &AzurermFrontdoorInvalidNameRule{
		resourceType:  "azurerm_frontdoor",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9]+([-a-zA-Z0-9]?[a-zA-Z0-9])*$`),
	}
}

// Name returns the rule name
func (r *AzurermFrontdoorInvalidNameRule) Name() string {
	return "azurerm_frontdoor_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermFrontdoorInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermFrontdoorInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermFrontdoorInvalidNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermFrontdoorInvalidNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9]+([-a-zA-Z0-9]?[a-zA-Z0-9])*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
