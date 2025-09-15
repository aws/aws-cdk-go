package awsconfig


// Provides the CustomPolicyDetails, the rule owner ( `AWS` for managed rules, `CUSTOM_POLICY` for Custom Policy rules, and `CUSTOM_LAMBDA` for Custom Lambda rules), the rule identifier, and the events that cause the evaluation of your AWS resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customPolicyDetailsProperty := &CustomPolicyDetailsProperty{
//   	EnableDebugLogDelivery: jsii.Boolean(false),
//   	PolicyRuntime: jsii.String("policyRuntime"),
//   	PolicyText: jsii.String("policyText"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-custompolicydetails.html
//
type CfnConfigRule_CustomPolicyDetailsProperty struct {
	// The boolean expression for enabling debug logging for your AWS Config Custom Policy rule.
	//
	// The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-custompolicydetails.html#cfn-config-configrule-custompolicydetails-enabledebuglogdelivery
	//
	EnableDebugLogDelivery interface{} `field:"optional" json:"enableDebugLogDelivery" yaml:"enableDebugLogDelivery"`
	// The runtime system for your AWS Config Custom Policy rule.
	//
	// Guard is a policy-as-code language that allows you to write policies that are enforced by AWS Config Custom Policy rules. For more information about Guard, see the [Guard GitHub Repository](https://docs.aws.amazon.com/https://github.com/aws-cloudformation/cloudformation-guard) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-custompolicydetails.html#cfn-config-configrule-custompolicydetails-policyruntime
	//
	PolicyRuntime *string `field:"optional" json:"policyRuntime" yaml:"policyRuntime"`
	// The policy definition containing the logic for your AWS Config Custom Policy rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-custompolicydetails.html#cfn-config-configrule-custompolicydetails-policytext
	//
	PolicyText *string `field:"optional" json:"policyText" yaml:"policyText"`
}

