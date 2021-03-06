// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermBotChannelsRegistrationInvalidSkuRule checks the pattern is valid
type AzurermBotChannelsRegistrationInvalidSkuRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermBotChannelsRegistrationInvalidSkuRule returns new rule with default attributes
func NewAzurermBotChannelsRegistrationInvalidSkuRule() *AzurermBotChannelsRegistrationInvalidSkuRule {
	return &AzurermBotChannelsRegistrationInvalidSkuRule{
		resourceType:  "azurerm_bot_channels_registration",
		attributeName: "sku",
		enum: []string{
			"F0",
			"S1",
		},
	}
}

// Name returns the rule name
func (r *AzurermBotChannelsRegistrationInvalidSkuRule) Name() string {
	return "azurerm_bot_channels_registration_invalid_sku"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermBotChannelsRegistrationInvalidSkuRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermBotChannelsRegistrationInvalidSkuRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermBotChannelsRegistrationInvalidSkuRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermBotChannelsRegistrationInvalidSkuRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as sku`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
