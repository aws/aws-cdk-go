package awsec2


// Describes an IAM instance profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamInstanceProfileSpecificationProperty := &iamInstanceProfileSpecificationProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnSpotFleet_IamInstanceProfileSpecificationProperty struct {
	// The Amazon Resource Name (ARN) of the instance profile.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

