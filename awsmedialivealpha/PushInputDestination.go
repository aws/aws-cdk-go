package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
)

// A destination for a push-type input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkRef INetworkRef
//
//   pushInputDestination := &PushInputDestination{
//   	Network: networkRef,
//   	NetworkRoutes: []NetworkRoute{
//   		&NetworkRoute{
//   			Cidr: jsii.String("cidr"),
//   			Gateway: jsii.String("gateway"),
//   		},
//   	},
//   	StaticIpAddress: jsii.String("staticIpAddress"),
//   	StreamName: jsii.String("streamName"),
//   }
//
// Experimental.
type PushInputDestination struct {
	// The MediaLive Anywhere network this push destination lives on.
	//
	// Required when the input's
	// `inputNetworkLocation` is `ON_PREMISES`.
	// Default: - AWS-managed network.
	//
	// Experimental.
	Network interfacesawsmedialive.INetworkRef `field:"optional" json:"network" yaml:"network"`
	// The routes to the push destination on the local network.
	//
	// Required for
	// on-premises (`ON_PREMISES`) push inputs.
	// Default: - no routes.
	//
	// Experimental.
	NetworkRoutes *[]*NetworkRoute `field:"optional" json:"networkRoutes" yaml:"networkRoutes"`
	// A static IP address to assign to the push destination on the local network.
	// Default: - MediaLive Anywhere uses one from the IP pool specified on the selected network (service default).
	//
	// Experimental.
	StaticIpAddress *string `field:"optional" json:"staticIpAddress" yaml:"staticIpAddress"`
	// The stream name for RTMP push destinations (application name/instance).
	// Default: - auto-generated.
	//
	// Experimental.
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

