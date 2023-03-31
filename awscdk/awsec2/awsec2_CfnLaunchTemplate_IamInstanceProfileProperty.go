package awsec2


// Specifies an IAM instance profile, which is a container for an IAM role for your instance.
//
// You can use an IAM role to distribute your AWS credentials to your instances.
//
// If you are creating the launch template for use with an Amazon EC2 Auto Scaling group, you can specify either the name or the ARN of the instance profile, but not both.
//
// `IamInstanceProfile` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamInstanceProfileProperty := &iamInstanceProfileProperty{
//   	arn: jsii.String("arn"),
//   	name: jsii.String("name"),
//   }
//
type CfnLaunchTemplate_IamInstanceProfileProperty struct {
	// The Amazon Resource Name (ARN) of the instance profile.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The name of the instance profile.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

