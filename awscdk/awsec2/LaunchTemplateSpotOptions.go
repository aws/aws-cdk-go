package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Interface for the Spot market instance options provided in a LaunchTemplate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var expiration expiration
//
//   launchTemplateSpotOptions := &LaunchTemplateSpotOptions{
//   	BlockDuration: cdk.Duration_Minutes(jsii.Number(30)),
//   	InterruptionBehavior: awscdk.Aws_ec2.SpotInstanceInterruption_STOP,
//   	MaxPrice: jsii.Number(123),
//   	RequestType: awscdk.*Aws_ec2.SpotRequestType_ONE_TIME,
//   	ValidUntil: expiration,
//   }
//
type LaunchTemplateSpotOptions struct {
	// Spot Instances with a defined duration (also known as Spot blocks) are designed not to be interrupted and will run continuously for the duration you select.
	//
	// You can use a duration of 1, 2, 3, 4, 5, or 6 hours.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-requests.html#fixed-duration-spot-instances
	//
	// Default: Requested spot instances do not have a pre-defined duration.
	//
	BlockDuration awscdk.Duration `field:"optional" json:"blockDuration" yaml:"blockDuration"`
	// The behavior when a Spot Instance is interrupted.
	// Default: Spot instances will terminate when interrupted.
	//
	InterruptionBehavior SpotInstanceInterruption `field:"optional" json:"interruptionBehavior" yaml:"interruptionBehavior"`
	// Maximum hourly price you're willing to pay for each Spot instance.
	//
	// The value is given
	// in dollars. ex: 0.01 for 1 cent per hour, or 0.001 for one-tenth of a cent per hour.
	// Default: Maximum hourly price will default to the on-demand price for the instance type.
	//
	MaxPrice *float64 `field:"optional" json:"maxPrice" yaml:"maxPrice"`
	// The Spot Instance request type.
	//
	// If you are using Spot Instances with an Auto Scaling group, use one-time requests, as the
	// Amazon EC2 Auto Scaling service handles requesting new Spot Instances whenever the group is
	// below its desired capacity.
	// Default: One-time spot request.
	//
	RequestType SpotRequestType `field:"optional" json:"requestType" yaml:"requestType"`
	// The end date of the request.
	//
	// For a one-time request, the request remains active until all instances
	// launch, the request is canceled, or this date is reached. If the request is persistent, it remains
	// active until it is canceled or this date and time is reached.
	// Default: The default end date is 7 days from the current date.
	//
	ValidUntil awscdk.Expiration `field:"optional" json:"validUntil" yaml:"validUntil"`
}

