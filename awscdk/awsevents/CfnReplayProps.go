package awsevents


// Properties for defining a `CfnReplay`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplayProps := &CfnReplayProps{
//   	Destination: &DestinationProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	EventEndTime: jsii.String("eventEndTime"),
//   	EventSourceArn: jsii.String("eventSourceArn"),
//   	EventStartTime: jsii.String("eventStartTime"),
//   	ReplayName: jsii.String("replayName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-replay.html
//
type CfnReplayProps struct {
	// A ReplayDestination object that includes details about the destination for the replay.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-replay.html#cfn-events-replay-destination
	//
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// A time stamp for the time to stop replaying events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-replay.html#cfn-events-replay-eventendtime
	//
	EventEndTime *string `field:"required" json:"eventEndTime" yaml:"eventEndTime"`
	// The ARN of the archive to replay events from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-replay.html#cfn-events-replay-eventsourcearn
	//
	EventSourceArn *string `field:"required" json:"eventSourceArn" yaml:"eventSourceArn"`
	// A time stamp for the time to start replaying events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-replay.html#cfn-events-replay-eventstarttime
	//
	EventStartTime *string `field:"required" json:"eventStartTime" yaml:"eventStartTime"`
	// The name of the replay.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-replay.html#cfn-events-replay-replayname
	//
	ReplayName *string `field:"required" json:"replayName" yaml:"replayName"`
}

