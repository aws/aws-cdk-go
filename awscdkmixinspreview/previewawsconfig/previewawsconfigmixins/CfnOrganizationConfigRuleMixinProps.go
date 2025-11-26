package previewawsconfigmixins


// Properties for CfnOrganizationConfigRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOrganizationConfigRuleMixinProps := &CfnOrganizationConfigRuleMixinProps{
//   	ExcludedAccounts: []*string{
//   		jsii.String("excludedAccounts"),
//   	},
//   	OrganizationConfigRuleName: jsii.String("organizationConfigRuleName"),
//   	OrganizationCustomPolicyRuleMetadata: &OrganizationCustomPolicyRuleMetadataProperty{
//   		DebugLogDeliveryAccounts: []*string{
//   			jsii.String("debugLogDeliveryAccounts"),
//   		},
//   		Description: jsii.String("description"),
//   		InputParameters: jsii.String("inputParameters"),
//   		MaximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   		OrganizationConfigRuleTriggerTypes: []*string{
//   			jsii.String("organizationConfigRuleTriggerTypes"),
//   		},
//   		PolicyText: jsii.String("policyText"),
//   		ResourceIdScope: jsii.String("resourceIdScope"),
//   		ResourceTypesScope: []*string{
//   			jsii.String("resourceTypesScope"),
//   		},
//   		Runtime: jsii.String("runtime"),
//   		TagKeyScope: jsii.String("tagKeyScope"),
//   		TagValueScope: jsii.String("tagValueScope"),
//   	},
//   	OrganizationCustomRuleMetadata: &OrganizationCustomRuleMetadataProperty{
//   		Description: jsii.String("description"),
//   		InputParameters: jsii.String("inputParameters"),
//   		LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
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
//   	OrganizationManagedRuleMetadata: &OrganizationManagedRuleMetadataProperty{
//   		Description: jsii.String("description"),
//   		InputParameters: jsii.String("inputParameters"),
//   		MaximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   		ResourceIdScope: jsii.String("resourceIdScope"),
//   		ResourceTypesScope: []*string{
//   			jsii.String("resourceTypesScope"),
//   		},
//   		RuleIdentifier: jsii.String("ruleIdentifier"),
//   		TagKeyScope: jsii.String("tagKeyScope"),
//   		TagValueScope: jsii.String("tagValueScope"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconfigrule.html
//
type CfnOrganizationConfigRuleMixinProps struct {
	// A comma-separated list of accounts excluded from organization AWS Config rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconfigrule.html#cfn-config-organizationconfigrule-excludedaccounts
	//
	ExcludedAccounts *[]*string `field:"optional" json:"excludedAccounts" yaml:"excludedAccounts"`
	// The name that you assign to organization AWS Config rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconfigrule.html#cfn-config-organizationconfigrule-organizationconfigrulename
	//
	OrganizationConfigRuleName *string `field:"optional" json:"organizationConfigRuleName" yaml:"organizationConfigRuleName"`
	// An object that specifies metadata for your organization's AWS Config Custom Policy rule.
	//
	// The metadata includes the runtime system in use, which accounts have debug logging enabled, and other custom rule metadata, such as resource type, resource ID of AWS resource, and organization trigger types that initiate AWS Config to evaluate AWS resources against a rule.
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

