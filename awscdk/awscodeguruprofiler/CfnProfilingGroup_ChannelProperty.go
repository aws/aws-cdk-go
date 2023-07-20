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
//   channelProperty := &ChannelProperty{
//   	ChannelUri: jsii.String("channelUri"),
//
//   	// the properties below are optional
//   	ChannelId: jsii.String("channelId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeguruprofiler-profilinggroup-channel.html
//
type CfnProfilingGroup_ChannelProperty struct {
	// The channel URI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeguruprofiler-profilinggroup-channel.html#cfn-codeguruprofiler-profilinggroup-channel-channeluri
	//
	ChannelUri *string `field:"required" json:"channelUri" yaml:"channelUri"`
	// The channel ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeguruprofiler-profilinggroup-channel.html#cfn-codeguruprofiler-profilinggroup-channel-channelid
	//
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
}

