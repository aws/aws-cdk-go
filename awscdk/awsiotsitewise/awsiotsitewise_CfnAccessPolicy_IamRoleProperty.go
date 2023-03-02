package awsiotsitewise


// Contains information about an AWS Identity and Access Management role.
//
// For more information, see [IAM roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html) in the *IAM User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamRoleProperty := &iamRoleProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnAccessPolicy_IamRoleProperty struct {
	// The ARN of the IAM role.
	//
	// For more information, see [IAM ARNs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html) in the *IAM User Guide* .
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

