package awsec2


// Describes an IAM instance profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamInstanceProfileSpecificationProperty := &IamInstanceProfileSpecificationProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-iaminstanceprofilespecification.html
//
type CfnSpotFleet_IamInstanceProfileSpecificationProperty struct {
	// The Amazon Resource Name (ARN) of the instance profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-iaminstanceprofilespecification.html#cfn-ec2-spotfleet-iaminstanceprofilespecification-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

