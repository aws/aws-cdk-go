package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTrafficMirrorSession`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficMirrorSessionProps := &cfnTrafficMirrorSessionProps{
//   	networkInterfaceId: jsii.String("networkInterfaceId"),
//   	sessionNumber: jsii.Number(123),
//   	trafficMirrorFilterId: jsii.String("trafficMirrorFilterId"),
//   	trafficMirrorTargetId: jsii.String("trafficMirrorTargetId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	packetLength: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	virtualNetworkId: jsii.Number(123),
//   }
//
type CfnTrafficMirrorSessionProps struct {
	// The ID of the source network interface.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions.
	//
	// The first session with a matching filter is the one that mirrors the packets.
	//
	// Valid values are 1-32766.
	SessionNumber *float64 `field:"required" json:"sessionNumber" yaml:"sessionNumber"`
	// The ID of the Traffic Mirror filter.
	TrafficMirrorFilterId *string `field:"required" json:"trafficMirrorFilterId" yaml:"trafficMirrorFilterId"`
	// The ID of the Traffic Mirror target.
	TrafficMirrorTargetId *string `field:"required" json:"trafficMirrorTargetId" yaml:"trafficMirrorTargetId"`
	// The description of the Traffic Mirror session.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The number of bytes in each packet to mirror.
	//
	// These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror. For example, if you set this value to 100, then the first 100 bytes that meet the filter criteria are copied to the target.
	//
	// If you do not want to mirror the entire packet, use the `PacketLength` parameter to specify the number of bytes in each packet to mirror.
	PacketLength *float64 `field:"optional" json:"packetLength" yaml:"packetLength"`
	// The tags to assign to a Traffic Mirror session.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The VXLAN ID for the Traffic Mirror session.
	//
	// For more information about the VXLAN protocol, see [RFC 7348](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc7348) . If you do not specify a `VirtualNetworkId` , an account-wide unique id is chosen at random.
	VirtualNetworkId *float64 `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

