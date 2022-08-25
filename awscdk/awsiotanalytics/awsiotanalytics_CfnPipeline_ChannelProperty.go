package awsiotanalytics


// Determines the source of the messages to be processed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelProperty := &channelProperty{
//   	channelName: jsii.String("channelName"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_ChannelProperty struct {
	// The name of the channel from which the messages are processed.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The name of the 'channel' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}

