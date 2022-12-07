package awsxray


// Properties for defining a `CfnResourcePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourcePolicyProps := &cfnResourcePolicyProps{
//   	policyDocument: jsii.String("policyDocument"),
//   	policyName: jsii.String("policyName"),
//
//   	// the properties below are optional
//   	bypassPolicyLockoutCheck: jsii.Boolean(false),
//   }
//
type CfnResourcePolicyProps struct {
	// `AWS::XRay::ResourcePolicy.PolicyDocument`.
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// `AWS::XRay::ResourcePolicy.PolicyName`.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// `AWS::XRay::ResourcePolicy.BypassPolicyLockoutCheck`.
	BypassPolicyLockoutCheck interface{} `field:"optional" json:"bypassPolicyLockoutCheck" yaml:"bypassPolicyLockoutCheck"`
}

