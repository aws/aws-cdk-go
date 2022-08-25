package awsiam


// Contains information about an attached policy.
//
// An attached policy is a managed policy that has been attached to a user, group, or role.
//
// For more information about managed policies, refer to [Managed Policies and Inline Policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the *IAM User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   policyProperty := &policyProperty{
//   	policyDocument: policyDocument,
//   	policyName: jsii.String("policyName"),
//   }
//
type CfnRole_PolicyProperty struct {
	// The policy document.
	PolicyDocument interface{} `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The friendly name (not ARN) identifying the policy.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

