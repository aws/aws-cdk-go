package awssso


// Specifies the name and path of a customer managed policy.
//
// You must have an IAM policy that matches the name and path in each AWS account where you want to deploy your permission set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerManagedPolicyReferenceProperty := &customerManagedPolicyReferenceProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	path: jsii.String("path"),
//   }
//
type CfnPermissionSet_CustomerManagedPolicyReferenceProperty struct {
	// The name of the IAM policy that you have configured in each account where you want to deploy your permission set.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The path to the IAM policy that you have configured in each account where you want to deploy your permission set.
	//
	// The default is `/` . For more information, see [Friendly names and paths](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names) in the *IAM User Guide* .
	Path *string `field:"optional" json:"path" yaml:"path"`
}

