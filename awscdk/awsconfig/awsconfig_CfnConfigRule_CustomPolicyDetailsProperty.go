package awsconfig


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customPolicyDetailsProperty := &customPolicyDetailsProperty{
//   	enableDebugLogDelivery: jsii.Boolean(false),
//   	policyRuntime: jsii.String("policyRuntime"),
//   	policyText: jsii.String("policyText"),
//   }
//
type CfnConfigRule_CustomPolicyDetailsProperty struct {
	// `CfnConfigRule.CustomPolicyDetailsProperty.EnableDebugLogDelivery`.
	EnableDebugLogDelivery interface{} `field:"optional" json:"enableDebugLogDelivery" yaml:"enableDebugLogDelivery"`
	// `CfnConfigRule.CustomPolicyDetailsProperty.PolicyRuntime`.
	PolicyRuntime *string `field:"optional" json:"policyRuntime" yaml:"policyRuntime"`
	// `CfnConfigRule.CustomPolicyDetailsProperty.PolicyText`.
	PolicyText *string `field:"optional" json:"policyText" yaml:"policyText"`
}

