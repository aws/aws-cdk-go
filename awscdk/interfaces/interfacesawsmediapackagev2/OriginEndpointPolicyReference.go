package interfacesawsmediapackagev2


// A reference to a OriginEndpointPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originEndpointPolicyReference := &OriginEndpointPolicyReference{
//   	ChannelGroupName: jsii.String("channelGroupName"),
//   	ChannelName: jsii.String("channelName"),
//   	OriginEndpointName: jsii.String("originEndpointName"),
//   }
//
type OriginEndpointPolicyReference struct {
	// The ChannelGroupName of the OriginEndpointPolicy resource.
	ChannelGroupName *string `field:"required" json:"channelGroupName" yaml:"channelGroupName"`
	// The ChannelName of the OriginEndpointPolicy resource.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The OriginEndpointName of the OriginEndpointPolicy resource.
	OriginEndpointName *string `field:"required" json:"originEndpointName" yaml:"originEndpointName"`
}

