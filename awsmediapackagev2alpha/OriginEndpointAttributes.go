package awsmediapackagev2alpha


// Represents an Origin Endpoint defined outside of this stack.
//
// Example:
//   var stack Stack
//
//   originEndpoint := awsmediapackagev2alpha.OriginEndpoint_FromOriginEndpointAttributes(stack, jsii.String("ImportedOriginEndpoint"), &OriginEndpointAttributes{
//   	ChannelGroupName: jsii.String("MyChannelGroup"),
//   	ChannelName: jsii.String("MyChannel"),
//   	OriginEndpointName: jsii.String("MyExampleOriginEndpoint"),
//   })
//
// Experimental.
type OriginEndpointAttributes struct {
	// The name that describes the channel group.
	//
	// The name is the primary identifier for the channel group.
	// Experimental.
	ChannelGroupName *string `field:"required" json:"channelGroupName" yaml:"channelGroupName"`
	// The name that describes the channel.
	//
	// The name is the primary identifier for the channel.
	// Experimental.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The name that describes the origin endpoint.
	// Experimental.
	OriginEndpointName *string `field:"required" json:"originEndpointName" yaml:"originEndpointName"`
	// The AWS region where the origin endpoint lives.
	//
	// Required for cross-region imports to construct the correct ARN.
	// Default: - the importing stack's region.
	//
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

