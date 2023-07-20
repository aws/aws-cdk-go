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
//   	OrganizationCustomPolicyRuleMetadata: &OrganizationCustomPolicyRuleMetadataProperty{
//   		PolicyText: jsii.String("policyText"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconfigrule.html
//
type CfnOrganizationConfigRuleProps struct {
	// The name that you assign to organization AWS Config rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconfigrule.html#cfn-config-organizationconfigrule-organizationconfigrulename
	//
	OrganizationConfigRuleName *string `field:"required" json:"organizationConfigRuleName" yaml:"organizationConfigRuleName"`
	// A comma-separated list of accounts excluded from organization AWS Config rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconfigrule.html#cfn-config-organizationconfigrule-excludedaccounts
	//
	ExcludedAccounts *[]*string `field:"optional" json:"excludedAccounts" yaml:"excludedAccounts"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconfigrule.html#cfn-config-organizationconfigrule-organizationcustompolicyrulemetadata
	//
	OrganizationCustomPolicyRuleMetadata interface{} `field:"optional" json:"organizationCustomPolicyRuleMetadata" yaml:"organizationCustomPolicyRuleMetadata"`
	// An `OrganizationCustomRuleMetadata` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconfigrule.html#cfn-config-organizationconfigrule-organizationcustomrulemetadata
	//
	OrganizationCustomRuleMetadata interface{} `field:"optional" json:"organizationCustomRuleMetadata" yaml:"organizationCustomRuleMetadata"`
	// An `OrganizationManagedRuleMetadata` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconfigrule.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata
	//
	OrganizationManagedRuleMetadata interface{} `field:"optional" json:"organizationManagedRuleMetadata" yaml:"organizationManagedRuleMetadata"`
}

