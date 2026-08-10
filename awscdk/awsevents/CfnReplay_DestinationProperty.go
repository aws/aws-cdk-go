package awsevents


// A ReplayDestination object that includes details about the destination for the replay.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationProperty := &DestinationProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-replay-destination.html
//
type CfnReplay_DestinationProperty struct {
	// The ARN of the event bus to replay events to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-replay-destination.html#cfn-events-replay-destination-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

