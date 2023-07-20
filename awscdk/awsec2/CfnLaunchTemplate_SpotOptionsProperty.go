package awsec2


// Specifies options for Spot Instances.
//
// `SpotOptions` is a property of [AWS::EC2::LaunchTemplate InstanceMarketOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-instancemarketoptions.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotOptionsProperty := &SpotOptionsProperty{
//   	BlockDurationMinutes: jsii.Number(123),
//   	InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   	MaxPrice: jsii.String("maxPrice"),
//   	SpotInstanceType: jsii.String("spotInstanceType"),
//   	ValidUntil: jsii.String("validUntil"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-spotoptions.html
//
type CfnLaunchTemplate_SpotOptionsProperty struct {
	// Deprecated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-spotoptions.html#cfn-ec2-launchtemplate-spotoptions-blockdurationminutes
	//
	BlockDurationMinutes *float64 `field:"optional" json:"blockDurationMinutes" yaml:"blockDurationMinutes"`
	// The behavior when a Spot Instance is interrupted.
	//
	// The default is `terminate` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-spotoptions.html#cfn-ec2-launchtemplate-spotoptions-instanceinterruptionbehavior
	//
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// The maximum hourly price you're willing to pay for the Spot Instances.
	//
	// We do not recommend using this parameter because it can lead to increased interruptions. If you do not specify this parameter, you will pay the current Spot price.
	//
	// > If you specify a maximum price, your Spot Instances will be interrupted more frequently than if you do not specify this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-spotoptions.html#cfn-ec2-launchtemplate-spotoptions-maxprice
	//
	MaxPrice *string `field:"optional" json:"maxPrice" yaml:"maxPrice"`
	// The Spot Instance request type.
	//
	// If you are using Spot Instances with an Auto Scaling group, use `one-time` requests, as the Amazon EC2 Auto Scaling service handles requesting new Spot Instances whenever the group is below its desired capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-spotoptions.html#cfn-ec2-launchtemplate-spotoptions-spotinstancetype
	//
	SpotInstanceType *string `field:"optional" json:"spotInstanceType" yaml:"spotInstanceType"`
	// The end date of the request, in UTC format ( *YYYY-MM-DD* T *HH:MM:SS* Z). Supported only for persistent requests.
	//
	// - For a persistent request, the request remains active until the `ValidUntil` date and time is reached. Otherwise, the request remains active until you cancel it.
	// - For a one-time request, `ValidUntil` is not supported. The request remains active until all instances launch or you cancel the request.
	//
	// Default: 7 days from the current date.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-spotoptions.html#cfn-ec2-launchtemplate-spotoptions-validuntil
	//
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
}

