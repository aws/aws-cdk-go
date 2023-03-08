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
	// The resource policy document, which can be up to 5kb in size.
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The name of the resource policy.
	//
	// Must be unique within a specific AWS account.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// `AWS::XRay::ResourcePolicy.BypassPolicyLockoutCheck`.
	BypassPolicyLockoutCheck interface{} `field:"optional" json:"bypassPolicyLockoutCheck" yaml:"bypassPolicyLockoutCheck"`
}

