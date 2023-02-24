package awsconfig


// Provides the runtime system, policy definition, and whether debug logging enabled.
//
// You can specify the following CustomPolicyDetails parameter values only for AWS Config Custom Policy rules.
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
type CfnConfigRule_CustomPolicyDetailsProperty struct {
	// The boolean expression for enabling debug logging for your AWS Config Custom Policy rule.
	//
	// The default value is `false` .
	EnableDebugLogDelivery interface{} `field:"optional" json:"enableDebugLogDelivery" yaml:"enableDebugLogDelivery"`
	// The runtime system for your AWS Config Custom Policy rule.
	//
	// Guard is a policy-as-code language that allows you to write policies that are enforced by AWS Config Custom Policy rules. For more information about Guard, see the [Guard GitHub Repository](https://docs.aws.amazon.com/https://github.com/aws-cloudformation/cloudformation-guard) .
	PolicyRuntime *string `field:"optional" json:"policyRuntime" yaml:"policyRuntime"`
	// The policy definition containing the logic for your AWS Config Custom Policy rule.
	PolicyText *string `field:"optional" json:"policyText" yaml:"policyText"`
}

