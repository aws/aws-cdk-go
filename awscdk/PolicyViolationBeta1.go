package awscdk


// Violation produced by the validation plugin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   policyViolationBeta1 := &PolicyViolationBeta1{
//   	Description: jsii.String("description"),
//   	RuleName: jsii.String("ruleName"),
//   	ViolatingResources: []policyViolatingResourceBeta1{
//   		&policyViolatingResourceBeta1{
//   			Locations: []*string{
//   				jsii.String("locations"),
//   			},
//   			ResourceLogicalId: jsii.String("resourceLogicalId"),
//   			TemplatePath: jsii.String("templatePath"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Fix: jsii.String("fix"),
//   	RuleMetadata: map[string]*string{
//   		"ruleMetadataKey": jsii.String("ruleMetadata"),
//   	},
//   	Severity: jsii.String("severity"),
//   }
//
type PolicyViolationBeta1 struct {
	// The description of the violation.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the rule.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// The resources violating this rule.
	ViolatingResources *[]*PolicyViolatingResourceBeta1 `field:"required" json:"violatingResources" yaml:"violatingResources"`
	// How to fix the violation.
	// Default: - no fix is provided.
	//
	Fix *string `field:"optional" json:"fix" yaml:"fix"`
	// Additional metadata to include with the rule results.
	//
	// This can be used to provide additional information that is
	// plugin specific. The data provided here will be rendered as is.
	// Default: - no rule metadata.
	//
	RuleMetadata *map[string]*string `field:"optional" json:"ruleMetadata" yaml:"ruleMetadata"`
	// The severity of the violation, only used for reporting purposes.
	//
	// This is useful for helping the user discriminate between warnings,
	// errors, information, etc.
	// Default: - no severity.
	//
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
}

