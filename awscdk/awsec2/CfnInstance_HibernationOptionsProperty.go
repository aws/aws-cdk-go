package awsec2


// Specifies the hibernation options for the instance.
//
// `HibernationOptions` is a property of the [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hibernationOptionsProperty := &HibernationOptionsProperty{
//   	Configured: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-hibernationoptions.html
//
type CfnInstance_HibernationOptionsProperty struct {
	// Set to `true` to enable your instance for hibernation.
	//
	// For Spot Instances, if you set `Configured` to `true` , either omit the `InstanceInterruptionBehavior` parameter (for [`SpotMarketOptions`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotMarketOptions.html) ), or set it to `hibernate` . When `Configured` is true:
	//
	// - If you omit `InstanceInterruptionBehavior` , it defaults to `hibernate` .
	// - If you set `InstanceInterruptionBehavior` to a value other than `hibernate` , you'll get an error.
	//
	// Default: `false`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-hibernationoptions.html#cfn-ec2-instance-hibernationoptions-configured
	//
	Configured interface{} `field:"optional" json:"configured" yaml:"configured"`
}

