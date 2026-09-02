package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
)

// Properties for push-type inputs (RTMP_PUSH, RTP_PUSH, UDP_PUSH).
//
// Example:
//   var stack Stack
//
//
//   network := medialive.NewNetwork(stack, jsii.String("Network"), &NetworkProps{
//   	NetworkName: jsii.String("on-prem-network"),
//   	IpPools: []*string{
//   		jsii.String("192.168.1.0/24"),
//   	},
//   })
//
//   medialive.NewInput(stack, jsii.String("OnPremInput"), &InputProps{
//   	InputName: jsii.String("on-prem-rtp"),
//   	InputNetworkLocation: medialive.InputNetworkLocation_ON_PREMISES(),
//   	Input: medialive.InputConfiguration_RtpPush(&PushInputProps{
//   		Destinations: []PushInputDestination{
//   			&PushInputDestination{
//   				Network: *Network,
//   				NetworkRoutes: []NetworkRoute{
//   					&NetworkRoute{
//   						Cidr: jsii.String("10.0.0.0/24"),
//   						Gateway: jsii.String("10.0.0.1"),
//   					},
//   				},
//   				StaticIpAddress: jsii.String("192.168.1.50"),
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type PushInputProps struct {
	// The destinations for push inputs.
	//
	// For RTMP push, each destination can have a stream name.
	// Default: - MediaLive auto-generates destinations.
	//
	// Experimental.
	Destinations *[]*PushInputDestination `field:"optional" json:"destinations" yaml:"destinations"`
	// The input security groups that control which CIDR blocks can push to this input.
	//
	// Required for
	// cloud inputs; must be omitted for on-premises inputs (`InputNetworkLocation.ON_PREMISES`),
	// which do not support security groups.
	// Default: - none (only valid for on-premises inputs).
	//
	// Experimental.
	InputSecurityGroups *[]interfacesawsmedialive.IInputSecurityGroupRef `field:"optional" json:"inputSecurityGroups" yaml:"inputSecurityGroups"`
}

