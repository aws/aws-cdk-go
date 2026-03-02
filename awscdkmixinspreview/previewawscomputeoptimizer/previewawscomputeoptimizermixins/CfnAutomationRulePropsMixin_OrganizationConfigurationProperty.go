package previewawscomputeoptimizermixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   organizationConfigurationProperty := &OrganizationConfigurationProperty{
//   	AccountIds: []*string{
//   		jsii.String("accountIds"),
//   	},
//   	RuleApplyOrder: jsii.String("ruleApplyOrder"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-organizationconfiguration.html
//
type CfnAutomationRulePropsMixin_OrganizationConfigurationProperty struct {
	// List of account IDs where the organization rule applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-organizationconfiguration.html#cfn-computeoptimizer-automationrule-organizationconfiguration-accountids
	//
	AccountIds *[]*string `field:"optional" json:"accountIds" yaml:"accountIds"`
	// When the rule should be applied relative to account rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-organizationconfiguration.html#cfn-computeoptimizer-automationrule-organizationconfiguration-ruleapplyorder
	//
	RuleApplyOrder *string `field:"optional" json:"ruleApplyOrder" yaml:"ruleApplyOrder"`
}

