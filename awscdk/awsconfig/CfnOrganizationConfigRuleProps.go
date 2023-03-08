package awsconfig


// Properties for defining a `CfnOrganizationConfigRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOrganizationConfigRuleProps := &CfnOrganizationConfigRuleProps{
//   	OrganizationConfigRuleName: jsii.String("organizationConfigRuleName"),
//
//   	// the properties below are optional
//   	ExcludedAccounts: []*string{
//   		jsii.String("excludedAccounts"),
//   	},
//   	OrganizationCustomCodeRuleMetadata: &OrganizationCustomCodeRuleMetadataProperty{
//   		CodeText: jsii.String("codeText"),
//   		Runtime: jsii.String("runtime"),
//
//   		// the properties below are optional
//   		DebugLogDeliveryAccounts: []*string{
//   			jsii.String("debugLogDeliveryAccounts"),
//   		},
//   		Description: jsii.String("description"),
//   		InputParameters: jsii.String("inputParameters"),
//   		MaximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   		OrganizationConfigRuleTriggerTypes: []*string{
//   			jsii.String("organizationConfigRuleTriggerTypes"),
//   		},
//   		ResourceIdScope: jsii.String("resourceIdScope"),
//   		ResourceTypesScope: []*string{
//   			jsii.String("resourceTypesScope"),
//   		},
//   		TagKeyScope: jsii.String("tagKeyScope"),
//   		TagValueScope: jsii.String("tagValueScope"),
//   	},
//   	OrganizationCustomRuleMetadata: &OrganizationCustomRuleMetadataProperty{
//   		LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   		OrganizationConfigRuleTriggerTypes: []*string{
//   			jsii.String("organizationConfigRuleTriggerTypes"),
//   		},
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   		InputParameters: jsii.String("inputParameters"),
//   		MaximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   		ResourceIdScope: jsii.String("resourceIdScope"),
//   		ResourceTypesScope: []*string{
//   			jsii.String("resourceTypesScope"),
//   		},
//   		TagKeyScope: jsii.String("tagKeyScope"),
//   		TagValueScope: jsii.String("tagValueScope"),
//   	},
//   	OrganizationManagedRuleMetadata: &OrganizationManagedRuleMetadataProperty{
//   		RuleIdentifier: jsii.String("ruleIdentifier"),
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   		InputParameters: jsii.String("inputParameters"),
//   		MaximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   		ResourceIdScope: jsii.String("resourceIdScope"),
//   		ResourceTypesScope: []*string{
//   			jsii.String("resourceTypesScope"),
//   		},
//   		TagKeyScope: jsii.String("tagKeyScope"),
//   		TagValueScope: jsii.String("tagValueScope"),
//   	},
//   }
//
type CfnOrganizationConfigRuleProps struct {
	// The name that you assign to organization AWS Config rule.
	OrganizationConfigRuleName *string `field:"required" json:"organizationConfigRuleName" yaml:"organizationConfigRuleName"`
	// A comma-separated list of accounts excluded from organization AWS Config rule.
	ExcludedAccounts *[]*string `field:"optional" json:"excludedAccounts" yaml:"excludedAccounts"`
	// `AWS::Config::OrganizationConfigRule.OrganizationCustomCodeRuleMetadata`.
	OrganizationCustomCodeRuleMetadata interface{} `field:"optional" json:"organizationCustomCodeRuleMetadata" yaml:"organizationCustomCodeRuleMetadata"`
	// An `OrganizationCustomRuleMetadata` object.
	OrganizationCustomRuleMetadata interface{} `field:"optional" json:"organizationCustomRuleMetadata" yaml:"organizationCustomRuleMetadata"`
	// An `OrganizationManagedRuleMetadata` object.
	OrganizationManagedRuleMetadata interface{} `field:"optional" json:"organizationManagedRuleMetadata" yaml:"organizationManagedRuleMetadata"`
}

