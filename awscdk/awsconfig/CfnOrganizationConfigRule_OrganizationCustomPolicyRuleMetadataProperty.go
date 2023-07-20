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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustompolicyrulemetadata.html
//
type CfnOrganizationConfigRule_OrganizationCustomPolicyRuleMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustompolicyrulemetadata.html#cfn-config-organizationconfigrule-organizationcustompolicyrulemetadata-policytext
	//
	PolicyText *string `field:"required" json:"policyText" yaml:"policyText"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustompolicyrulemetadata.html#cfn-config-organizationconfigrule-organizationcustompolicyrulemetadata-runtime
	//
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustompolicyrulemetadata.html#cfn-config-organizationconfigrule-organizationcustompolicyrulemetadata-debuglogdeliveryaccounts
	//
	DebugLogDeliveryAccounts *[]*string `field:"optional" json:"debugLogDeliveryAccounts" yaml:"debugLogDeliveryAccounts"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustompolicyrulemetadata.html#cfn-config-organizationconfigrule-organizationcustompolicyrulemetadata-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustompolicyrulemetadata.html#cfn-config-organizationconfigrule-organizationcustompolicyrulemetadata-inputparameters
	//
	InputParameters *string `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustompolicyrulemetadata.html#cfn-config-organizationconfigrule-organizationcustompolicyrulemetadata-maximumexecutionfrequency
	//
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustompolicyrulemetadata.html#cfn-config-organizationconfigrule-organizationcustompolicyrulemetadata-organizationconfigruletriggertypes
	//
	OrganizationConfigRuleTriggerTypes *[]*string `field:"optional" json:"organizationConfigRuleTriggerTypes" yaml:"organizationConfigRuleTriggerTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustompolicyrulemetadata.html#cfn-config-organizationconfigrule-organizationcustompolicyrulemetadata-resourceidscope
	//
	ResourceIdScope *string `field:"optional" json:"resourceIdScope" yaml:"resourceIdScope"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustompolicyrulemetadata.html#cfn-config-organizationconfigrule-organizationcustompolicyrulemetadata-resourcetypesscope
	//
	ResourceTypesScope *[]*string `field:"optional" json:"resourceTypesScope" yaml:"resourceTypesScope"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustompolicyrulemetadata.html#cfn-config-organizationconfigrule-organizationcustompolicyrulemetadata-tagkeyscope
	//
	TagKeyScope *string `field:"optional" json:"tagKeyScope" yaml:"tagKeyScope"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationcustompolicyrulemetadata.html#cfn-config-organizationconfigrule-organizationcustompolicyrulemetadata-tagvaluescope
	//
	TagValueScope *string `field:"optional" json:"tagValueScope" yaml:"tagValueScope"`
}

