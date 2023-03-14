package awsconfig


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationCustomCodeRuleMetadataProperty := &organizationCustomCodeRuleMetadataProperty{
//   	codeText: jsii.String("codeText"),
//   	runtime: jsii.String("runtime"),
//
//   	// the properties below are optional
//   	debugLogDeliveryAccounts: []*string{
//   		jsii.String("debugLogDeliveryAccounts"),
//   	},
//   	description: jsii.String("description"),
//   	inputParameters: jsii.String("inputParameters"),
//   	maximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   	organizationConfigRuleTriggerTypes: []*string{
//   		jsii.String("organizationConfigRuleTriggerTypes"),
//   	},
//   	resourceIdScope: jsii.String("resourceIdScope"),
//   	resourceTypesScope: []*string{
//   		jsii.String("resourceTypesScope"),
//   	},
//   	tagKeyScope: jsii.String("tagKeyScope"),
//   	tagValueScope: jsii.String("tagValueScope"),
//   }
//
type CfnOrganizationConfigRule_OrganizationCustomCodeRuleMetadataProperty struct {
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.CodeText`.
	CodeText *string `field:"required" json:"codeText" yaml:"codeText"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.Runtime`.
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.DebugLogDeliveryAccounts`.
	DebugLogDeliveryAccounts *[]*string `field:"optional" json:"debugLogDeliveryAccounts" yaml:"debugLogDeliveryAccounts"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.InputParameters`.
	InputParameters *string `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.MaximumExecutionFrequency`.
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.OrganizationConfigRuleTriggerTypes`.
	OrganizationConfigRuleTriggerTypes *[]*string `field:"optional" json:"organizationConfigRuleTriggerTypes" yaml:"organizationConfigRuleTriggerTypes"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.ResourceIdScope`.
	ResourceIdScope *string `field:"optional" json:"resourceIdScope" yaml:"resourceIdScope"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.ResourceTypesScope`.
	ResourceTypesScope *[]*string `field:"optional" json:"resourceTypesScope" yaml:"resourceTypesScope"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.TagKeyScope`.
	TagKeyScope *string `field:"optional" json:"tagKeyScope" yaml:"tagKeyScope"`
	// `CfnOrganizationConfigRule.OrganizationCustomCodeRuleMetadataProperty.TagValueScope`.
	TagValueScope *string `field:"optional" json:"tagValueScope" yaml:"tagValueScope"`
}

