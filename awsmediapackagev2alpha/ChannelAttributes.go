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
}

