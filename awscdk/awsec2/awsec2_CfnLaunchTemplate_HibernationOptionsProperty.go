package awsec2


// Specifies whether your instance is configured for hibernation.
//
// This parameter is valid only if the instance meets the [hibernation prerequisites](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html#hibernating-prerequisites) . For more information, see [Hibernate Your Instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the *Amazon EC2 User Guide* .
//
// `HibernationOptions` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hibernationOptionsProperty := &hibernationOptionsProperty{
//   	configured: jsii.Boolean(false),
//   }
//
type CfnLaunchTemplate_HibernationOptionsProperty struct {
	// If you set this parameter to `true` , the instance is enabled for hibernation.
	//
	// Default: `false`.
	Configured interface{} `field:"optional" json:"configured" yaml:"configured"`
}

