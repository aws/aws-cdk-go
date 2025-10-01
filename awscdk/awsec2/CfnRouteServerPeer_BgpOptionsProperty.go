package awsec2


// The BGP configuration options for this peer, including ASN (Autonomous System Number) and BFD (Bidrectional Forwarding Detection) settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bgpOptionsProperty := &BgpOptionsProperty{
//   	PeerAsn: jsii.Number(123),
//   	PeerLivenessDetection: jsii.String("peerLivenessDetection"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-routeserverpeer-bgpoptions.html
//
type CfnRouteServerPeer_BgpOptionsProperty struct {
	// The Border Gateway Protocol (BGP) Autonomous System Number (ASN) for the appliance.
	//
	// Valid values are from 1 to 4294967295. We recommend using a private ASN in the 64512–65534 (16-bit ASN) or 4200000000–4294967294 (32-bit ASN) range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-routeserverpeer-bgpoptions.html#cfn-ec2-routeserverpeer-bgpoptions-peerasn
	//
	PeerAsn *float64 `field:"optional" json:"peerAsn" yaml:"peerAsn"`
	// The liveness detection protocol used for the BGP peer.
	//
	// The requested liveness detection protocol for the BGP peer.
	//
	// - `bgp-keepalive` : The standard BGP keep alive mechanism ( [RFC4271](https://docs.aws.amazon.com/https://www.rfc-editor.org/rfc/rfc4271#page-21) ) that is stable but may take longer to fail-over in cases of network impact or router failure.
	// - `bfd` : An additional Bidirectional Forwarding Detection (BFD) protocol ( [RFC5880](https://docs.aws.amazon.com/https://www.rfc-editor.org/rfc/rfc5880) ) that enables fast failover by using more sensitive liveness detection.
	//
	// Defaults to `bgp-keepalive` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-routeserverpeer-bgpoptions.html#cfn-ec2-routeserverpeer-bgpoptions-peerlivenessdetection
	//
	PeerLivenessDetection *string `field:"optional" json:"peerLivenessDetection" yaml:"peerLivenessDetection"`
}

