package awscdk


// The report emitted by the plugin after evaluation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   policyValidationPluginReportBeta1 := &PolicyValidationPluginReportBeta1{
//   	Success: jsii.Boolean(false),
//   	Violations: []PolicyViolationBeta1{
//   		&PolicyViolationBeta1{
//   			Description: jsii.String("description"),
//   			RuleName: jsii.String("ruleName"),
//   			ViolatingResources: []PolicyViolatingResourceBeta1{
//   				&PolicyViolatingResourceBeta1{
//   					Locations: []*string{
//   						jsii.String("locations"),
//   					},
//   					ResourceLogicalId: jsii.String("resourceLogicalId"),
//   					TemplatePath: jsii.String("templatePath"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Fix: jsii.String("fix"),
//   			RuleMetadata: map[string]*string{
//   				"ruleMetadataKey": jsii.String("ruleMetadata"),
//   			},
//   			Severity: jsii.String("severity"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Metadata: map[string]*string{
//   		"metadataKey": jsii.String("metadata"),
//   	},
//   	PluginVersion: jsii.String("pluginVersion"),
//   }
//
// Deprecated: Use `PolicyValidationPluginReport` instead.
type PolicyValidationPluginReportBeta1 struct {
	// Whether or not the report was successful.
	// Deprecated: Use `PolicyValidationPluginReport` instead.
	Success *bool `field:"required" json:"success" yaml:"success"`
	// List of violations in the report.
	// Deprecated: Use `PolicyValidationPluginReport` instead.
	Violations *[]*PolicyViolationBeta1 `field:"required" json:"violations" yaml:"violations"`
	// Arbitrary information about the report.
	// Default: - no metadata.
	//
	// Deprecated: Use `PolicyValidationPluginReport` instead.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// The version of the plugin that created the report.
	// Default: - no version.
	//
	// Deprecated: Use `PolicyValidationPluginReport` instead.
	PluginVersion *string `field:"optional" json:"pluginVersion" yaml:"pluginVersion"`
}

