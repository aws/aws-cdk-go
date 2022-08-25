package awsiotsitewise


// Contains information about an AWS Identity and Access Management user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamUserProperty := &iamUserProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnAccessPolicy_IamUserProperty struct {
	// The ARN of the IAM user. For more information, see [IAM ARNs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html) in the *IAM User Guide* .
	//
	// > If you delete the IAM user, access policies that contain this identity include an empty `arn` . You can delete the access policy for the IAM user that no longer exists.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

