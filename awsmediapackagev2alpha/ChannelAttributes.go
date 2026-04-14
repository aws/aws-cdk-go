package awsmediapackagev2alpha


// Represents a Channel defined outside of this stack.
//
// Example:
//   var stack Stack
//
//   channel := awsmediapackagev2alpha.Channel_FromChannelAttributes(stack, jsii.String("ImportedChannel"), &ChannelAttributes{
//   	ChannelName: jsii.String("MyChannel"),
//   	ChannelGroupName: jsii.String("MyChannelGroup"),
//   })
//
// Experimental.
type ChannelAttributes struct {
	// The name that describes the channel group.
	// Experimental.
	ChannelGroupName *string `field:"required" json:"channelGroupName" yaml:"channelGroupName"`
	// The name that describes the channel.
	// Experimental.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The AWS region where the channel lives.
	//
	// Required for cross-region imports to construct the correct ARN.
	// Default: - the importing stack's region.
	//
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

