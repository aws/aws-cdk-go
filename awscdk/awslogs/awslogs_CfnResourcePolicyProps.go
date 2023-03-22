package awslogs


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
//   }
//
type CfnResourcePolicyProps struct {
	// The details of the policy.
	//
	// It must be formatted in JSON, and you must use backslashes to escape characters that need to be escaped in JSON strings, such as double quote marks.
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The name of the resource policy.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

