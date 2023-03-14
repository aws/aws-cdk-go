package awscodeguruprofiler


// Notification medium for users to get alerted for events that occur in application profile.
//
// We support SNS topic as a notification channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelProperty := &channelProperty{
//   	channelUri: jsii.String("channelUri"),
//
//   	// the properties below are optional
//   	channelId: jsii.String("channelId"),
//   }
//
type CfnProfilingGroup_ChannelProperty struct {
	// The channel URI.
	ChannelUri *string `field:"required" json:"channelUri" yaml:"channelUri"`
	// The channel ID.
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
}

