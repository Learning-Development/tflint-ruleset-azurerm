// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermCosmosdbMongoDatabaseInvalidAccountNameRule checks the pattern is valid
type AzurermCosmosdbMongoDatabaseInvalidAccountNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermCosmosdbMongoDatabaseInvalidAccountNameRule returns new rule with default attributes
func NewAzurermCosmosdbMongoDatabaseInvalidAccountNameRule() *AzurermCosmosdbMongoDatabaseInvalidAccountNameRule {
	return &AzurermCosmosdbMongoDatabaseInvalidAccountNameRule{
		resourceType:  "azurerm_cosmosdb_mongo_database",
		attributeName: "account_name",
		pattern:       regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*`),
	}
}

// Name returns the rule name
func (r *AzurermCosmosdbMongoDatabaseInvalidAccountNameRule) Name() string {
	return "azurerm_cosmosdb_mongo_database_invalid_account_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermCosmosdbMongoDatabaseInvalidAccountNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermCosmosdbMongoDatabaseInvalidAccountNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermCosmosdbMongoDatabaseInvalidAccountNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AzurermCosmosdbMongoDatabaseInvalidAccountNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z0-9]+(-[a-z0-9]+)*`),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}