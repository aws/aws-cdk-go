package awsec2


// BGP Options.
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
	// BGP ASN of the Route Server Peer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-routeserverpeer-bgpoptions.html#cfn-ec2-routeserverpeer-bgpoptions-peerasn
	//
	PeerAsn *float64 `field:"optional" json:"peerAsn" yaml:"peerAsn"`
	// BGP Liveness Detection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-routeserverpeer-bgpoptions.html#cfn-ec2-routeserverpeer-bgpoptions-peerlivenessdetection
	//
	PeerLivenessDetection *string `field:"optional" json:"peerLivenessDetection" yaml:"peerLivenessDetection"`
}

