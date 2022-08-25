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
//   spotOptionsProperty := &spotOptionsProperty{
//   	blockDurationMinutes: jsii.Number(123),
//   	instanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   	maxPrice: jsii.String("maxPrice"),
//   	spotInstanceType: jsii.String("spotInstanceType"),
//   	validUntil: jsii.String("validUntil"),
//   }
//
type CfnLaunchTemplate_SpotOptionsProperty struct {
	// The required duration for the Spot Instances (also known as Spot blocks), in minutes.
	//
	// This value must be a multiple of 60 (60, 120, 180, 240, 300, or 360).
	BlockDurationMinutes *float64 `field:"optional" json:"blockDurationMinutes" yaml:"blockDurationMinutes"`
	// The behavior when a Spot Instance is interrupted.
	//
	// The default is `terminate` .
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// The maximum hourly price you're willing to pay for the Spot Instances.
	MaxPrice *string `field:"optional" json:"maxPrice" yaml:"maxPrice"`
	// The Spot Instance request type.
	//
	// If you are using Spot Instances with an Auto Scaling group, use `one-time` requests, as the Amazon EC2 Auto Scaling service handles requesting new Spot Instances whenever the group is below its desired capacity.
	SpotInstanceType *string `field:"optional" json:"spotInstanceType" yaml:"spotInstanceType"`
	// The end date of the request.
	//
	// For a one-time request, the request remains active until all instances launch, the request is canceled, or this date is reached. If the request is persistent, it remains active until it is canceled or this date and time is reached. The default end date is 7 days from the current date.
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
}

