package awsmediapackagev2alpha


// Attributes to enable import of a Channel Group, which in turn can be used to create a Channel and OriginEndpoint).
//
// Example:
//   var stack Stack
//
//   channelGroup := awsmediapackagev2alpha.ChannelGroup_FromChannelGroupAttributes(stack, jsii.String("ImportedChannelGroup"), &ChannelGroupAttributes{
//   	ChannelGroupName: jsii.String("MyChannelGroup"),
//   })
//
// Experimental.
type ChannelGroupAttributes struct {
	// Channel Group Name.
	// Experimental.
	ChannelGroupName *string `field:"required" json:"channelGroupName" yaml:"channelGroupName"`
	// The egress domain where packaged content is available.
	//
	// Use this as the origin domain when configuring a CDN such as Amazon CloudFront.
	// Default: - not available on imported channel groups.
	//
	// Experimental.
	EgressDomain *string `field:"optional" json:"egressDomain" yaml:"egressDomain"`
}

