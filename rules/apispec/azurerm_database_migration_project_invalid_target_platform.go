// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermDatabaseMigrationProjectInvalidTargetPlatformRule checks the pattern is valid
type AzurermDatabaseMigrationProjectInvalidTargetPlatformRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermDatabaseMigrationProjectInvalidTargetPlatformRule returns new rule with default attributes
func NewAzurermDatabaseMigrationProjectInvalidTargetPlatformRule() *AzurermDatabaseMigrationProjectInvalidTargetPlatformRule {
	return &AzurermDatabaseMigrationProjectInvalidTargetPlatformRule{
		resourceType:  "azurerm_database_migration_project",
		attributeName: "target_platform",
		enum: []string{
			"SQLDB",
			"Unknown",
		},
	}
}

// Name returns the rule name
func (r *AzurermDatabaseMigrationProjectInvalidTargetPlatformRule) Name() string {
	return "azurerm_database_migration_project_invalid_target_platform"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDatabaseMigrationProjectInvalidTargetPlatformRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDatabaseMigrationProjectInvalidTargetPlatformRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDatabaseMigrationProjectInvalidTargetPlatformRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermDatabaseMigrationProjectInvalidTargetPlatformRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as target_platform`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
