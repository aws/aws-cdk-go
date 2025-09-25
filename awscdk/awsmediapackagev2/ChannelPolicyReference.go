package awsmediapackagev2


// A reference to a ChannelPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelPolicyReference := &ChannelPolicyReference{
//   	ChannelGroupName: jsii.String("channelGroupName"),
//   	ChannelName: jsii.String("channelName"),
//   }
//
type ChannelPolicyReference struct {
	// The ChannelGroupName of the ChannelPolicy resource.
	ChannelGroupName *string `field:"required" json:"channelGroupName" yaml:"channelGroupName"`
	// The ChannelName of the ChannelPolicy resource.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
}

