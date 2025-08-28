package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRouteServerPeer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouteServerPeerProps := &CfnRouteServerPeerProps{
//   	BgpOptions: &BgpOptionsProperty{
//   		PeerAsn: jsii.Number(123),
//   		PeerLivenessDetection: jsii.String("peerLivenessDetection"),
//   	},
//   	PeerAddress: jsii.String("peerAddress"),
//   	RouteServerEndpointId: jsii.String("routeServerEndpointId"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverpeer.html
//
type CfnRouteServerPeerProps struct {
	// The BGP configuration options for this peer, including ASN (Autonomous System Number) and BFD (Bidrectional Forwarding Detection) settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverpeer.html#cfn-ec2-routeserverpeer-bgpoptions
	//
	BgpOptions interface{} `field:"required" json:"bgpOptions" yaml:"bgpOptions"`
	// The IPv4 address of the peer device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverpeer.html#cfn-ec2-routeserverpeer-peeraddress
	//
	PeerAddress *string `field:"required" json:"peerAddress" yaml:"peerAddress"`
	// The ID of the route server endpoint associated with this peer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverpeer.html#cfn-ec2-routeserverpeer-routeserverendpointid
	//
	RouteServerEndpointId *string `field:"required" json:"routeServerEndpointId" yaml:"routeServerEndpointId"`
	// Any tags assigned to the route server peer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverpeer.html#cfn-ec2-routeserverpeer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

