package awsconfig


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationCustomPolicyRuleMetadataProperty := &OrganizationCustomPolicyRuleMetadataProperty{
//   	PolicyText: jsii.String("policyText"),
//   	Runtime: jsii.String("runtime"),
//
//   	// the properties below are optional
//   	DebugLogDeliveryAccounts: []*string{
//   		jsii.String("debugLogDeliveryAccounts"),
//   	},
//   	Description: jsii.String("description"),
//   	InputParameters: jsii.String("inputParameters"),
//   	MaximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   	OrganizationConfigRuleTriggerTypes: []*string{
//   		jsii.String("organizationConfigRuleTriggerTypes"),
//   	},
//   	ResourceIdScope: jsii.String("resourceIdScope"),
//   	ResourceTypesScope: []*string{
//   		jsii.String("resourceTypesScope"),
//   	},
//   	TagKeyScope: jsii.String("tagKeyScope"),
//   	TagValueScope: jsii.String("tagValueScope"),
//   }
//
type CfnOrganizationConfigRule_OrganizationCustomPolicyRuleMetadataProperty struct {
	// `CfnOrganizationConfigRule.OrganizationCustomPolicyRuleMetadataProperty.PolicyText`.
	PolicyText *string `field:"required" json:"policyText" yaml:"policyText"`
	// `CfnOrganizationConfigRule.OrganizationCustomPolicyRuleMetadataProperty.Runtime`.
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// `CfnOrganizationConfigRule.OrganizationCustomPolicyRuleMetadataProperty.DebugLogDeliveryAccounts`.
	DebugLogDeliveryAccounts *[]*string `field:"optional" json:"debugLogDeliveryAccounts" yaml:"debugLogDeliveryAccounts"`
	// `CfnOrganizationConfigRule.OrganizationCustomPolicyRuleMetadataProperty.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `CfnOrganizationConfigRule.OrganizationCustomPolicyRuleMetadataProperty.InputParameters`.
	InputParameters *string `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// `CfnOrganizationConfigRule.OrganizationCustomPolicyRuleMetadataProperty.MaximumExecutionFrequency`.
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// `CfnOrganizationConfigRule.OrganizationCustomPolicyRuleMetadataProperty.OrganizationConfigRuleTriggerTypes`.
	OrganizationConfigRuleTriggerTypes *[]*string `field:"optional" json:"organizationConfigRuleTriggerTypes" yaml:"organizationConfigRuleTriggerTypes"`
	// `CfnOrganizationConfigRule.OrganizationCustomPolicyRuleMetadataProperty.ResourceIdScope`.
	ResourceIdScope *string `field:"optional" json:"resourceIdScope" yaml:"resourceIdScope"`
	// `CfnOrganizationConfigRule.OrganizationCustomPolicyRuleMetadataProperty.ResourceTypesScope`.
	ResourceTypesScope *[]*string `field:"optional" json:"resourceTypesScope" yaml:"resourceTypesScope"`
	// `CfnOrganizationConfigRule.OrganizationCustomPolicyRuleMetadataProperty.TagKeyScope`.
	TagKeyScope *string `field:"optional" json:"tagKeyScope" yaml:"tagKeyScope"`
	// `CfnOrganizationConfigRule.OrganizationCustomPolicyRuleMetadataProperty.TagValueScope`.
	TagValueScope *string `field:"optional" json:"tagValueScope" yaml:"tagValueScope"`
}

